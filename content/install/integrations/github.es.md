+++
date = "2017-04-15T14:39:04+02:00"
title = "GitHub"
url = "es/install-for-github"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-github-es"
  weight = 1
+++

Drone incluye soporte para Github y Github Enterprise. Para habilitar Github deberías configurar el contenedor de Drone usando las siguientes variables de ambiente:

```diff
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
+     - DRONE_GITHUB=true
+     - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
+     - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
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

# Registro

Registrar tu aplicación con github para crear tu llave de cliente y llave secreta. Es muy importante que la URL de callback de autorización tenga el mismo esquema (http(s)) y nombre de host con la ruta `authorize`.

Por favor usa esta captura de pantalla como referencia:

![configuración oauth de github](images/github_oauth.png)

# Configuración

Esta es la lista completa de opciones de configuración. Por favor nota que muchas de estas opciones usan una configuración por defecto que deberían funcionar para la mayoría de instalaciones.

DRONE_GITHUB=true
: Define que el manejador de Github estará activado.

DRONE_GITHUB_URL=`https://github.com`
: URL del servidor de Github

DRONE_GITHUB_CLIENT
: Llave de cliente oauth2 de Github.

DRONE_GITHUB_SECRET
: Llave secreta  de oauth2 de Github.

DRONE_GITHUB_SCOPE=repo,repo:status,user:email,read:org
: Lista separada por comas con el alcance de github.

DRONE_GITHUB_GIT_USERNAME
: Opcional. Usa una cuenta única para clonar todos los repositorios.

DRONE_GITHUB_GIT_PASSWORD
: Opcional. Usa una contraseña de un usuario github para clonar todos los repositorios.

DRONE_GITHUB_PRIVATE_MODE=false
: Define si Github está ejecutándose en modo privado.

DRONE_GITHUB_MERGE_REF=true
: Define que se use `refs/pulls/%d/merge` en vez de `refs/pulls/%d/head`

DRONE_GITHUB_CONTEXT=continuous-integration/drone
: Personaliza el mensaje de estado de Github.

DRONE_GITHUB_SKIP_VERIFY=false
: Deshabilita la verificación SSL.
