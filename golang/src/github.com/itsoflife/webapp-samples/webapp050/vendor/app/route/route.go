package route

import (
	"net/http"

	"app/controller"
)

// intialize all routes
func Init() *http.ServeMux {
	r := http.NewServeMux()

	// handler for static files
	fs := http.FileServer(http.Dir("static"))

	// route top level request to `index` function.
	r.HandleFunc("/", controller.Index)
	r.HandleFunc("/aboutme", controller.AboutMe)
	r.HandleFunc("/skillset", controller.SkillSetH)
	r.HandleFunc("/time", controller.TimeApi)
	r.HandleFunc("/timer", controller.Timer)
	r.HandleFunc("/users", controller.Users)
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	return r
}
