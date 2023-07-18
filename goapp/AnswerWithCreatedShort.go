package main

import (
	"encoding/json"
	"net/http"
)

type CreatedShortContentAnswer struct {
	Url   string `json:"url"`
	Short string `json:"short"`
}

func AnswerWithCreatedShort(
	resp http.ResponseWriter,
	urlToShorten string,
	shortString string,
) {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(201)

	var contentAnswer CreatedShortContentAnswer
	contentAnswer.Short = shortString
	contentAnswer.Url = urlToShorten
	jsonedData, _ := json.Marshal(contentAnswer)

	resp.Write(jsonedData)
}
