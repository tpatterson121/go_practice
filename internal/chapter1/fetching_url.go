// Fetch prints the content found at a URL.

package chapter1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetchurl() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", resp.Status)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("$s", b)
	}
}


// RESULT EXAMPLE

// $ go build gopl.io/ch1/fetch
// $ ./fetch http://gopl.io
// <html>
// <head>
// <title>The Go Programming Language</title>
// ...
// If the HTTP request fails, fetch reports the failure instead:
//  $ ./fetch http://bad.gopl.io
// fetch: Get http://bad.gopl.io: dial tcp: lookup bad.gopl.io: no such host
// in either error case, Os.Exit(1) causes the process to exit with a status code of 1