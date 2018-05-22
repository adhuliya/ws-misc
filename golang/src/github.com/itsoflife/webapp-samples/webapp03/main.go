// A basic hello world webapp.
// Run by: go run main.go logconfig.go
package main

import (
	"fmt"      //to print to the client response
	"net/http" //library for http based interaction

	"app/logs"
)

// function to handel the top level request
func index(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Trace("Received request: ", r) // logging

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "Hello Gopher! <a href='static/static.html'>A Static Page</a>")
}

func main() {
	// defined in ./vendor/app/logs/logconfig.go
	// GOOD FOR DEBUGGING
	logs.InitLogger("configs/seelog.xml")
	defer logs.Logger.Flush()

	//handler for static files
	fs := http.FileServer(http.Dir("static"))

	// route top level request to `index` function.
	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// configure server
	server := http.Server{
		Addr: "0.0.0.0:9090",
	}

	logs.Logger.Info("STARTING SERVER : ", server)

	// run the server and start listening
	err := server.ListenAndServe()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		logs.Logger.Error("Server Error: ", err)
	}
}
