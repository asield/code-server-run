// cmd/destroy.go
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var cleanup bool

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Detiene y elimina el entorno de desarrollo y sus archivos.",
	Long:  `Ejecuta 'docker-compose down' para detener y eliminar los contenedores, redes y volúmenes. Opcionalmente, puede limpiar los archivos de configuración.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("🔥 Destruyendo el entorno de desarrollo...")

		composeCmd := exec.Command("docker-compose", "-f", ".dev-env/compose.yml", "down", "--volumes")
		if err := composeCmd.Run(); err != nil {
			log.Printf("⚠️  Advertencia: 'docker-compose down' falló. Puede que el entorno no existiera. Error: %v", err)
		} else {
			log.Println("✅ Contenedores, redes y volúmenes con nombre eliminados.")
		}

		configDir := ".dev-env"
		if _, err := os.Stat(configDir); !os.IsNotExist(err) {
			if cleanup {
				log.Printf("🧹 Limpiando el directorio de configuración '%s'...", configDir)
				os.RemoveAll(configDir)
			} else {
				fmt.Printf("🧹 ¿Deseas eliminar el directorio de configuración '%s' y todo su contenido? [y/N]: ", configDir)
				reader := bufio.NewReader(os.Stdin)
				input, _ := reader.ReadString('\n')
				if strings.ToLower(strings.TrimSpace(input)) == "y" {
					log.Printf("🧹 Limpiando el directorio de configuración '%s'...", configDir)
					os.RemoveAll(configDir)
				}
			}
		}

		log.Println("👍 Proceso de destrucción finalizado.")
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
	destroyCmd.Flags().BoolVarP(&cleanup, "cleanup", "c", false, "Elimina el directorio de configuración .dev-env sin preguntar.")
}