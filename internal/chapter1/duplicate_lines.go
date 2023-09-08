package chapter1

import (
	"bufio"
	"fmt"
	"os"
)

func FindDuplicates() {
	counts := make(map[string]int)      // counts is a map, whose key is a string and it returns an int
	input := bufio.NewScanner(os.Stdin) // input is the first file provided as a cmdline argument

	for input.Scan() {                  //for is a loop that will keep using scan, as long as scan is true, there is more to be read
		counts[input.Text()]++ 			//given a string, this map becomes the int
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) 
		}
	}
}
