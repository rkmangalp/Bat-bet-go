package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config holds the configuration settings for the application
type Config struct {
	Port     string `json:"port"`
	Database struct {
		URL      string `json:"url"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

// AppConfig is the global configuration variable
var AppConfig *Config

// LoadConfig loads the configuration from the specified file
func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	AppConfig = &Config{}
	err = decoder.Decode(AppConfig)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
