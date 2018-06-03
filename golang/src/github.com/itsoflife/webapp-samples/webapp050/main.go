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
	"app/model"
	"app/route"
	"app/template"
	"app/util"
)

func main() {
	// defined in ./vendor/app/logs/logconfig.go
	// GOOD FOR DEBUGGING & MAINTENANCE
	// STEP 0
	logs.Init("configs/seelog.xml") // as early as possible
	logs.Logger.Info("PROCESS STARTED.")
	logs.Logger.Flush()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	logs.Logger.Trace("Server kill wait Duration: ", wait)

	// STEP 1
	model.Init(&model.DbConf{
		Db:       "postgres",
		UserName: "hop",
		Password: "anshuisneo",
		DbName:   "hop",
		Host:     "127.0.0.1",
		Port:     "5432",
		SslMode:  "disable",
	})

	// STEP 2: setup templates
	template.Init("template/layout/", "template/")

	// STEP 3: setup controllers
	controller.Init()

	// STEP 4: setup routes
	r := route.Init()

	// SETP 5: setup server
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

	// STEP 6: setup graceful shutdown
	// Prepare the server for a graceful shutdown.
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

	logs.Logger.Warn("SERVER STOPPED.")
	logs.Logger.Flush()
	os.Exit(0)
}
