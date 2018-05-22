// A file containing the log initializtion elements. All logger configuration is self contained in this file.

// Off -> Trace -> Debug -> Info -> Warn -> Error -> Fatal

// Set level to Fatal will enable all levels.
// Set level to Info will disable Warn, Error, and Fatal only.
// Set level to Off will disable all.

// The InitLogger has to be called from main() function,
// to initialize the logger.
// e.g. it can be initialized as,
// DISABLE Logging:
// DisableLogger()
// or
// InitLogger(ioutil.Discard,   // conveniently discard Trace
//            ioutil.Discard,   // Debug
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

// log file
var file io.Writer
var logLevel int

const (
	dirName  = "logs"
	fileName = "main.log"

	OFF = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var (
	Trace *log.Logger
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Fatal *log.Logger
)

func InitLogger(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	fatalHandle io.Writer) {

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
	SetLogLevel(FATAL) // enable all by default
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

	SetLogLevel(FATAL) // enable all by default
}

func DisableLogger() {
	SetLogLevel(OFF)
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
	case OFF:
		InitLogger(discard, discard, discard, discard, discard, discard)
	case TRACE:
		InitLogger(output, discard, discard, discard, discard, discard)
	case DEBUG:
		InitLogger(output, output, discard, discard, discard, discard)
	case INFO:
		InitLogger(output, output, output, discard, discard, discard)
	case WARN:
		InitLogger(output, output, output, output, discard, discard)
	case ERROR:
		InitLogger(output, output, output, output, output, discard)
	case FATAL:
		InitLogger(output, output, output, output, output, output)
	}
}
