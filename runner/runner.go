// runner/runner.go
package runner

import (
	"bufio"
	"bytes"
	"code-server-run/config"
	"code-server-run/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func newComposeCmd(args ...string) *exec.Cmd {
	composeFile := filepath.Join(config.ConfigDir, "compose.yml")
	baseArgs := []string{"compose", "-f", composeFile}
	fullArgs := append(baseArgs, args...)
	return exec.Command("docker", fullArgs...)
}

// Start ya no necesita el puerto.
func Start(uid, gid string) error {
	log.Println("   - Obteniendo GID de Docker...")
	dockerGid, err := utils.GetDockerGID()
	if err != nil {
		log.Printf("⚠️  Advertencia: No se pudo obtener el GID de Docker. Usando valor por defecto. Error: %v", err)
		dockerGid = "999"
	}
	log.Printf("   - Usando UID=%s, GID=%s, DOCKER_GID=%s", uid, gid, dockerGid)

	log.Println("🐳 Ejecutando docker compose...")

	cmd := newComposeCmd("up", "--build", "-d")
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("UID=%s", uid),
		fmt.Sprintf("GID=%s", gid),
		fmt.Sprintf("DOCKER_GID=%s", dockerGid),
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al ejecutar docker compose: %v\n%s", err, stderr.String())
	}

	serverURL := "http://localhost:3000"
	if err := utils.WaitForServer(serverURL, 90*time.Second); err != nil {
		return err
	}

	utils.OpenBrowser(serverURL)
	return nil
}

// StartExisting ya no necesita el puerto.
func StartExisting() error {
	log.Println("▶️  Iniciando el entorno de desarrollo...")
	cmd := newComposeCmd("up", "-d")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al iniciar el entorno: %w", err)
	}

	log.Println("✅ ¡Entorno iniciado correctamente!")
	serverURL := "http://localhost:3000"
	utils.OpenBrowser(serverURL)
	return nil
}

func Stop() error {
	log.Println("⏸️  Deteniendo el entorno de desarrollo...")
	cmd := newComposeCmd("stop")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al detener el entorno. ¿Estás seguro de que hay uno activo? Error: %w", err)
	}
	log.Println("✅ ¡Entorno detenido! Puedes reanudarlo con 'dev-env start'.")
	return nil
}

func Destroy(cleanup bool) error {
	log.Println("🔥 Destruyendo el entorno de desarrollo...")

	cmd := newComposeCmd("down", "--volumes")
	if err := cmd.Run(); err != nil {
		log.Printf("⚠️  Advertencia: 'docker compose down' falló. Puede que el entorno no existiera. Error: %v", err)
	} else {
		log.Println("✅ Contenedores, redes y volúmenes con nombre eliminados.")
	}

	if _, err := os.Stat(config.ConfigDir); !os.IsNotExist(err) {
		shouldDelete := cleanup
		if !shouldDelete {
			fmt.Printf("🧹 ¿Deseas eliminar el directorio de configuración '%s' y todo su contenido? [y/N]: ", config.ConfigDir)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(input)) == "y" {
				shouldDelete = true
			}
		}

		if shouldDelete {
			log.Printf("🧹 Limpiando el directorio de configuración '%s'...", config.ConfigDir)
			if err := os.RemoveAll(config.ConfigDir); err != nil {
				return fmt.Errorf("no se pudo eliminar el directorio %s: %w", config.ConfigDir, err)
			}
		}
	}

	log.Println("👍 Proceso de destrucción finalizado.")
	return nil
}