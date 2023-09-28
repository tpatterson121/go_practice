// Server1 is a minimal "echo" server.

package chapter1

import (
	"fmt"
	"log"
	"net/http"
)

func Webserver1() {
	http.HandleFunc("/", handler) //each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(w, "URL.Path = %q\n", r.URL.Path) //ask brandon about w
}
