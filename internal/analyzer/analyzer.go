package analyzer

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

type FileNotFoundError struct{ Path string }

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.Path)
}

type ParseError struct{ Path string }

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing: %s", e.Path)
}

func AnalyzeLog(id, path, typ string) (Result, error) {
	return AnalyzeLogWithRandom(id, path, typ, true)
}

func AnalyzeLogWithRandom(id, path, typ string, enableRandom bool) (Result, error) {
	if _, err := os.Stat(path); err != nil {
		return Result{
			LogID:        id,
			FilePath:     path,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: err.Error(),
		}, &FileNotFoundError{Path: path}
	}

	if enableRandom {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(r.Intn(150)+50) * time.Millisecond)
		if r.Intn(10) == 0 { // 10% d'erreur de parsing
			return Result{
				LogID:        id,
				FilePath:     path,
				Status:       "FAILED",
				Message:      "Erreur de parsing.",
				ErrorDetails: "erreur de parsing simulée",
			}, &ParseError{Path: path}
		}
	}

	return Result{
		LogID:        id,
		FilePath:     path,
		Status:       "OK",
		Message:      "Analyse terminée avec succès.",
		ErrorDetails: "",
	}, nil
}
