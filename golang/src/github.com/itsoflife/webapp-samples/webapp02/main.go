// A basic hello world webapp.
// Run by: go run main.go logconfig.go
package main

import (
	"fmt"      //to print to the client response
	"net/http" //library for http based interaction
	//"io/ioutil"
	//"os"
)

// function to handel the top level request
func index(w http.ResponseWriter, r *http.Request) {
	Info.Println("Received request: ", r) // logging

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "Hello Gopher! <a href='static/static.html'>Static Page</a> <a href='static/'>Static Content</a>")
}

func main() {
	// defined in ./logconfig.go
	// GOOD FOR DEBUGGING
	//EnableLogger()
	EnableFileLogger()
	SetLogLevel(FATAL) // see logconfigv2.go file

	// To disable logging,
	//DisableLogger()

	server := http.Server{
		Addr: "0.0.0.0:9090",
	}

	Info.Println("AD: Starting Server at address: ", server)

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
		Error.Println("Server Error: ", err)
	}
}
