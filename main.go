//**Created by Kyle Prouty
//**Winter Break 2016

package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.template"))
}

func main() {
	http.HandleFunc("/", idx)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("pub/"))))
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.template", nil)
	fmt.Println(req.URL.Path)
}