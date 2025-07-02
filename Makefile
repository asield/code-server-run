# Makefile para el proyecto dev-env

BINARY_NAME=dev-env
VERSION=1.0.0
MAIN_PACKAGE=.
BIN_DIR=bin
DIST_DIR=dist

.DEFAULT_GOAL := help

build:
	@echo "🔨 Construyendo el binario nativo en la carpeta '$(BIN_DIR)'..."
	@mkdir -p $(BIN_DIR)
	@go build -buildvcs=false -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "✅ Binario '$(BIN_DIR)/$(BINARY_NAME)' creado."

run:
	@echo "🚀 Ejecutando la aplicación..."
	@go run $(MAIN_PACKAGE)

install: build
	@echo "📦 Instalando '$(BINARY_NAME)' desde '$(BIN_DIR)'..."
	@sudo mv $(BIN_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "✅ ¡Instalación completada!"

clean:
	@echo "🧹 Limpiando las carpetas '$(BIN_DIR)' y '$(DIST_DIR)'..."
	@rm -rf $(BIN_DIR) $(DIST_DIR)


PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

cross-compile: clean
	@echo "🌍 Realizando compilación cruzada en la carpeta '$(DIST_DIR)'..."
	@mkdir -p $(DIST_DIR)
	@for platform in $(PLATFORMS); do \
		echo "   -> Construyendo para $$platform..."; \
		GOOS_ARCH=$$(echo $$platform | tr '/' ' '); \
		GOOS=$$(echo $$GOOS_ARCH | awk '{print $$1}'); \
		GOARCH=$$(echo $$GOOS_ARCH | awk '{print $$2}'); \
		BINARY_SUFFIX=""; \
		if [ "$$GOOS" = "windows" ]; then \
			BINARY_SUFFIX=".exe"; \
		fi; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -buildvcs=false -ldflags "-s -w" -o $(DIST_DIR)/$(BINARY_NAME)-$$GOOS-$$GOARCH$$BINARY_SUFFIX $(MAIN_PACKAGE); \
	done
	@echo "✅ Compilación cruzada finalizada."

help:
	@echo "--------------------------------------------------"
	@echo " Ayuda para el Makefile del proyecto dev-env"
	@echo "--------------------------------------------------"
	@echo " Comandos disponibles:"
	@echo "   make build           - Compila el proyecto en la carpeta '$(BIN_DIR)'."
	@echo "   make run             - Compila y ejecuta la aplicación."
	@echo "   make install         - Instala el binario desde '$(BIN_DIR)'."
	@echo "   make cross-compile   - Compila para todas las plataformas en la carpeta '$(DIST_DIR)'."
	@echo "   make clean           - Elimina las carpetas de compilados."
	@echo "   make help            - Muestra este mensaje de ayuda."
	@echo "--------------------------------------------------"

.PHONY: build run install clean cross-compile help