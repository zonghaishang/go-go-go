package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	args := os.Args[1:]

	fmt.Fprintf(os.Stdout, "arg[0] = %v, args = %v\n", os.Args[0], args)

	handler := ServeHTTP
	http.HandleFunc("/", handler)

	log.Fatalln(http.ListenAndServe("localhost:8000", nil))

}

func ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	fmt.Fprintf(w, "path = %s\n", path)
}
