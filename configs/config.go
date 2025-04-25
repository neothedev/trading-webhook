package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// TradeStationEndpoints holds endpoint paths with formatting
type TradeStationEndpoints struct {
	Accounts     string `json:"Accounts"`
	Balances     string `json:"Balances"`
	Positions    string `json:"Positions"`
	Orders       string `json:"Orders"`
	Quote        string `json:"Quote"`
	BarChart     string `json:"BarChart"`
	OptionChain  string `json:"OptionChain"`
	OptionQuotes string `json:"OptionQuotes"`
	MarketDepth  string `json:"MarketDepth"`
}

// Config holds the entire app configuration
type Config struct {
	Port                   string                `json:"port"`
	LogLevel               string                `json:"logLevel"`
	Timezone               string                `json:"timezone"`
	WebhookPath            string                `json:"webhookPath"`
	TradeStationBaseURL    string                `json:"TradeStationBaseURL"`
	TradeStationAPIRoot    string                `json:"TradeStationAPIRoot"`
	TradeStationEndpoints  TradeStationEndpoints `json:"TradeStationEndpoints"`
}

// LoadConfig loads config file based on the environment name
func LoadConfig(env string) Config {
	if env == "" {
		env = "dev" // default fallback
	}
	file := fmt.Sprintf("configs/config.%s.json", env)

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	return config
}
