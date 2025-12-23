package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Config struct {
	UploadDir    string `json:"upload_dir"`
	DefaultTitle string `json:"default_title"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config, err := loadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	http.HandleFunc("/", indexHandler)
	log.Printf("Starting server on :8080 with upload directory: %s", config.UploadDir)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
