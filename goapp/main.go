package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handleRequest(resp http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/api/short" {
		if req.Method != http.MethodPost {
			resp.Header().Add("Allow", "POST")
			AnswerWithError(resp, 405, "Method Not Allowed")
			return
		}

		contentType := req.Header.Get("Content-Type")
		if contentType != "application/json" {
			AnswerWithError(resp, 415, "Unsupported Media Type")
			return
		}

		HandleRequestCreateShortUrl(req, resp)
		return
	}

	if strings.HasPrefix(req.RequestURI, "/api/short/") {

		urlParts := strings.Split(req.RequestURI, "/")
		if len(urlParts) != 4 {
			AnswerWithError(resp, 404, "Not found")
			return
		}

		if req.Method != http.MethodGet {
			resp.Header().Add("Allow", "GET")
			AnswerWithError(resp, 405, "Method Not Allowed")
			return
		}

		shortStr := urlParts[3]
		if len(shortStr) < 1 || len(shortStr) > 8 {
			AnswerWithError(resp, 400, "Wrong short string length")
			return
		}

		HandleRequestGetUrlByShortString(resp, shortStr)
		return
	}

	AnswerWithError(resp, 404, "404 Not Found")
}

func main() {
	fmt.Println("Initing Redis client")
	InitRedis()
	fmt.Println("Redis client was inited")

	// addr := "localhost:11200"
	addr := "0.0.0.0:11200"
	http.HandleFunc("/", handleRequest)

	fmt.Println("Starting listening: " + addr)

	err := http.ListenAndServe(addr, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
