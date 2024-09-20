package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

var startTime time.Time

type configData struct {
	Listenport int
}

type healthData struct {
	Uptime   time.Duration `json:"uptime"`
	Hostname string        `json:"hostname"`
}

type resultResponse struct {
	Result string `json:"result"`
}

func init() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func health(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Fprintf(w, "{\"hostname\":\"%s\"}\n", "undefined")
	}
	data := &healthData{Uptime: uptime(), Hostname: hostname}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	fmt.Fprintf(w, "%s\n", jsonData)
	// fmt.Fprintf(w, "{\n\"uptime\":\"%s\"\n\"hostname\":\"%s\"\n}\n", uptime(), hostname)
}

func hello(w http.ResponseWriter, req *http.Request) {
	formatresp(w, "Hello")
	// fmt.Fprintf(w, "hello\n")
}

func world(w http.ResponseWriter, req *http.Request) {
	formatresp(w, "Hello world!")
	// fmt.Fprintf(w, "Hello world!\n")
}

func hi(w http.ResponseWriter, req *http.Request) {
	formatresp(w, "Hi")
	// fmt.Fprintf(w, "Hi\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func formatresp(w http.ResponseWriter, result string) {
	data := &resultResponse{Result: result}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	fmt.Fprintf(w, "%s\n", jsonData)
}

func defaultresp(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Well then hello, %q", html.EscapeString(req.URL.Path))
}

func readConf(filename string) (*configData, error) {
	c := &configData{}
	// read config file
	yamlfile, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("error: %v\n", err)
		return c, err
	}

	err = yaml.Unmarshal(yamlfile, c)
	if err != nil {
		log.Printf("error: %v\n", err)
		return c, err
	}

	return c, nil
}

func main() {

	c, err := readConf("config.yaml")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	if c.Listenport == 0 {
		log.Fatalf("Listenport not defined!\n")
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/health", health)
	http.HandleFunc("/", defaultresp)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Listenport), nil))
}
