// A basic hello world webapp.
// Run by: go run main.go logconfig.go
package main

import (
    "fmt"           //to print to the client response
    "net/http"      //library for http based interaction

    "io/ioutil"
    "os"
)


// function to handel the top level request
func index(w http.ResponseWriter, r *http.Request) {
    Info.Println("Received request: ", r) // logging
    fmt.Fprintf(w, "Hello Gopher!")
}


func main() {
    // defined in ./logconfig.go
    // GOOD FOR DEBUGGING
    InitLogger(ioutil.Discard,   // conveniently discard Trace
               os.Stdout,        // Info to os.Stdout
               os.Stdout,        // Warn to os.Stdout
               os.Stderr,        // Error to os.Stderr
               os.Stderr)        // Fatal to os.Stderr

    server := http.Server {
        Addr: "0.0.0.0:9090",
    }

    Info.Println("AD: Starting Server at address: ", server)

    // route top level request to `index` function.
    http.HandleFunc("/", index)

    // run the server and start listening
    err := server.ListenAndServe()
    checkError(err)
}


func checkError(err error) {
    if err != nil {
        Fatal.Println("Server Error: ", err)
    }
}
