// A file containing the log initializtion elements. All logger configuration is self contained in this file.

// Debug -> Fatal -> Error -> Warn -> Info -> Trace

// Set level to Fatal will disable Debug and enable all the rest.
// Set level to Error will disable Fatal and Debug and enable all the rest.

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
//            ioutil.Discard)   // Debug
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// log file
var file io.Writer
var logLevel int

const (
	dirName  = "logs"
	fileName = "main.log"

	DEBUG = iota
	FATAL
	ERROR
	WARN
	INFO
	TRACE
	DISABLE
)

var (
	Debug *log.Logger
	Fatal *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Trace *log.Logger
)

func InitLogger(
	debugHandle io.Writer,
	fatalHandle io.Writer,
	errorHandle io.Writer,
	warningHandle io.Writer,
	infoHandle io.Writer,
	traceHandle io.Writer) {

	if traceHandle == ioutil.Discard {
		Trace = log.New(ioutil.Discard, "", 0)
	} else {
		Trace = log.New(traceHandle,
			"TRACE  : ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}

	if infoHandle == ioutil.Discard {
		Info = log.New(ioutil.Discard, "", 0)
	} else {
		Info = log.New(infoHandle,
			"INFO   : ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}

	if warningHandle == ioutil.Discard {
		Warn = log.New(ioutil.Discard, "", 0)
	} else {
		Warn = log.New(warningHandle,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}

	if errorHandle == ioutil.Discard {
		Error = log.New(ioutil.Discard, "", 0)
	} else {
		Error = log.New(errorHandle,
			"ERROR  : ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}

	if fatalHandle == ioutil.Discard {
		Fatal = log.New(ioutil.Discard, "", 0)
	} else {
		Fatal = log.New(fatalHandle,
			"FATAL  : ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}

	if debugHandle == ioutil.Discard {
		Debug = log.New(ioutil.Discard, "", 0)
	} else {
		Debug = log.New(debugHandle,
			"DEBUG  : ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func EnableLogger() {
	SetLogLevel(DEBUG) // enable all by default
}

// Log to a preset log file
// Enables all levels.
func EnableFileLogger() {
	if file == nil {
		// attempt to create the directory and ignore errors
		_ = os.Mkdir(dirName, 0755)

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(dirName+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		file = f // set to the global var to avoid recreating
	}

	SetLogLevel(DEBUG) // enable all by default
}

func DisableLogger() {
	SetLogLevel(DISABLE)
}

func SetLogLevel(level int) {
	logLevel = level

	var discard = ioutil.Discard
	var output io.Writer = nil
	if file == nil {
		output = os.Stdout
	} else {
		output = file
	}

	switch level {
	case DEBUG:
		InitLogger(output, output, output, output, output, output)
	case FATAL:
		InitLogger(discard, output, output, output, output, output)
	case ERROR:
		InitLogger(discard, discard, output, output, output, output)
	case WARN:
		InitLogger(discard, discard, discard, output, output, output)
	case INFO:
		InitLogger(discard, discard, discard, discard, output, output)
	case TRACE:
		InitLogger(discard, discard, discard, discard, discard, output)
	case DISABLE:
		InitLogger(discard, discard, discard, discard, discard, discard)
	}
}
