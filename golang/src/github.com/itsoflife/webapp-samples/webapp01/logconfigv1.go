// All logging related code is in one file for convenience.
// This simplest form of logging is for the most basic projects.
// To use simple initialize the logger:
// StartLogger()   // initializes the logger
// Log.Println("Log this message") // using the logger
// Log.Fatal("Log this message") // using the logger
// StopLogger()  // stops the logger output from showing

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// global logger variable: for very simple projects
var Log *log.Logger

// Initialize a logger to the given output stream
func InitLogger(w io.Writer) {
	Log = log.New(w,
		"LOG  : ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// convenience function over InitLogger(io.Writer)
func StartLogger() {
	InitLogger(os.Stderr)
}

// stops logger by redirecting output to ioutil.Discard
func StopLogger(logger *log.Logger) {
	if Log != nil {
		Log.SetOutput(ioutil.Discard)
	}
}
