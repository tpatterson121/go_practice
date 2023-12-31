// Fetchall fetches URLs in parallel and reports their times and sizes.
// go over with brandon

package chapter1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Fetchmultiple() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) //ask brandon about 2 args
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// EXAMPLE OF RESULT

// $ go build gopl.io/ch1/fetchall
// $ ./fetchall https://golang.org http://gopl.io https://godoc.org
// 0.14s 6852 https://godoc.org
// 0.16s 7261 https://golang.org
// 0.48s 2475 http://gopl.io
// 0.48s elapsed
