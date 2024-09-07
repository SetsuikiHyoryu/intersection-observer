package handler

import (
	"fmt"
	"net/http"
)

func PingHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(response, "pong")
}
