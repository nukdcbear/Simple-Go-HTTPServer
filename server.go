package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"strconv"
)

var startTime time.Time

type configData struct {
	Listenport int
}

func init() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"uptime\":\"%s\"}\n", uptime())
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func world(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func defaultresp(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Well then hello, %q", html.EscapeString(req.URL.Path))
}

func readConf(filename string) (*configData, error) {
	c := &configData{}
	// read config file
	yamlfile, err := ioutil.ReadFile(filename)
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

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(c.Listenport), nil))
}
