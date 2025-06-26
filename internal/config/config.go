package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func ReadConfig(path string) ([]LogConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier de config : %w", err)
	}
	defer file.Close()
	var logs []LogConfig
	dec := json.NewDecoder(file)
	if err := dec.Decode(&logs); err != nil {
		return nil, fmt.Errorf("erreur de décodage JSON : %w", err)
	}
	return logs, nil
}
