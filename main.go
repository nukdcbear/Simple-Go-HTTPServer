package main

import (
    "fmt"
	"html"
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
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
	http.HandleFunc("/hi", hi)
    http.HandleFunc("/headers", headers)
	http.HandleFunc("/", defaultresp)

    log.Fatal(http.ListenAndServe(":8090", nil))
}