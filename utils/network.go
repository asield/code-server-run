// utils/network.go
package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer espera a que un servidor HTTP esté disponible en una URL específica.
func WaitForServer(url string, timeout time.Duration) error {
	fmt.Print("⏳ Esperando que el servidor esté listo... ")

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		resp, err := client.Get(url)
		// Aceptamos cualquier código de estado 2xx o 3xx como una señal de que el servidor está activo.
		if err == nil && resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
			if resp.Body != nil {
				resp.Body.Close()
			}
			fmt.Println("¡OK!")
			log.Println("✅ ¡Servidor listo!")
			return nil
		}
		if resp != nil {
			resp.Body.Close()
		}

		fmt.Print(".")
		time.Sleep(2 * time.Second)
	}

	fmt.Println()
	return fmt.Errorf("el servidor en %s no respondió en %v", url, timeout)
}