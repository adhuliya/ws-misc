package tmpl

import (
	"app/logs" // of this project in `vendor/app/logs` folder
	"fmt"
	"github.com/oxtoacart/bpool"
	"html/template"
	"net/http"
	"path/filepath"
)

// Example usage.
//func index(w http.ResponseWriter, r *http.Request) {
//	tmpl.RenderTemplate(w, "index.tmpl", nil)
//}

// This package expects the following directory layout:
// template
// ├── aboutme.tmpl
// ├── index.tmpl
// ├── layout
// │   ├── base.tmpl
// │   ├── js.tmpl
// │   └── style.tmpl
// └── skillset.tmpl

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var templates map[string]*template.Template
var bufpool *bpool.BufferPool

const BUFF_POOL_SIZE = 1 << 6

type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

var templateConfig TemplateConfig

// usage: tmpl.InitTemplates("template/layout/", "template/")
func InitTemplates(layoutPath, includePath string) {
	loadConfiguration(layoutPath, includePath)
	loadTemplates()
}

func loadConfiguration(layoutPath, includePath string) {
	templateConfig.TemplateLayoutPath = layoutPath
	templateConfig.TemplateIncludePath = includePath
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.tmpl")
	if err != nil {
		logs.Logger.Error(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.tmpl")
	if err != nil {
		logs.Logger.Error(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		logs.Logger.Error(err)
	}

	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			logs.Logger.Error(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	logs.Logger.Info("Templates Loading Successful")

	bufpool = bpool.NewBufferPool(BUFF_POOL_SIZE)
	logs.Logger.Info("Buffer Allocation Successful. Size: ", BUFF_POOL_SIZE)
}

// InitTemplates() must be called for initialization
// before calling this function.
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		logs.Logger.Error("Template `", name, "` doesn't exist")
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		logs.Logger.Error("Rendering template `", name,
			"` failed. With data: ", data)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
