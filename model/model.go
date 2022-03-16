package model

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
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

func (arc Arc) DisplayStoryOnCLI() {
	fmt.Println("*************************")
	fmt.Println(arc.Title)
	fmt.Println("*************************")
	fmt.Println(strings.Join(arc.Story, ". "))
	if len(arc.Options) > 0 {
		fmt.Printf("\n\nOptions  \n")
		for index, option := range arc.Options {
			fmt.Printf("%v:%v\n", index+1, option.Text)
		}
	}
}

func (storArcs StoryArcs) CLIStoryTell(arc Arc, isDisplayable bool) {
	if isDisplayable {
		arc.DisplayStoryOnCLI()
	}
	if len(arc.Options) > 0 {
		var choice int
		fmt.Printf("\nEnter you choice::")
		fmt.Scan(&choice)
		if choice > 0 && choice <= len(arc.Options) {
			storArcs.CLIStoryTell(storArcs[arc.Options[choice-1].Arc], true)
		} else {
			fmt.Printf("\n\n\nXXXXXXXX Alert XXXXXXXX\n")
			fmt.Println("Enter valid option please")
			storArcs.CLIStoryTell(arc, false)
		}
	} else {
		os.Exit(0)
	}
}
