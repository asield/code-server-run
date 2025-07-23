// cmd/destroy.go
package cmd

import (
	"code-server-run/runner"
	"log"

	"github.com/spf13/cobra"
)

var cleanup bool

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Detiene y elimina el entorno de desarrollo y sus archivos.",
	Long:  `Ejecuta 'docker compose down' para detener y eliminar los contenedores, redes y volúmenes. Opcionalmente, puede limpiar los archivos de configuración.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Llamada al runner, pasando el valor del flag
		if err := runner.Destroy(cleanup); err != nil {
			log.Fatalf("❌ Error durante la destrucción: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
	destroyCmd.Flags().BoolVarP(&cleanup, "cleanup", "c", false, "Elimina el directorio de configuración .dev-env sin preguntar.")
}