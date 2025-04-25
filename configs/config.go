package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Port string `json:"port"`
}

func LoadConfig(env string) Config {
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