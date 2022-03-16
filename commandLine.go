package main

import "story/controller"

func main() {
	storyArcs := controller.ReadFromJson("docs/input.json")
	storyArcs.CLIStoryTell(storyArcs["intro"], true)
}
