// runner/runner.go
package runner

import (
	"bytes"
	"fmt"
	"log"
	"code-server-run/utils"
	"os"
	"os/exec"
	"time"
)

const serverURL = "http://localhost:8080"

func Start(uid, gid string) error {
	log.Println("   - Obteniendo GID de Docker...")
	dockerGid, err := utils.GetDockerGID()
	if err != nil {
		log.Printf("⚠️  Advertencia: No se pudo obtener el GID de Docker. Usando valor por defecto. Error: %v", err)
		dockerGid = "999"
	}
	log.Printf("   - Usando UID=%s, GID=%s, DOCKER_GID=%s", uid, gid, dockerGid)

	log.Println("🐳 Ejecutando docker-compose...")

	cmd := exec.Command("docker-compose", "-f", ".dev-env/compose.yml", "up", "--build", "-d")
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("UID=%s", uid),
		fmt.Sprintf("GID=%s", gid),
		fmt.Sprintf("DOCKER_GID=%s", dockerGid),
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al ejecutar docker-compose: %v\n%s", err, stderr.String())
	}

	if err := utils.WaitForServer(serverURL, 60*time.Second); err != nil {
		return err
	}

	utils.OpenBrowser(serverURL)

	return nil
}