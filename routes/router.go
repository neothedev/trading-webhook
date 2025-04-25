package routes

import (
	"net/http"

	"github.com/neothedev/trading-webhook/controllers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", controllers.WebhookHandler)
	return mux
}