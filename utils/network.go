// utils/network.go
package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string, timeout time.Duration) error {
	fmt.Print("⏳ Esperando que el servidor esté listo... ")

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		resp, err := client.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
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
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	return fmt.Errorf("el servidor no respondió en %v", timeout)
}