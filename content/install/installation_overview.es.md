+++
date = "2017-04-15T14:39:04+02:00"
title = "Instalación"
url = "es/installation"

[menu.install]
  weight = 1
  identifier = "installation-es"
  parent = "install_overview"
+++

Drone es una plataforma de despliegue continuo liviana y poderosa  construida para contenedores. Drone está enpaquetado y distribuido como una imagen de Docker y puede ser descargado desde Dockerhub.

```text
docker pull drone/drone:{{% version %}}
```

# Docker Compose

Esta sección provee instrucciones básicas para la instalación de Drone usando [docker-compose](https://docs.docker.com/compose/). La configuración que mostraremos acontinuación puede ser usada para iniciar un servidor de docker con un único agente.

```yaml
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}

  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

Drone se integra con múltiples proveedores de control de versiones, configurables usando variables de entorno. Este ejemplo muestra la integración básica con Github.

{{% alert %}}
Debes registrar Drone con Github para obtener las llaves de cliente y secreta. El callback de autorización debe coincidir con  `<scheme>://<host>/authorize`
{{% /alert %}}

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_ORGS=dolores,dogpatch
      - DRONE_ADMIN=johnsmith,janedoe
+     - DRONE_GITHUB=true
+     - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
+     - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

Drone monta volúmenes con la máquina anfitrión para almacenar la base de datos sqlite.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
+   volumes:
+     - ./drone:/var/lib/drone/
    restart: always
```

Drone necesita saber su propia dirección. Por esta razón debes proveeder la dirección en el formato `<scheme>://<hostname>`.


```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
+     - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

El agente Drone requiere acceso al demonio Docker en la máquina anfitrión.

```diff
services:
  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on: [ drone-server ]
+   volumes:
+     - /var/run/docker.sock:/var/run/docker.sock
```

Los agentes Drone requieren la dirección del servidor para comunicación agente-a-servidor.

```diff
services:
  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on: [ drone-server ]
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
+     - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

El servidor y los agentes Drone usan una clave secreta para autenticar su comunicación. Esto debería ser una cadena aleatoria de caracteres de tu elección y debería mantenerse en privado.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
+     - DRONE_SECRET=${DRONE_SECRET}
  drone-agent:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_DEBUG: true
+     - DRONE_SECRET=${DRONE_SECRET}
```

El registro a Drone está cerrado por defecto. Este ejemplo habilita el registro abierto para miembros aprovados por las organizaciones de Github.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_OPEN=true
+     - DRONE_ORGS=dolores,dogpatch
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```

Los administradores de Drone pueden ser enumerados en tu configucación.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_ORGS=dolores,dogpatch
+     - DRONE_ADMIN=johnsmith,janedoe
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
```
