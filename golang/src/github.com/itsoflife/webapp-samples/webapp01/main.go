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
	// always initialize logger global variable as early as possible
	StartLogger()

	// to disable logging, uncomment the line below
	//StopLogger()

	server := http.Server{
		Addr: "0.0.0.0:9090",
	}

	Log.Println("AD: Starting Server at address: ", server)

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
		Log.Fatal("AD: Server Error: ", err)
	}
}
