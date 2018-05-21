// A file containing the log initializtion elements. All logger configuration is self contained in this file.
// The InitLogger has to be called from main() function,
// to initialize the logger.
// e.g. it can be initialized as,
// InitLogger(ioutil.Discard,   // conveniently discard Trace
//            os.Stdout,        // Info to os.Stdout
//            os.Stdout,        // Warn to os.Stdout
//            os.Stderr,        // Error to os.Stderr
//            os.Stderr)        // Fatal to os.Stderr
// DISABLE Logging:
// InitLogger(ioutil.Discard,   // conveniently discard Trace
//            ioutil.Discard,   // Info
//            ioutil.Discard,   // Warn
//            ioutil.Discard,   // Error
//            ioutil.Discard)   // Fatal
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug *log.Logger
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Fatal *log.Logger
)

func InitLogger(
	debugHandle io.Writer,
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	fatalHandle io.Writer) {

	Debug = log.New(debugHandle,
		"DEBUG  : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Trace = log.New(traceHandle,
		"TRACE  : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO   : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR  : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Fatal = log.New(fatalHandle,
		"FATAL  : ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func StartLogger() {
	InitLogger(ioutil.Discard, // conveniently discard Trace
		os.Stdout, // Debug to os.Stdout
		os.Stdout, // Info to os.Stdout
		os.Stdout, // Warn to os.Stdout
		os.Stderr, // Error to os.Stderr
		os.Stderr) // Fatal to os.Stderr
}

// stops logging by redirecting input to ioutil.Discard
func StopLogger() {
	Debug = log.New(ioutil.Discard, "", 0)
	Trace = log.New(ioutil.Discard, "", 0)
	Info = log.New(ioutil.Discard, "", 0)
	Warn = log.New(ioutil.Discard, "", 0)
	Error = log.New(ioutil.Discard, "", 0)
	Fatal = log.New(ioutil.Discard, "", 0)
}
