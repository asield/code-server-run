# dev-env: Entornos de Desarrollo Portables / Portable Development Environments

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://golang.org) [![Build with Makefile](https://img.shields.io/badge/Build-Makefile-brightgreen)](Makefile) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**dev-env** is a powerful command-line tool that creates and manages portable, containerized development environments for multiple languages. It leverages Docker, Docker Compose, and code-server (VS Code in the browser) to provide a consistent, high-performance workspace on any machine.

Esta herramienta genera automáticamente todos los archivos de configuración necesarios (`Dockerfile`, `compose.yml`, `.env`) adaptados a tu sistema y proyecto, permitiéndote levantar un entorno de desarrollo completo con un solo comando.

---

[**🇬🇧 English**](#-english) | [**🇪🇸 Español**](#-español)

---

<a name="english"></a>

## 🇬🇧 English

### ✨ Key Features

* **Multi-Language Support:** Create development environments for Go, Python, Node.js, Rust, and Java with a single flag.
* **Zero Configuration:** Generates all required files on the fly based on your current user and directory.
* **VS Code in the Browser:** Provides a full-featured code-server instance, accessible from any web browser at [http://localhost:8080](http://localhost:8080).
* **Seamless File Permissions:** Automatically handles user (UID) and group (GID) IDs to prevent permission conflicts inside the container.
* **Docker-in-Docker Ready:** Includes the Docker CLI and mounts the host's Docker socket, allowing you to run Docker commands from within the dev environment.
* **Fully Contained:** The Go binary includes all templates, making it a single, portable executable.
* **Professional CLI:** Built with Cobra, providing intuitive `create`, `destroy`, `start`, and `stop` commands.
* **Pre-configured:** Comes with essential extensions for each language (e.g., golang.go, ms-python.python, rust-analyzer), plus Docker and Gemini Code Assist.

---

### 📋 Prerequisites

* **Go** (version 1.24+ to build the tool).
* **Docker Engine** & **Docker Compose** (V2 plugin syntax is used).
* **Make** build-automation tool.
* A **Linux-based OS** or **WSL2** on Windows is recommended for full feature compatibility.

---

### 📦 Installation

Clone this repository and use the included Makefile to compile and install the binary system-wide.

```bash
# 1. Clone the repository
git clone <YOUR_REPOSITORY_URL>
cd <REPOSITORY_NAME>

# 2. Compile and install the binary into /usr/local/bin
# This will ask for your password as it uses 'sudo'.
make install
```

After this, the `dev-env` command will be available anywhere on your system.

---

### 🚀 Usage

Using `dev-env` is simple and intuitive.

1. **Create a new project directory and navigate into it:**

   ```bash
   mkdir ~/my-project
   cd ~/my-project
   ```

2. **Create the development environment:**
   Use the `--lang` flag to specify the language.

   ```bash
   # Example for a Go environment
   dev-env create --lang go

   # Example for a Python environment
   dev-env create --lang python
   ```

   The tool will prompt you for a password for code-server, build the Docker image, start the container, and open the environment in your default browser at [http://localhost:8080](http://localhost:8080).

3. **Stop and start the environment:**

   ```bash
   dev-env stop      # Stops the running container
   dev-env start     # Resumes the stopped container
   ```

4. **Destroy the environment:**
   When you finish working, run the `destroy` command from the same project directory.

   ```bash
   dev-env destroy
   ```

   This will stop and remove all containers, networks, and named volumes. It will also ask if you want to clean up the generated configuration files.

   To force cleanup without being prompted:

   ```bash
   dev-env destroy --cleanup
   ```

---

### 🛠️ Building & Development

This project uses a Makefile to automate common tasks.

* `make build`: Compiles the binary for your current system into the `bin/` directory.
* `make cross-compile`: Compiles binaries for Linux, macOS, and Windows into the `dist/` directory.
* `make clean`: Removes all build artifacts from the `bin/` and `dist/` directories.
* `make help`: Displays a list of all available commands.

---

### ⚠️ Security Notice

This tool mounts the host's Docker socket (`/var/run/docker.sock`) into the container. This grants the container privileged access equivalent to `root` on your host machine. **Use this tool only in trusted local development environments.**

---

<a name="español"></a>

## 🇪🇸 Español

### ✨ Características Principales

* **Soporte Multi-lenguaje:** Crea entornos de desarrollo para Go, Python, Node.js, Rust y Java con un solo flag.
* **Cero Configuración:** Genera todos los archivos necesarios al vuelo, basándose en tu usuario y directorio actual.
* **VS Code en el Navegador:** Provee una instancia completa de code-server, accesible desde cualquier navegador web en [http://localhost:8080](http://localhost:8080).
* **Permisos de Archivo Perfectos:** Gestiona automáticamente los IDs de tu usuario (UID) y grupo (GID) para prevenir conflictos de permisos dentro del contenedor.
* **Listo para Docker-in-Docker:** Incluye el CLI de Docker y monta el socket de Docker del anfitrión, permitiéndote ejecutar comandos de Docker desde dentro del entorno.
* **Autocontenido:** El binario de Go incluye todas las plantillas, convirtiéndolo en un único ejecutable portable.
* **CLI Profesional:** Construido con Cobra, proveyendo comandos intuitivos como `create`, `destroy`, `start` y `stop`.
* **Pre-configurado:** Viene con extensiones esenciales para cada lenguaje (ej. golang.go, ms-python.python, rust-analyzer), además de Docker y Gemini Code Assist.

---

### 📋 Prerrequisitos

* **Go** (versión 1.24+ para compilar la herramienta).
* **Motor de Docker** y **Docker Compose** (se usa la sintaxis de plugin V2).
* Herramienta de automatización **Make**.
* Un sistema operativo basado en **Linux** (o **WSL2** en Windows) es recomendado para la compatibilidad total de las características.

---

### 📦 Instalación

Clona este repositorio y usa el Makefile incluido para compilar e instalar el binario en todo el sistema.

```bash
# 1. Clona el repositorio
git clone <URL-DE-TU-REPOSITORIO>
cd <NOMBRE-DEL-REPOSITORIO>

# 2. Compila e instala el binario en /usr/local/bin
# Te pedirá tu contraseña ya que usa 'sudo'.
make install
```

Después de esto, el comando `dev-env` estará disponible desde cualquier lugar en tu sistema.

---

### 🚀 Uso

Usar `dev-env` está diseñado para ser simple e intuitivo.

1. **Crea una nueva carpeta de proyecto y entra en ella:**

   ```bash
   mkdir ~/mi-proyecto
   cd ~/mi-proyecto
   ```

2. **Crea el entorno de desarrollo:**
   Usa el flag `--lang` para especificar el lenguaje.

   ```bash
   # Ejemplo para un entorno de Go
   dev-env create --lang go

   # Ejemplo para un entorno de Python
   dev-env create --lang python
   ```

   La herramienta te pedirá una contraseña para code-server, construirá la imagen de Docker, iniciará el contenedor y abrirá el entorno en tu navegador por defecto en [http://localhost:8080](http://localhost:8080).

3. **Detén e inicia el entorno:**

   ```bash
   dev-env stop      # Detiene el contenedor en ejecución
   dev-env start     # Reanuda el contenedor detenido
   ```

4. **Destruye el entorno:**
   Cuando termines de trabajar, ejecuta el comando `destroy` desde la misma carpeta del proyecto.

   ```bash
   dev-env destroy
   ```

   Esto detendrá y eliminará todos los contenedores, redes y volúmenes nombrados. También te preguntará si deseas limpiar los archivos de configuración generados.

   Para forzar la limpieza sin que te pregunte:

   ```bash
   dev-env destroy --cleanup
   ```

---

### 🛠️ Compilación y Desarrollo

Este proyecto utiliza un Makefile para automatizar tareas comunes.

* `make build`: Compila el binario para tu sistema actual en la carpeta `bin/`.
* `make cross-compile`: Compila binarios para Linux, macOS y Windows en la carpeta `dist/`.
* `make clean`: Elimina todos los artefactos de compilación de las carpetas `bin/` y `dist/`.
* `make help`: Muestra una lista de todos los comandos disponibles.

---

### ⚠️ Aviso de Seguridad

Esta herramienta monta el socket de Docker del anfitrión (`/var/run/docker.sock`) en el contenedor. Esto otorga al contenedor privilegios de acceso equivalentes a `root` en tu máquina anfitriona. **Usa esta herramienta únicamente en entornos de desarrollo locales y de confianza.**
