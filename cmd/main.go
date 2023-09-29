package main

import (
	"github.com/tpatterson121/go_practice/internal/chapter1"
)

func main() {
	// myfile, err := os.Create("AnimatedGifs.gif")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(myfile)
	// defer myfile.Close()
	// fmt.Println("Hello world")
	// // chapter1.Echo1()
	// // chapter1.Echo2()
	// // chapter1.FindDuplicates()
	// // chapter1.Lissa()
	// chapter1.Lissajous(myfile)
	// chapter1.Fetchurl()
	// chapter1.Fetchmultiple()
	//chapter1.Webserver1()
	chapter1.Webserver2()

}
