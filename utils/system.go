// utils/system.go
package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func GetDockerGID() (string, error) {
	cmd := exec.Command("sh", "-c", "getent group docker | cut -d: -f3")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("el comando para obtener GID falló: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func OpenBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		if exec.Command("sh", "-c", "grep -qE '(Microsoft|WSL)' /proc/version").Run() == nil {
			cmd = "explorer.exe"
		} else {
			cmd = "xdg-open"
		}
		args = []string{url}
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		log.Printf("Sistema operativo no soportado para abrir el navegador automáticamente.")
		log.Printf("Por favor, abre esta URL manualmente: %s", url)
		return
	}

	log.Printf("🌐 Abriendo %s en tu navegador...", url)
	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Printf("⚠️  No se pudo abrir el navegador: %v. Por favor, abre la URL manualmente.", err)
	}
}