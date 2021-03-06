package util

import (
	"fmt"
	"log"
	"net/http"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

func CfInfo(w http.ResponseWriter, req *http.Request) {
	appEnv, err := cfenv.Current()
	if err != nil {
		panic("source the env")
	}

	r := "Binary INFO:\n"
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
