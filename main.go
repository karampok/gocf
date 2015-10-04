package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/DavidHuie/gomigrate"
	cfenv "github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/go-sql-driver/mysql"
	mig "github.com/karampok/gocf/migrations"
)

var (
	buildstamp string
	githash    string
	db         *sql.DB
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/info", info)
	log.Printf("Listening at %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func init() {
	log.SetOutput(os.Stdout)
	migPath := fmt.Sprintf("%s/migrations_data", os.Getenv("TMPDIR"))
	files := mig.AssetNames()
	for _, f := range files {
		err := mig.RestoreAsset(os.Getenv("TMPDIR"), f)
		if err != nil {
			log.Fatal("Migration were not found")
		}
	}
	cStr := getConnectionStr()
	log.Printf("Connect str %s ", cStr)
	db, err := sql.Open("mysql", cStr)
	if err != nil {
		log.Fatal("Can not open DB:", err)
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Printf("Connected to:%s\n", version)

	if err := db.Ping(); err != nil {
		log.Fatal("Error2", err.Error())
	}

	if _, err := os.Stat(migPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", migPath)
		return
	}
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Mariadb{}, migPath)
	if err := migrator.Migrate(); err != nil {
		log.Fatal("Can not execute migrations:", err)
	}
}

func getConnectionStr() string {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "mysql2://userOnoma:userKodikos@127.0.0.1:3306/72C61CC6BEED?reconnect=true"
	}
	u, err := url.Parse(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s@tcp(%s)%s", u.User.String(), u.Host, u.Path)
}
func getMariaService() string {
	appEnv, _ := cfenv.Current()
	s, err := appEnv.Services.WithName("kka-mariadb")
	if err != nil {
		log.Printf("NO SERVICE %s ", err)
	} else {

		log.Printf("s.name=%s\n", s.Name)
		log.Printf("s.label=%s\n", s.Label)
		log.Printf("s.tage=%v\n", s.Tags)
		log.Printf("s.plan=%s\n", s.Plan)
		for k, v := range s.Credentials {
			log.Printf("s.Credentials.%s : %v \n", k, v)
		}
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		s.Credentials["username"], s.Credentials["password"],
		s.Credentials["host"], s.Credentials["port"],
		s.Credentials["database"])
}
func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}

func info(w http.ResponseWriter, req *http.Request) {
	r := "Binary INFO:\n"
	r += fmt.Sprintf("githash=%s\n", githash)
	r += fmt.Sprintf("buildstamp=%s\n", buildstamp)
	r += fmt.Sprintf("\n\nENV Variables\n")
	for _, e := range os.Environ() {
		r += fmt.Sprintf("%s\n", e)
	}

	r += fmt.Sprintf("\n\nCF Variables\n")
	appEnv, err := cfenv.Current()
	if err != nil {
		panic("source the env")
	}
	r += fmt.Sprintf("ID:%s\n", appEnv.ID)
	r += fmt.Sprintf("Index:%s\n", appEnv.Index)
	r += fmt.Sprintf("Name:%s\n", appEnv.Name)
	r += fmt.Sprintf("Host:%s\n", appEnv.Host)
	r += fmt.Sprintf("Port:%s\n", appEnv.Port)
	r += fmt.Sprintf("Version:%s\n", appEnv.Version)
	r += fmt.Sprintf("Home:%s\n", appEnv.Home)
	r += fmt.Sprintf("MemoryLimit:%s\n", appEnv.MemoryLimit)
	r += fmt.Sprintf("WorkingDir:%s\n", appEnv.WorkingDir)
	r += fmt.Sprintf("TempDir:%s\n", appEnv.TempDir)
	r += fmt.Sprintf("User:%s\n", appEnv.User)
	r += fmt.Sprintf("Services:%s\n", appEnv.Services)
	r += fmt.Sprintf("\nALL:%v\n", appEnv)

	fmt.Fprintln(w, r)
	return
}
