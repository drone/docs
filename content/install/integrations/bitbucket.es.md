+++
date = "2017-04-15T14:39:04+02:00"
title = "Bitbucket Cloud"
url = "es/install-for-bitbucket-cloud"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-bitucket-cloud-es"
  weight = 5
+++

Drone viene con soporte para Bitbucket cloud. Para habilitar Bitbucket cloud debes habilitar el contenedor de Drone con las siguientes variables de ambiente:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_BITBUCKET=true
+     - DRONE_BITBUCKET_CLIENT=95c0282573633eb25e82
+     - DRONE_BITBUCKET_SECRET=30f5064039e6b359e075
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

Esta es la lista completa de opciones de configuración. Por favor nota que muchas de estas opciones usan la configuración por defecto, que deberían funcionar para la mayoría de instalaciones.

DRONE_BITBUCKET=true
: Define que se habilitará el servidor de Bitbucket

DRONE_BITBUCKET_CLIENT
: Llave de cliente oauth2 de Bitbucket

DRONE_BITBUCKET_SECRET
: Llave secreta oauth2 de cliente.

# Registro

Debes registrar tu aplicación con Bitbucket para poder generar una llave de cliente y una llave secreta. Navega en tu configuración de cuenta y escoge OAth desde el menú y luego Agregar Cliente.

Por favor usa la URL de callback:

```nohighlight
http://drone.mycompany.com/authorize
```

Por favor asegúrate de verificar los siguientes permisos:

```nohighlight
Account:Email
Account:Read
Team Membership:Read
Repositories:Read
Webhooks:Read and Write
```
