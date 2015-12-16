package util

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/go-sql-driver/mysql"
	mig "github.com/karampok/gocf/migrations"
)

func RestoreData() {
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
