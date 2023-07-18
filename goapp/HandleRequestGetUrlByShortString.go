package main

import (
	"net/http"

	"github.com/go-redis/redis/v8"
)

func HandleRequestGetUrlByShortString(resp http.ResponseWriter, shortStr string) {
	url, err := RedisGet(shortStr)
	if err == redis.Nil {
		AnswerWithError(resp, 404, "Not found")
		return
	} else if err != nil {
		AnswerWithError(resp, 500, "Internal Server Error")
		return
	}

	AnswerWithUrlFromShort(resp, url, shortStr)
}
