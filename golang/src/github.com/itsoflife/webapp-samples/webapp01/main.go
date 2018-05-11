// A basic hello world webapp.
// Run by: go run main.go
package main

import (
  "fmt"           //to print to the client response
  "net/http"      //library for http based interaction
  "log"           //GOOD FOR DEBUGGING
)


// function to handel the top level request
func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Gopher!")
}


func main() {
  server := http.Server {
    Addr: "0.0.0.0:9090",
  }

  log.Println("AD: Starting Server at address: ", server)

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
    log.Fatal("AD: Server Error: ", err)
  }
}
