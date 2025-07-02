// config/generator.go
package config

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type TemplateData struct {
	Username    string
	ProjectName string
}

func GenerateFiles(data TemplateData, password string) error {
	tmpl, err := template.ParseFS(Templates, "*.tmpl")
	if err != nil {
		return fmt.Errorf("error al parsear plantillas desde FS: %w", err)
	}

	configDir := ".dev-env"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("no se pudo crear el directorio de configuración: %w", err)
	}

	if err := renderTemplate(filepath.Join(configDir, "Dockerfile"), "Dockerfile.tmpl", tmpl, data); err != nil {
		return err
	}


	if err := renderTemplate(filepath.Join(configDir, "compose.yml"), "compose.yml.tmpl", tmpl, data); err != nil {
		return err
	}

	envContent := fmt.Sprintf("DEV_PASSWORD=%s\n", password)
	return os.WriteFile(filepath.Join(configDir, ".env"), []byte(envContent), 0644)
}

func renderTemplate(path, templateName string, tmpl *template.Template, data interface{}) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("error al crear archivo %s: %w", path, err)
		}
		defer file.Close()

		if err := tmpl.ExecuteTemplate(file, templateName, data); err != nil {
			return fmt.Errorf("error al ejecutar plantilla %s: %w", path, err)
		}
	}
	return nil
}