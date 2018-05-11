// A file containing the log initializtion elements.
// The InitLogger has to be called from main() function,
// to initialize the logger.
// e.g. 
// InitLogger(ioutil.Discard,   // conveniently discard Trace
//            os.Stdout,        // Info to os.Stdout
//            os.Stdout,        // Warn to os.Stdout
//            os.Stderr,        // Error to os.Stderr
//            os.Stderr)        // Fatal to os.Stderr
package main

import (
    "io"
//    "io/ioutil"
    "log"
//    "os"
)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warn    *log.Logger
    Error   *log.Logger
    Fatal   *log.Logger
)

func InitLogger(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer,
    fatalHandle io.Writer) {

    Trace = log.New(traceHandle,
        "TRACE  : ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO   : ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warn    = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR  : ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Fatal = log.New(fatalHandle,
        "FATAL  : ",
        log.Ldate|log.Ltime|log.Lshortfile)
}


