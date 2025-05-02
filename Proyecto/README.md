# Generador automático de proyectos Bun + Elysia + Docker

Este proyecto contiene un script llamado `generador.sh` que automatiza la creación de un proyecto basado en Bun y Elysia, además de preparar los archivos necesarios para contenerizar la aplicación usando Docker (o Podman).

## ¿Qué hace este proyecto?

- Crea un nuevo proyecto Bun con Elysia usando el comando `bun create elysia`.
- Elimina cualquier carpeta previa con el mismo nombre para evitar conflictos.
- Genera automáticamente un `Dockerfile` y un archivo `.dockerignore` en la carpeta del nuevo proyecto.
- Construye la imagen del contenedor usando Podman.
- Ejecuta el contenedor exponiendo el puerto 3000.

## Uso

Desde la carpeta `Proyecto`, ejecuta el script con el nombre del proyecto que desees crear:

```sh
bash generador.sh -n <nombre_del_proyecto>
```

o

```sh
bash generador.sh --name <nombre_del_proyecto>
```

Ejemplo:

```sh
bash generador.sh -n mi-api-bun
```

Esto creará una carpeta `mi-api-bun` en el directorio superior, con la estructura de un proyecto Bun + Elysia y los archivos necesarios para Docker.

## Estructura generada

```
<nombre_del_proyecto>/
├── Dockerfile
├── .dockerignore
├── package.json
├── bun.lockb
├── src/
│   └── index.ts
└── ...
```

## Requisitos

- [Bun](https://bun.sh/)
- [Podman](https://podman.io/) (puedes adaptar los comandos a Docker si lo prefieres)

## Notas

- El comando de inicio en el Dockerfile asume que el entrypoint es `src/index.ts`. Modifícalo si tu plantilla cambia.
- El contenedor se ejecuta en modo detached y expone el puerto 3000.

---

Generado automáticamente por [`generador.sh`](Proyecto/generador.sh)
