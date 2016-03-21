package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"time"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
	"github.com/karampok/gocf/util"
)

type appconfig struct {
	elastic, logstash, kibana string
}

func (a appconfig) String() (r string) {
	r += fmt.Sprintf("Elasticsearch URL [%s]\n", a.elastic)
	r += fmt.Sprintf("Logstash URL [%s]\n", a.logstash)
	r += fmt.Sprintf("Kibana URL [%s]\n", a.kibana)
	return
}

var (
	buildstamp string
	githash    string
	app        appconfig
	rlog       *syslog.Writer
)

func init() {
	log.SetOutput(os.Stdout)
	e, l, k := util.SetELK("loadtest")
	app = appconfig{e, l, k}
	fmt.Println(app)
	//util.RestoreData()
}

func main() {
	//need a timeout
	rlog, err := syslog.Dial("tcp", app.logstash, syslog.LOG_INFO, "FromGO")
	defer rlog.Close()
	if err != nil {
		log.Fatalf("Cannot connect to remote syslog server,  %s", err)
	}

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}

	for i := 0; i < 1; i++ {
		go longJob()
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/elk", playelk)
	http.HandleFunc("/kibana", playkibana)
	http.HandleFunc("/elastic", playelastic)
	//http.HandleFunc("/info", info)
	//http.HandleFunc("/cfinfo", util.CfInfo)
	log.Printf("Listening at %s", port)
	rlog.Info(fmt.Sprintf("Listening at %s", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

type Tweet struct {
	User     string `json:"user"`
	Message  string `json:"message"`
	Location string `json:"location"`
	IP       string `json:"ip"`
}

func feederRandomJson() {
	t := Tweet{
		User:     randomdata.SillyName(),
		Message:  randomdata.Paragraph(),
		Location: randomdata.Country(randomdata.FullCountry),
		IP:       randomdata.IpV4Address(),
	}
	j, _ := json.Marshal(t)
	fmt.Println(string(j))
}

func longJob() {
	for {
		time.Sleep(100 * time.Millisecond)
		feederRandomJson()
	}
}

func playelk(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello at swisscom elk")
}

func getHttp(url string) (int, []byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, []byte{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, []byte{}, err
	}
	return resp.StatusCode, body, nil
}

func playelastic(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Fprintln(w, string(path))
	_, b, _ := getHttp(fmt.Sprintf("%s/%s", app.elastic, "_cat/indices?v"))
	//fmt.Fprintln(w, "cf service-connector --skip-ssl-validation   9090 %s", app.elastic)
	fmt.Fprintln(w, string(b))
}

func playkibana(w http.ResponseWriter, req *http.Request) {
	log.Printf("play kibana, play")
	fmt.Fprintln(w, "hello swisscom kibana!")
	fmt.Fprintln(w, "cf service-connector --skip-ssl-validation   8090 %s", app.kibana)
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

	r += app.String()
	fmt.Fprintln(w, r)

	return
}
