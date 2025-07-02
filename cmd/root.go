// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-env",
	Short: "Herramienta CLI para gestionar entornos de desarrollo portables con Docker.",
	Long: `dev-env es una utilidad de línea de comandos que te permite crear
y destruir entornos de desarrollo completos basados en Go, Docker y code-server
directamente en cualquier directorio de proyectos.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
