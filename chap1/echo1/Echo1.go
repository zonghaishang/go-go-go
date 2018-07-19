package main

import (
	"os"
	"fmt"
)

/**
 * go run chap1/echo1/Echo1.go hello world
 */
func main() {

	var echo, space string
	for i := 1; i < len(os.Args); i++ {
		echo += space + os.Args[i]
		space = " "
	}

	fmt.Println(echo)
}
