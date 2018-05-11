// A basic hello world webapp.
// Run by: go run main.go logconfig.go
package main

import (
    "net/http"      //library for http based interaction

    "app/logs"
    "app/tmpl"
    "app/util"
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

func main() {
  // defined in ./vendor/app/logs/logconfig.go
  // GOOD FOR DEBUGGING
  logs.InitLogger("configs/seelog.xml")
  defer logs.Logger.Flush()
  tmpl.InitTemplates("template/layout/", "template/")

  //handler for static files
  fs := http.FileServer(http.Dir("static"))

  // route top level request to `index` function.
	http.HandleFunc("/", index)
	http.HandleFunc("/aboutme", aboutMe)
  http.HandleFunc("/skillset", skillSet)
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // configure server
  server := http.Server {
      Addr: "0.0.0.0:9090",
  }

  logs.Logger.Info("Starting Server with config: ", server)

  // run the server to start listening
  err := server.ListenAndServe()
  util.CheckError(err)
}


