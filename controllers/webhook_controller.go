package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/neothedev/trading-webhook/models"
	"github.com/neothedev/trading-webhook/services"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload models.AlertPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	formattedTime, err := services.ProcessAlert(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Alert received at %s", formattedTime)
}