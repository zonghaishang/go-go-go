package main

import (
	"os"
	"fmt"
)

func main()  {

	s, step := "", " "

	for _, value := range os.Args[1:] {
		s += step + value
	}

	fmt.Println(s)

}
