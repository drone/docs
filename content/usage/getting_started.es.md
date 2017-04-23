+++
date = "2017-04-15T14:39:04+02:00"
title = "Empezando"
url = "es/getting-started"

[menu.usage]
  weight = 1
  identifier = "getting-started-es"
  parent = "usage_overview"
+++

Bienvenido a la comunidad de Drone. este documento explica brevenemte el proceso para activar y configurar un flujo de integración continua.

# Activación

Para activar nuestro proyecto, navega a la configuración de tu cuenta. Verás  una lista de repositorios que pueden ser activados con un simple botón. Cuando activas un repositorio, Drone automáticamente agrega los Webhooks a tu sistema de control de veersiones (Ej: GitHub).

Los Webhooks son usados para ejecutar el flujo de ejecuciones. Cuando realizas un push a tu repositorio, abres un pull request, o creas una etiqueta, tu sistema de control de versiones automáticamente enviará un Webhook a Drone el cual ejecutará el flujo de ejecución.

![Lista de repositorios](/images/drone_repo_list.png)

# Configuración

Para configurar tu flujo de ejecución debes crear un archivo `.drone.yml` en la raiz de tu repositorio. El archivo `drone.yml` es usado para definir los pasos del flujo de ejecución. Es un super conjunto del muy conocido archivo `docker-compose`.

Ejemplo de una configuración de flujo de ejecución.

```
pipeline:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test

services:
  postgres:
    image: postgres:9.4.5
    environment:
      - POSTGRES_USER=myapp
```

Ejemplo de la configuración de un flujo de ejecución con múltiples pasos sucesivos:

```
pipeline:
  backend:
    image: golang
    commands:
      - go get
      - go build
      - go test

  frontend:
    image: node:6
    commands:
      - npm install
      - npm test

  notify:
    image: plugins/slack
    channel: developers
    username: drone
```

# Ejecución

Para ejecutar tu primer flujo de ejecución puedes ejecutar un push a tu repositorio, abrir un pull request, o crear una etiqueta. Cualquiera de estos eventos lanza un Webhook desde tu sistema de control de versiones y ejecuta el flujo de ejecución.

Puedes revisar el flujo de ejecución en tiempo real usando la interfaz de usuario.

![running build](/images/drone_build_running.png)
