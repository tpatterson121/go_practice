// Server2 is a minimal "echo" and counter server.

package chapter1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func Webserver2() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// handler echoes the HTTP request

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {   // could've written parseform as err := r.ParseForm() if err != nil {log.Print(error)}
		log.Print(err)                      // combining the statements is shorter and reduces the scope of var err (scope is section 2.7)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v[])
	}
}

// This uses the fields of the http.Request struct to produce output like this:
// GET /?q=query HTTP/1.1
// Header["Accept-Encoding"] = ["gzip, deflate, sdch"]
// Header["Accept-Language"] = ["en-US,en;q=0.8"]
// Header["Connection"] = ["keep-alive"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;..."]
// Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5)..."]
// Host = "localhost:8000"
// RemoteAddr = "127.0.0.1:59911"
// Form["q"] = ["query"]