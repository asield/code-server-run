// cmd/start.go
package cmd

import (
	"code-server-run/config"
	"code-server-run/runner"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia un entorno de desarrollo detenido en el directorio actual.",
	Long:  `Este comando verifica la configuración y ejecuta 'docker compose up -d' para reanudar el entorno.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := filepath.Join(config.ConfigDir, "config.json")
		configFile, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("❌ No se pudo encontrar el archivo de configuración. ¿Creaste el entorno con 'dev-env create' en este directorio? Error: %v", err)
		}
		defer configFile.Close()

		var envConfig config.TemplateData
		if err := json.NewDecoder(configFile).Decode(&envConfig); err != nil {
			log.Fatalf("❌ Error al leer el archivo de configuración: %v", err)
		}

		log.Printf("▶️  Iniciando el entorno para el proyecto '%s'...", envConfig.ProjectName)

		if err := runner.StartExisting(); err != nil {
			log.Fatalf("❌ Error fatal durante la ejecución: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}