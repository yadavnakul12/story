package model

import (
	"html/template"
	"net/http"
)

type Arc struct {
	Title   string
	Story   []string
	Options []Option
}

type Option struct {
	Text, Arc string
}

type StoryArcs map[string]Arc

func (arc Arc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	myTemplate := template.Must(template.ParseFiles("view/arc.html"))
	myTemplate.Execute(w, arc)
}
