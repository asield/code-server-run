// cmd/create.go
package cmd

import (
	"code-server-run/config"
	"code-server-run/runner"
	//"code-server-run/utils"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var language string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Crea y levanta el entorno de desarrollo en el directorio actual.",
	Long:  `Este comando genera los archivos de configuración y luego ejecuta 'docker compose up' para iniciar el entorno en el puerto 3000.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("🚀 Iniciando la creación del entorno de desarrollo...")

		if _, err := os.Stat(config.ConfigDir); !os.IsNotExist(err) {
			log.Fatalf("❌ Ya existe un directorio '%s' aquí. Si quieres recrear el entorno, primero usa 'dev-env destroy --cleanup'.", config.ConfigDir)
		}

		if !config.IsLanguageSupported(language) {
			log.Fatalf("❌ Lenguaje '%s' no soportado. Los lenguajes disponibles son: go, python, node, rust, java.", language)
		}

		currentUser, err := user.Current()
		if err != nil {
			log.Fatalf("❌ No se pudo obtener el usuario actual: %v", err)
		}

		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("❌ No se pudo obtener el directorio actual: %v", err)
		}

		data := config.TemplateData{
			Username:    currentUser.Username,
			ProjectName: filepath.Base(wd),
			Language:    language,
		}

		fmt.Print("🔑 Por favor, introduce la contraseña para code-server: ")
		passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("❌ No se pudo leer la contraseña: %v", err)
		}
		fmt.Println()

		if err := config.GenerateFiles(data, string(passwordBytes)); err != nil {
			log.Fatalf("❌ Error al generar los archivos de configuración: %v", err)
		}

		log.Println("▶️  Archivos listos. Iniciando el entorno de Docker...")
		if err := runner.Start(currentUser.Uid, currentUser.Gid); err != nil {
			log.Fatalf("❌ Error fatal durante la ejecución: %v", err)
		}

		log.Printf("🎉 ¡Entorno de %s listo y funcionando en http://localhost:3000!\n", language)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&language, "lang", "l", "go", "El lenguaje de programación para el entorno (go, python, node, rust, java)")
}