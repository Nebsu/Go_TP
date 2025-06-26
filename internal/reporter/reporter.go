package reporter

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Nebsu/Go_TP/internal/analyzer"
)

func ExportReport(path string, results []analyzer.Result) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(results)
}
