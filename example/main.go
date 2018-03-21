package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/njuettner/alexa"
)

func alexaDispatchIntentHandler(req alexa.Request) (*alexa.Response, error) {
	switch req.RequestBody.Intent.Name {
	case "hello":
		return alexaGreeterIntentHandler(req)
	default:
		return alexaHelpHandler()
	}
}

func main() {
	lambda.Start(alexaDispatchIntentHandler)
}

func alexaGreeterIntentHandler(req alexa.Request) (*alexa.Response, error) {
	var output string
	switch req.RequestBody.Locale {
	case alexa.LocaleGerman:
		output = "Hallo"
	case alexa.LocaleJapanese:
		output = "こんにちは"
	default:
		output = "Hello"
	}

	simpleResponse := &alexa.SimpleResponse{
		OutputSpeechText: output,
		CardTitle:        "Greeter",
		CardContent:      "Greeter Content",
	}

	//standardResponse := &alexa.StandardResponse{
	//	OutputSpeechText: "Hello World",
	//	CardTitle:        "Title",
	//	CardText:         "Card Text",
	//	CardImage: struct {
	//		SmallImageUrl string
	//		LargeImageUrl string
	//	}{
	//		SmallImageUrl: "https:/foobar.com/small_image.jpg",
	//		LargeImageUrl: "https:/foobar.com/large_image.jpg",
	//	},
	//}

	return alexa.NewResponse(simpleResponse), nil
}

func alexaHelpHandler() (*alexa.Response, error) {
	helpResponse := &alexa.SimpleResponse{
		OutputSpeechText: "Please ask hello to get greetings from your Alexa",
		CardTitle:        "Help for Greeter",
		CardContent:      "Card Content",
	}
	return alexa.NewResponse(helpResponse), nil
}
