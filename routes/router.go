package routes

import (
	"net/http"

	"../controllers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", controllers.WebhookHandler)
	return mux
}