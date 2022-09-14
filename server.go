package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

var startTime time.Time

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

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/health", health)
	http.HandleFunc("/", defaultresp)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
