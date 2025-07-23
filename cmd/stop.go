// cmd/stop.go
package cmd

import (
	"code-server-run/runner"
	"log"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Detiene (pausa) el entorno de desarrollo en el directorio actual.",
	Long:  `Este comando ejecuta 'docker compose stop' para detener los contenedores sin eliminarlos.`,
	// La corrección está aquí: cobra.Command en lugar de cobra.command
	Run: func(cmd *cobra.Command, args []string) {
		// Llamada única al runner
		if err := runner.Stop(); err != nil {
			log.Fatalf("❌ %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
