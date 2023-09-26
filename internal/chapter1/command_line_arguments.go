package main // echo1 prints its command-line arguments

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// echo2 prints its command-line arguments

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "*"
	}
	fmt.Println(s)
}

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}