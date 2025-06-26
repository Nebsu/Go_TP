package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Nebsu/Go_TP/internal/analyzer"
	"github.com/Nebsu/Go_TP/internal/config"
	"github.com/Nebsu/Go_TP/internal/reporter"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les logs depuis un fichier de configuration JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, _ := cmd.Flags().GetString("config")
		outputPath, _ := cmd.Flags().GetString("output")
		statusFilter, _ := cmd.Flags().GetString("status")
		logs, err := config.ReadConfig(configPath)
		if err != nil {
			return err
		}

		resultsCh := make(chan analyzer.Result, len(logs))
		done := make(chan struct{})
		var wg sync.WaitGroup

		for _, l := range logs {
			wg.Add(1)
			go func(log config.LogConfig) {
				defer wg.Done()
				res, err := analyzer.AnalyzeLog(log.ID, log.Path, log.Type)
				if err != nil {
					var fnf *analyzer.FileNotFoundError
					var parseErr *analyzer.ParseError
					switch {
					case errors.As(err, &fnf):
						res.Message += " (erreur personnalisée : fichier introuvable)"
					case errors.As(err, &parseErr):
						res.Message += " (erreur personnalisée : parsing)"
					case errors.Is(err, os.ErrNotExist):
						res.Message += " (erreur système : fichier non existant)"
					}
				}
				resultsCh <- res
			}(l)
		}

		go func() {
			wg.Wait()
			close(resultsCh)
			close(done)
		}()

		var results []analyzer.Result
		for res := range resultsCh {
			if statusFilter == "" || res.Status == statusFilter {
				results = append(results, res)
				fmt.Printf("[%s] %s : %s - %s\n", res.LogID, res.FilePath, res.Status, res.Message)
			}
		}
		<-done

		if outputPath != "" {
			importPath := outputPath
			if ext := filepath.Ext(outputPath); ext == ".json" {
				base := outputPath[:len(outputPath)-len(ext)]
				importPath = fmt.Sprintf("%s_%s%s", base, time.Now().Format("060102"), ext)
			}
			if err := reporter.ExportReport(importPath, results); err != nil {
				return err
			}
			fmt.Printf("Rapport exporté : %s\n", importPath)
		}
		return nil
	},
}

func init() {
	analyzeCmd.Flags().StringP("config", "c", "", "Chemin du fichier de configuration JSON")
	analyzeCmd.Flags().StringP("output", "o", "", "Chemin du fichier de rapport JSON à exporter")
	analyzeCmd.Flags().String("status", "", "Filtrer les résultats par statut (OK ou FAILED)")
	analyzeCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(analyzeCmd)
}
