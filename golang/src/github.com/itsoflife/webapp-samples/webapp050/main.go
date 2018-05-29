// A basic hello world webapp.
// Run by: go run main.go
package main

import (
	"net/http" //library for http based interaction

	"app/logs"
	"app/tmpl"
	"app/util"
	"encoding/json"
	"fmt"
	"time"
)

type UserData struct {
	Name        string
	City        string
	Nationality string
}

type SkillSet struct {
	Language string
	Level    string
}

type SkillSets []*SkillSet

func index(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Trace("Received request: ", r) // logging
	tmpl.RenderTemplate(w, "index.tmpl", nil)
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	userData := &UserData{Name: "Asit Dhal", City: "Bhubaneswar", Nationality: "Indian"}
	tmpl.RenderTemplate(w, "aboutme.tmpl", userData)
}

func skillSet(w http.ResponseWriter, r *http.Request) {
	skillSets := SkillSets{&SkillSet{Language: "Golang", Level: "Beginner"},
		&SkillSet{Language: "C++", Level: "Advanced"},
		&SkillSet{Language: "Python", Level: "Advanced"}}
	tmpl.RenderTemplate(w, "skillset.tmpl", skillSets)
}

var timeMap map[string]string

func timeApi(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Info("Time requested.")

	timeMap["time"] = fmt.Sprintf("%s", time.Now())
	jsn, err := json.Marshal(timeMap)
	util.CheckError(err)

	w.Header().Set("Content-Type", "text/json; charset=utf-8")

	fmt.Fprintf(w, string(jsn))
}

func timer(w http.ResponseWriter, r *http.Request) {
	tmpl.RenderTemplate(w, "timer.tmpl", SkillSets{})
}

func main() {
	// defined in ./vendor/app/logs/logconfig.go
	// GOOD FOR DEBUGGING
	logs.InitLogger("configs/seelog.xml")
	defer logs.Logger.Flush()

	tmpl.InitTemplates("template/layout/", "template/")

	timeMap = make(map[string]string)
	//handler for static files
	fs := http.FileServer(http.Dir("static"))

	// route top level request to `index` function.
	http.HandleFunc("/", index)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/skillset", skillSet)
	http.HandleFunc("/time", timeApi)
	http.HandleFunc("/timer", timer)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// configure server
	server := http.Server{
		Addr: "0.0.0.0:9090",
	}

	logs.Logger.Info("STARTING SERVER : ", server)

	// some meaningful output at command line
	fmt.Println("SERVING AT: ", server.Addr)

	// run the server to start listening
	err := server.ListenAndServe()
	util.FatalError(err)
}
