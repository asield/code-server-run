# dev-env: Portable Go Development Environment

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://golang.org) [![Build with Makefile](https://img.shields.io/badge/Build-Makefile-brightgreen)](Makefile) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**`dev-env`** is a powerful command-line tool that creates and manages portable, containerized development environments for Go. It leverages Docker, Docker Compose, and `code-server` (VS Code in the browser) to provide a consistent, high-performance workspace on any machine.

This tool automatically generates all necessary configuration files (`Dockerfile`, `compose.yml`, `.env`) tailored to your system and project, allowing you to spin up a complete development environment with a single command.

---

[**🇬🇧 English**](#-english) | [**🇪🇸 Español**](#-español)

---

<a name="english"></a>
## 🇬🇧 English

### ✨ Key Features

* **Zero Configuration:** Generates all required files on the fly based on your current user and directory.
* **VS Code in the Browser:** Provides a full-featured `code-server` instance, accessible from any web browser.
* **Seamless File Permissions:** Automatically syncs host user (UID) and group (GID) IDs with the container to prevent permission conflicts.
* **Docker-in-Docker Ready:** Includes the Docker CLI and mounts the host's Docker socket, allowing you to run Docker commands from within the dev environment.
* **Fully Contained:** The Go binary includes all templates, making it a single, portable executable.
* **Professional CLI:** Built with Cobra, providing intuitive `create` and `destroy` commands.
* **Pre-configured:** Comes with the official Go extension and Dracula theme ready to use.

---

### 📋 Prerequisites

* **Go** (version 1.16+ to build the tool).
* **Docker Engine** & **Docker Compose**.
* **Make** build-automation tool.
* A **Linux-based** OS or **WSL2** on Windows is recommended for full feature compatibility.

---

### 📦 Installation

Clone this repository and use the included `Makefile` to compile and install the binary system-wide.

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

Using `dev-env` is designed to be simple and intuitive.

1. **Create a new project directory and navigate into it:**

    ```bash
    mkdir ~/my-go-api
    cd ~/my-go-api
    ```

2. **Create the development environment:**

    ```bash
    dev-env create
    ```

    The tool will prompt you for a password for `code-server` and then automatically generate the config files, build the Docker image, and start the containers. Finally, it will open the environment in your default web browser at `http://localhost:8080`.

3. **Destroy the environment:**  
   When you are finished working, run the `destroy` command from the same project directory.

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

This project uses a `Makefile` to automate common tasks.

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

* **Cero Configuración:** Genera todos los archivos necesarios al vuelo, basándose en tu usuario y directorio actual.
* **VS Code en el Navegador:** Provee una instancia completa de `code-server`, accesible desde cualquier navegador web.
* **Permisos de Archivo Perfectos:** Sincroniza automáticamente los IDs de tu usuario (UID) y grupo (GID) del anfitrión con el contenedor para prevenir conflictos de permisos.
* **Listo para Docker-in-Docker:** Incluye el CLI de Docker y monta el socket de Docker del anfitrión, permitiéndote ejecutar comandos de Docker desde dentro del entorno.
* **Autocontenido:** El binario de Go incluye todas las plantillas, convirtiéndolo en un único ejecutable portable.
* **CLI Profesional:** Construido con Cobra, proveyendo comandos intuitivos como `create` y `destroy`.
* **Pre-configurado:** Viene con la extensión oficial de Go y el tema Dracula listos para usar.

---

### 📋 Prerrequisitos

* **Go** (versión 1.16+ para compilar la herramienta).
* **Motor de Docker** y **Docker Compose**.
* Herramienta de automatización **Make**.
* Un sistema operativo basado en **Linux** (o **WSL2** en Windows) es recomendado para la compatibilidad total de las características.

---

### 📦 Instalación

Clona este repositorio y usa el `Makefile` incluido para compilar e instalar el binario en todo el sistema.

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
    mkdir ~/mi-api-en-go
    cd ~/mi-api-en-go
    ```

2. **Crea el entorno de desarrollo:**

    ```bash
    dev-env create
    ```

    La herramienta te pedirá una contraseña para `code-server` y luego generará automáticamente los archivos de configuración, construirá la imagen de Docker e iniciará los contenedores. Finalmente, abrirá el entorno en tu navegador por defecto en `http://localhost:8080`.

3. **Destruye el entorno:**  
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

Este proyecto utiliza un `Makefile` para automatizar tareas comunes.

* `make build`: Compila el binario para tu sistema actual en la carpeta `bin/`.
* `make cross-compile`: Compila binarios para Linux, macOS y Windows en la carpeta `dist/`.
* `make clean`: Elimina todos los artefactos de compilación de las carpetas `bin/` y `dist/`.
* `make help`: Muestra una lista de todos los comandos disponibles.

---

### ⚠️ Aviso de Seguridad

Esta herramienta monta el socket de Docker del anfitrión (`/var/run/docker.sock`) en el contenedor. Esto otorga al contenedor privilegios de acceso equivalentes a `root` en tu máquina anfitriona. **Usa esta herramienta únicamente en entornos de desarrollo locales y de confianza.**
