package api

import (
	"fmt"
	"net/http"
	"time"
	"github.com/danchleo/serverless-learn/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime + handlers.Test())
}
