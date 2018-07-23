package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

/**
 * go run Fetch.go http://www.baidu.com
 */
func main()  {

	urls := os.Args[1:]
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
		resp.Body.Close()
		fmt.Printf("%s", b)
	}
}
