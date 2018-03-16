package alexa

type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	ResponseBody      responseBody           `json:"response"`
}

type responseBody struct {
	OutputSpeech     outputSpeech  `json:"outputSpeech,omitempty"`
	Card             card          `json:"card,omitempty"`
	Reprompt         outputSpeech  `json:"reprompt,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession,omitempty"`
	Directives       []interface{} `json:"directives,omitempty"`
}

type outputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

type card struct {
	Type    string    `json:"type"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Text    string    `json:"text,omitempty"`
	Image   cardImage `json:"image,omitempty"`
}

type cardImage struct {
	SmallImageUrl string `json:"smallImageUrl,omitempty"`
	LargeImageUrl string `json:"largeImageUrl,omitempty"`
}

type responder interface {
	newResponse() *Response
}

type SimpleResponse struct {
	OutputSpeechText string
	CardTitle        string
	CardContent      string
}

type StandardResponse struct {
	OutputSpeechText string
	CardTitle        string
	CardText         string
	CardImage        struct {
		SmallImageUrl string
		LargeImageUrl string
	}
}

type LinkAccountResponse struct {
	OutputSpeechText string
}

func (res *SimpleResponse) newResponse() *Response {
	return &Response{
		Version: "1.0",
		ResponseBody: responseBody{
			OutputSpeech: outputSpeech{
				Type: "PlainText",
				Text: res.OutputSpeechText,
			},
			Card: card{
				Type:    "Simple",
				Title:   res.CardTitle,
				Content: res.CardContent,
			},
			ShouldEndSession: true,
		},
	}
}

func (res *StandardResponse) newResponse() *Response {
	return &Response{
		Version: "1.0",
		ResponseBody: responseBody{
			OutputSpeech: outputSpeech{
				Type: "PlainText",
				Text: res.OutputSpeechText,
			},
			Card: card{
				Type:  "Standard",
				Title: res.CardTitle,
				Text:  res.CardText,
				Image: cardImage{
					SmallImageUrl: res.CardImage.SmallImageUrl,
					LargeImageUrl: res.CardImage.LargeImageUrl,
				},
			},
			ShouldEndSession: true,
		},
	}
}

func (res *LinkAccountResponse) newResponse() *Response {
	return &Response{
		Version: "1.0",
		ResponseBody: responseBody{
			OutputSpeech: outputSpeech{
				Type: "PlainText",
				Text: res.OutputSpeechText,
			},
			Card: card{
				Type: "LinkAccount",
			},
			ShouldEndSession: true,
		},
	}
}

func NewResponse(r responder) *Response {
	return r.newResponse()
}
