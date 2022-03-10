package main

import (
	"story/controller"
)

func main() {
	storyArcs := controller.ReadFromJson("docs/input.json")
	//fmt.Printf("%v\n", storyArcs)
	controller.MyHandlerFunc(storyArcs)
}
