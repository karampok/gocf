package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/karampok/gocf/util"
)

var (
	buildstamp string
	githash    string
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/cfinfo", util.CfInfo)
	http.HandleFunc("/elk", playelk)
	http.HandleFunc("/info", info)
	log.Printf("Listening at %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func init() {
	log.SetOutput(os.Stdout)
	util.RestoreData()
}

func playelk(w http.ResponseWriter, req *http.Request) {
	var eUrl, kUrl, lUrl string
	appEnv, enverr := cfenv.Current()
	if enverr != nil {
		eUrl = "http://localhost:9200"
		kUrl = "http://localhost:5601"
		lUrl = "http://localhost:5000"
	} else {
		elk, err := appEnv.Services.WithTag("myelk")
		fmt.Println(appEnv)
		if err == nil {
			{ //BORING very specific
				user, _ := elk[0].Credentials["elasticSearchUsername"]
				pass, _ := elk[0].Credentials["elasticSearchPassword"]
				host, _ := elk[0].Credentials["elasticSearchHost"]
				port, _ := elk[0].Credentials["elasticSearchPort"]
				eUrl = fmt.Sprintf("http://%s:%s@%s:%f", user, pass, host, port)
			}
			{ //VERY BAD because we cannot insert easy the user/pass in url
				user, _ := url.Parse(elk[0].Credentials["kibanaUsername"].(string))
				pass, _ := url.Parse(elk[0].Credentials["kibanaPassword"].(string))
				url, _ := url.Parse(elk[0].Credentials["kibanaUrl"].(string))
				kUrl = fmt.Sprintf("%s  -u %s -p %s", url, user, pass)
			}
			{ //GOOD because I could use standard parsing libray (
				u, _ := url.Parse(elk[0].Credentials["APROPERLOGSTASHURL"].(string))
				lUrl = fmt.Sprintf("http://%s@%s", u.User.String(), u.Host)
			}
		} else {
			log.Fatal("Unable to find elastic search service")
		}
	}
	log.Printf("Elasticsearch url [%s]", eUrl)
	log.Printf("Kibana url [%s]", kUrl)
	log.Printf("Logstash url [%s]", lUrl)

	fmt.Fprintf(w, "hello at swisscom elk")
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello swisscom cloud!")
}

func info(w http.ResponseWriter, req *http.Request) {
	r := "Binary INFO:\n"
	r += fmt.Sprintf("buildstamp %s\n", buildstamp)
	r += fmt.Sprintf("githash %s\n", githash)

	r += fmt.Sprintf("\n\nENV Variables\n")
	for _, e := range os.Environ() {
		r += fmt.Sprintf("%s\n", e)
	}
	fmt.Fprintln(w, r)

	return
}
