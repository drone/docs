+++
date = "2017-04-15T14:39:04+02:00"
title = "GitLab"
url = "es/install-for-gitlab"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-gitlab-es"
  weight = 2
+++

Drone viene con soporte para Gitlab versión 8.2 y superios. Para habilitar Gitlab deberías configurar las siguientes variables de entorno en el contenedor de drone:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_GITLAB=true
+     - DRONE_GITLAB_CLIENT=95c0282573633eb25e82
+     - DRONE_GITLAB_SECRET=30f5064039e6b359e075
+     - DRONE_GITLAB_URL=http://gitlab.mycompany.com
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

# Configuración

Esta es una lista completa de opciones de configuración. Por favor nota que muchas de estas opciones utilizan valores por defecto que deberían funcionar para la mayoría de instalaciones.

DRONE_GITLAB=true
: Se define para habilitar el manejador de Gitlab.

DRONE_GITLAB_URL=`https://gitlab.com`
: Dirección del servidor Gitlab.

DRONE_GITLAB_CLIENT
: Llave de cliente oauth2 de Gitlab.

DRONE_GITLAB_SECRET
: Llave secreta de oath2

DRONE_GITLAB_GIT_USERNAME
: Opcional. Usa un nombre de usuario único para clonar todos los repositorios.

DRONE_GITLAB_GIT_PASSWORD
: Opcional. Usa una contraseña de usuario única para clonar todos los repositorios.

DRONE_GITLAB_SKIP_VERIFY=false
: Habilita la verificación SSL.

DRONE_GITLAB_PRIVATE_MODE=false
: Habilita que Gitlab esté ejecutándose en modo privado.

# Registro

Debes registrar tu aplicación con Gitlab para poder generar una llave de cliente y una llave secreta. Navega en tu Configuración de Cuenta y escoge Aplicaciones y luego Nueva aplicación.

Por favot usa `http://drone.mycompany.com/authorize` como URL de callback de autorización. 
