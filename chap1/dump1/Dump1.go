package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	for input.Scan() {
		fmt.Println("Received:" + input.Text())
		if len(input.Text()) <= 0 {
			break
		}
		words[input.Text()]++
	}

	for word, n := range words {
		if n > 1 {
			fmt.Printf("%d %s\n", n, word)
		}
	}
}
