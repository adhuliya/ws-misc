// A basic hello world webapp.
// Run by: go run main.go
package main

import (
	"fmt"      //to print to the client response
	"net/http" //library for http based interaction
)

// function to handel the top level request
func index(w http.ResponseWriter, r *http.Request) {
	// Content-Type needs to be set because net/http package
	// tries to determine the content type using file extension,
	// and we are directly writing to the ResponseWriter using fmt.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "Hello Gopher! <a href='static/static.html'>A Static Page</a>")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:9090",
	}

	//handler for static files
	fs := http.FileServer(http.Dir("static"))

	// route top level request to `index` function.
	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// run the server and start listening
	err := server.ListenAndServe()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error", err)
	}
}
