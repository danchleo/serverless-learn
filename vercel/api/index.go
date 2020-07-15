package api

import (
	"fmt"
	"net/http"
	"time"
	"serverless_learn/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime + handlers.Test())
}
