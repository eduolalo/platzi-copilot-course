#!/bin/bash

# Nombre por defecto del proyecto
PROJECT_NAME="no_name_bun_project"

# Procesar argumentos de línea de comandos
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -n|--name) PROJECT_NAME="$2"; shift ;; # Captura el valor después de -n o --name
        *) echo "Opción desconocida: $1"; exit 1 ;;
    esac
    shift # Mueve al siguiente argumento
done

# Verifica si se proporcionó un nombre de proyecto
if [ -z "$PROJECT_NAME" ]; then
  echo "Error: Debes proporcionar un nombre para el proyecto usando -n <nombre> o --name <nombre>"
  exit 1
fi

# salir al directorio padre
cd ..
rm -rf "$PROJECT_NAME" # Elimina el directorio del proyecto si existe

# Crea un proyecto de Bun y Elysia con el nombre proporcionado
echo "Creando proyecto Bun y Elysia llamado '$PROJECT_NAME'..."
bun create elysia ./"$PROJECT_NAME"

# Agrega un Dockerfile al proyecto
echo "Agregando Dockerfile al proyecto '$PROJECT_NAME'..."
cat <<EOL > ./"$PROJECT_NAME"/Dockerfile
# Usa la imagen oficial de Bun
FROM oven/bun:latest as base

# Establece el directorio de trabajo
WORKDIR /usr/src/app

# Copia los archivos de manifiesto de dependencias
COPY package.json bun.lockb* ./

# Instala las dependencias de producción
RUN bun install --frozen-lockfile --production

# Copia el resto del código fuente de la aplicación
COPY . .

# Expone el puerto 3000
EXPOSE 3000

# Comando para iniciar la aplicación (ajustado para usar el entrypoint directamente)
CMD ["bun", "run", "src/index.ts"] # Asegúrate que este sea el comando correcto para tu plantilla
EOL

# Crear archivo .dockerignore
echo "Creando archivo .dockerignore en '$PROJECT_NAME'..."
cat <<EOL > ./"$PROJECT_NAME"/.dockerignore
Dockerfile
.dockerignore
node_modules
.git
.gitignore
README.md
EOL


# Hacer el build de la imagen usando Podman
echo "Construyendo la imagen de Docker con el tag '$PROJECT_NAME'..."
# Usa --tag en lugar de -t para mayor claridad
podman build --tag "$PROJECT_NAME":latest -f ./"$PROJECT_NAME"/Dockerfile ./"$PROJECT_NAME"

# Ejecuta el contenedor
echo "Ejecutando el contenedor con nombre '$PROJECT_NAME'..."
# Agrega --name para darle un nombre al contenedor y -d para modo detached
podman run -d --name "$PROJECT_NAME" -p 3000:3000 "$PROJECT_NAME":latest --build

echo "Script completado. Proyecto '$PROJECT_NAME' creado y contenedor ejecutándose."