package controller

import (
	"net/http"
	"story/model"
)

func MyHandlerFunc(storyArcs model.StoryArcs) {
	myArc := storyArcs["intro"]
	http.Handle("/", myArc)
	for k, v := range storyArcs {
		http.Handle("/"+k, v)
	}
	http.ListenAndServe(":7007", nil)
}
