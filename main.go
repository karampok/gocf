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
	log.Printf("Connected to :%s\n", version)

	if err := db.Ping(); err != nil {
		log.Fatal("DB Error Ping", err.Error())
	}

	if _, err := os.Stat(migPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", migPath)
		return
	}
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Mariadb{}, migPath)
	if err := migrator.Migrate(); err != nil {
		log.Fatal("Can not execute migrations:", err)
	}
	dbstuff(db)
}

func dbstuff(db *sql.DB) {
	if _, err := db.Exec("INSERT INTO port_ranges(port_from,port_to) VALUES(45,55)"); err != nil {
		log.Printf("Can not insert data:%s", err)
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
	fmt.Fprintln(w, "hello swisscom cloud!")
}

func info(w http.ResponseWriter, req *http.Request) {
	appEnv, err := cfenv.Current()
	if err != nil {
		panic("source the env")
	}

	r := "Binary INFO:\n"
	r += fmt.Sprintf("githash=%s\n", githash)
	r += fmt.Sprintf("buildstamp=%s\n", buildstamp)
	r += fmt.Sprintf("\n\nENV Variables - execute\n")
	r += fmt.Sprintf("cf logs %s\n", appEnv.Name)
	r += fmt.Sprintf("cf logs --recent %s\n", appEnv.Name)

	fmt.Fprintln(w, r)

	log.Printf("\n\nENV Variables\n")
	for _, e := range os.Environ() {
		log.Printf("%s\n", e)
	}

	log.Printf("\n\nCF Variables\n")
	log.Printf("ID:%s\n", appEnv.ID)
	log.Printf("Index:%s\n", appEnv.Index)
	log.Printf("Name:%s\n", appEnv.Name)
	log.Printf("Host:%s\n", appEnv.Host)
	log.Printf("Port:%s\n", appEnv.Port)
	log.Printf("Version:%s\n", appEnv.Version)
	log.Printf("Home:%s\n", appEnv.Home)
	log.Printf("MemoryLimit:%s\n", appEnv.MemoryLimit)
	log.Printf("WorkingDir:%s\n", appEnv.WorkingDir)
	log.Printf("TempDir:%s\n", appEnv.TempDir)
	log.Printf("User:%s\n", appEnv.User)
	log.Printf("Services:%s\n", appEnv.Services)
	log.Printf("\nALL:%v\n", appEnv)

	return
}
