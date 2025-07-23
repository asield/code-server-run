// config/generator.go
package config

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

const ConfigDir = ".dev-env"

// TemplateData ya no necesita el puerto.
type TemplateData struct {
	Username    string `json:"username"`
	ProjectName string `json:"projectName"`
	Language    string `json:"language"`
}

var supportedLanguages = map[string]bool{
	"go":     true,
	"python": true,
	"node":   true,
	"rust":   true,
	"java":   true,
}

func IsLanguageSupported(lang string) bool {
	_, ok := supportedLanguages[lang]
	return ok
}

func GenerateFiles(data TemplateData, password string) error {
	tmpl, err := template.ParseFS(Templates, "*.tmpl")
	if err != nil {
		return fmt.Errorf("error al parsear plantillas desde FS: %w", err)
	}

	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return fmt.Errorf("no se pudo crear el directorio de configuración: %w", err)
	}

	dockerfilePath := filepath.Join(ConfigDir, "Dockerfile")
	composePath := filepath.Join(ConfigDir, "compose.yml")
	configPath := filepath.Join(ConfigDir, "config.json")
	envPath := filepath.Join(ConfigDir, ".env")

	if err := renderTemplate(dockerfilePath, "Dockerfile.tmpl", tmpl, data); err != nil {
		return err
	}

	if err := renderTemplate(composePath, "compose.yml.tmpl", tmpl, data); err != nil {
		return err
	}

	if err := saveConfig(configPath, data); err != nil {
		return err
	}

	envContent := fmt.Sprintf("DEV_PASSWORD=%s\n", password)
	return os.WriteFile(envPath, []byte(envContent), 0644)
}

func renderTemplate(path, templateName string, tmpl *template.Template, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error al crear archivo %s: %w", path, err)
	}
	defer file.Close()

	if err := tmpl.ExecuteTemplate(file, templateName, data); err != nil {
		return fmt.Errorf("error al ejecutar plantilla %s: %w", path, err)
	}
	return nil
}

func saveConfig(path string, data TemplateData) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error al crear archivo de configuración %s: %w", path, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("error al codificar datos a JSON: %w", err)
	}
	return nil
}