// All logging related code is in one file for convenience.
// This simplest form of logging is for the most basic projects.

// To use, initialize the logger:
// EnableLogger()   // initializes the logger
// Log.Println("Log this message") // using the logger
// Log.Fatal("Log this message") // using the logger
// DisableLogger()  // disables the logger output from showing

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// global logger variable: for very simple projects
var Log *log.Logger
var file io.Writer

const (
	dirName  = "logs"
	fileName = "main.log"
)

// Initialize a logger to the given output stream
// Use EnalbeLogger() instead for convenience.
func InitLogger(w io.Writer) {
	Log = log.New(w,
		"LOG  : ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// convenience function over InitLogger(io.Writer)
func EnableLogger() {
	InitLogger(os.Stderr)
}

// Log to a preset log file
func EnableFileLogger() {
	if file == nil {
		// attempt to create the directory and ignore errors
		_ = os.Mkdir(dirName, 0755)

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(dirName+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		file = f // store in global variable for record
	}

	InitLogger(file)
}

// stops logger by redirecting output to ioutil.Discard
func DisableLogger(logger *log.Logger) {
	if Log != nil {
		Log.SetOutput(ioutil.Discard)
	}
}
