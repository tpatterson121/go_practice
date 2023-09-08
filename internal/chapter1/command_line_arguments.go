package chapter1

import (
	"fmt"
	"os"
)

func Echo() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "*"
	}
	fmt.Println(s)
}
