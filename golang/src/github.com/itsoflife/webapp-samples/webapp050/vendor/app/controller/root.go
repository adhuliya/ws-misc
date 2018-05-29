package controller

import (
	"net/http" //library for http based interaction

	"app/logs"
	tmpl "app/template"
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

func Index(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Trace("Received request: ", r) // logging
	tmpl.RenderTemplate(w, "index.tmpl", nil)
}

func AboutMe(w http.ResponseWriter, r *http.Request) {
	userData := &UserData{Name: "Asit Dhal", City: "Bhubaneswar", Nationality: "Indian"}
	tmpl.RenderTemplate(w, "aboutme.tmpl", userData)
}

func SkillSetH(w http.ResponseWriter, r *http.Request) {
	skillSets := SkillSets{&SkillSet{Language: "Golang", Level: "Beginner"},
		&SkillSet{Language: "C++", Level: "Advanced"},
		&SkillSet{Language: "Python", Level: "Advanced"}}
	tmpl.RenderTemplate(w, "skillset.tmpl", skillSets)
}

func TimeApi(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Trace("Time requested.")

	jsonStr := fmt.Sprintf(`{"time":"%s"}`, time.Now())
	jsn, err := json.Marshal(jsonStr)
	util.CheckError(err)

	w.Header().Set("Content-Type", "text/json; charset=utf-8")

	fmt.Fprintf(w, string(jsn))
}

func Timer(w http.ResponseWriter, r *http.Request) {
	tmpl.RenderTemplate(w, "timer.tmpl", SkillSets{})
}
