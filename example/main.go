package main

import (
	"encoding/json"
	"os"

	"github.com/njuettner/alexa"
)

func main() {
	simpleResponse := &alexa.SimpleResponse{
		OutputSpeechText: "Hello World",
		CardTitle:        "Title",
		CardContent:      "Card Content",
	}

	standardResponse := &alexa.StandardResponse{
		OutputSpeechText: "Hello World",
		CardTitle:        "Title",
		CardText:         "Card Text",
		CardImage: struct {
			SmallImageUrl string
			LargeImageUrl string
		}{
			SmallImageUrl: "https:/foobar.com/small_image.jpg",
			LargeImageUrl: "https:/foobar.com/large_image.jpg",
		},
	}

	resp := alexa.NewResponse(simpleResponse)
	json.NewEncoder(os.Stdout).Encode(resp)
	resp = alexa.NewResponse(standardResponse)
	json.NewEncoder(os.Stdout).Encode(resp)
}
