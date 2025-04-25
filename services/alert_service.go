package services

import (
	"fmt"
	"log"
	"time"

	"../models"
)

func ProcessAlert(payload models.AlertPayload) (string, error) {
	t := time.UnixMilli(payload.BarTime)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return "", fmt.Errorf("Could not load time zone")
	}
	tInEST := t.In(loc)
	formattedTime := tInEST.Format("2006-01-02 03:04:05.000 PM MST")

	log.Printf("Received alert:\nTicker: %s\nAlert: %s\nLead: %.2f\nLag: %.2f\nClose: %.2f\nTime: %s\n",
		payload.Ticker, payload.Alert,
		payload.NeoCloud.Lead, payload.NeoCloud.Lag,
		payload.Ohlcv.Close, formattedTime,
	)

	return formattedTime, nil
}