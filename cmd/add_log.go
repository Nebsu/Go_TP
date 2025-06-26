package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Nebsu/Go_TP/internal/config"
	"github.com/spf13/cobra"
)

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajoute une configuration de log au fichier config.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetString("id")
		path, _ := cmd.Flags().GetString("path")
		typeLog, _ := cmd.Flags().GetString("type")
		file, _ := cmd.Flags().GetString("file")
		if id == "" || path == "" || typeLog == "" || file == "" {
			return fmt.Errorf("tous les drapeaux --id, --path, --type, --file sont requis")
		}
		logs, _ := config.ReadConfig(file)
		logs = append(logs, config.LogConfig{ID: id, Path: path, Type: typeLog})
		f, err := os.Create(file)
		if err != nil {
			return err
		}
		defer f.Close()
		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		if err := enc.Encode(logs); err != nil {
			return err
		}
		fmt.Printf("Log ajouté à %s\n", file)
		return nil
	},
}

func init() {
	addLogCmd.Flags().String("id", "", "Identifiant du log")
	addLogCmd.Flags().String("path", "", "Chemin du fichier de log")
	addLogCmd.Flags().String("type", "", "Type de log")
	addLogCmd.Flags().String("file", "", "Chemin du fichier config.json")
	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
	addLogCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(addLogCmd)
}
