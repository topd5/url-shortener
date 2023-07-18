package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func HandleRequestCreateShortUrl(req *http.Request, resp http.ResponseWriter) {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		AnswerWithError(resp, 500, "Internal Server Error")
		return
	}

	var decodedPayload interface{}

	err = json.Unmarshal(bodyBytes, &decodedPayload)
	if err != nil {
		AnswerWithError(resp, 400, "Request payload must be a valid json data")
		return
	}

	payloadMap, ok := decodedPayload.(map[string]any)
	if !ok {
		AnswerWithError(resp, 400, "Request payload must be a valid json object")
		return
	}

	mayBeUrlString := payloadMap["url"]

	if mayBeUrlString == nil {
		AnswerWithError(resp, 400, "`url` field must exist and not be null")
		return
	}

	urlToShorten, ok := mayBeUrlString.(string)
	if !ok {
		AnswerWithError(resp, 400, "`url` value must be a string")
		return
	}

	if !(strings.HasPrefix(urlToShorten, "http://") || strings.HasPrefix(urlToShorten, "https://")) {
		AnswerWithError(resp, 400, "`url` value must start with http:// or https://")
		return
	}

	// Checking for spaces in url
	if strings.Contains(urlToShorten, " ") {
		AnswerWithError(resp, 400, "`url` value is not a valid url")
		return
	}

	// url.Parse() is a bad validator, but there is still no better one
	_, err = url.Parse(urlToShorten)
	if err != nil {
		AnswerWithError(resp, 400, "`url` value is not a valid url")
		return
	}

	shortName := GenerateRandomString()

	err = RedisSet(shortName, urlToShorten)
	if err != nil {
		AnswerWithError(resp, 500, "Internal Server Error")
		return
	}

	AnswerWithCreatedShort(resp, urlToShorten, shortName)
}
