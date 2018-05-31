// A basic hello world webapp.
// Run by: go run main.go
package main

import (
	"context"
	"flag" // for command line arguments
	"fmt"
	"net/http" //library for http based interaction
	"os"
	"os/signal"
	"time"

	"app/controller"
	"app/logs"
	"app/route"
	"app/template"
	"app/util"
)

func main() {
	// defined in ./vendor/app/logs/logconfig.go
	// GOOD FOR DEBUGGING
	logs.Init("configs/seelog.xml")
	logs.Logger.Info("PROCESS STARTED.")
	logs.Logger.Flush()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	logs.Logger.Trace("Server kill wait Duration: ", wait)

	template.Init("template/layout/", "template/")

	controller.Init()

	r := route.Init()

	// configure server
	srv := &http.Server{
		Addr: "0.0.0.0:9090",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logs.Logger.Info("STARTING SERVER : ", srv)
		// some meaningful output at command line
		fmt.Println("SERVING AT: ", srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			util.FatalError(err)
		}
	}()

	// prepare the server for a graceful shutdown.
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	logs.Logger.Warn("USER INTERRUPT! SHUTTING DOWN.")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	logs.Logger.Error("SERVER STOPPED.")
	logs.Logger.Flush()
	os.Exit(0)
}
