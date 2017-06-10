+++
date = "2017-04-15T14:39:04+02:00"
title = "Gogs"
url = "es/install-for-gogs"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-gogs-es"
  weight = 3
+++

Drone viene con soporte para la última versión estable de Gogs. Para habilitar Gogs debes configurar el contenedor de Drone las siguientes variables de ambiente:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_GOGS=true
+     - DRONE_GOGS_URL=http://gogs.mycompany.com
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

# Autenticación

Drone pedirá nombre de usuario y contraseña para autenticarse. Deberías usar tu nombre de usuario y contraseña de Gogs. Esto es requerido desafortunadamente pues Gogs no tiene soporte para oauth2.

# Configuración

Esta es la lista completa de opciones de configuración. Por favor note que muchas de estas opciones usan los valores de configuración por defecto que deberían funcionar para la mayoría de instalaciones.

DRONE_GOGS=true
: Habilita el manejador de Gogs.

DRONE_GOGS_URL
: Servidor URL de Gogs.

DRONE_GOGS_GIT_USERNAME
: Opcional. Usa una cuenta de usuario única para clonar todos los repositorios.

DRONE_GOGS_GIT_PASSWORD
: Opcional. Usa una contraseña de usuario única para clonar todos los repositorios.

DRONE_GOGS_PRIVATE_MODE=false
: Habilita que Gogs esté ejecutándose en modo privado.

DRONE_GOGS_SKIP_VERIFY=false
: Deshabilita la verificación SSL.
