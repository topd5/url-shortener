package main

import (
	"encoding/json"
	"net/http"
)

type UrlFromShortContentAnswer struct {
	Short string `json:"short"`
	Url   string `json:"url"`
}

func AnswerWithUrlFromShort(
	resp http.ResponseWriter,
	url string,
	shortString string,
) {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(200)

	var contentAnswer UrlFromShortContentAnswer
	contentAnswer.Url = url
	contentAnswer.Short = shortString
	jsonedData, _ := json.Marshal(contentAnswer)

	resp.Write(jsonedData)
}
