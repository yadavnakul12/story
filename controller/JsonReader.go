package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"story/model"
)

func ReadFromJson(fileName string) model.StoryArcs {
	fileContent, readError := ioutil.ReadFile(fileName)
	if readError != nil {
		fmt.Println(readError.Error())
		panic(readError)
	}
	var storyArcs model.StoryArcs
	parseError := json.Unmarshal(fileContent, &storyArcs)
	if parseError != nil {
		log.Fatal(parseError.Error())
	}
	return storyArcs
}
