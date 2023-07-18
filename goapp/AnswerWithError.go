package main

import (
	"encoding/json"
	"net/http"
)

type ErrorContentAnswer struct {
	Message string `json:"message"`
}

func AnswerWithError(
	resp http.ResponseWriter,
	httpCode int,
	message string,
) {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(httpCode)

	var contentAnswer ErrorContentAnswer
	contentAnswer.Message = message
	jsonedData, _ := json.Marshal(contentAnswer)

	resp.Write(jsonedData)
}
