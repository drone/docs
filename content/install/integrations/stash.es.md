+++
date = "2017-04-15T14:39:04+02:00"
title = "Bitbucket Server"
url = "es/install-for-bitbucket-server"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-bitucket-server-es"
  weight = 6
+++


Drone viene con soporte experimental para el servidor de Bitbucket, Antes conocido como Atlassian Stash. Para habilitar el servidor debes habilitar las siguientes variables de entorno en el contenedor de Drone:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_STASH=true
+     - DRONE_STASH_GIT_USERNAME=foo
+     - DRONE_STASH_GIT_PASSWORD=bar
+     - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
+     - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
    volumes:
+     - /path/to/key.pem:/path/to/key.pem

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
# Archivo de llave privada

El proceso de autenticación con Bitbucket requiere un certificado privado y público RSA. Aquí está como crear un certificado privado RSA.


```nohighlight
openssl genrsa -out /etc/bitbucket/key.pem 1024
```

Esto almacena el certificado privado en `key.pem`. El siguiente comanod genera un certificado público y lo almacena en `key.pub`

```nohighlight
openssl rsa -in /etc/bitbucket/key.pem -pubout >> /etc/bitbucket/key.pub
```

Por favor nota que el archivo de llave privado puede ser montado en tu contenedor de Drone en tiempo de ejecución o como una variable de ambiente.

Montando llave privada por medio de un volumen en tiempo de ejecución:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
    - DRONE_OPEN=true
    - DRONE_HOST=${DRONE_HOST}
      - DRONE_STASH=true
      - DRONE_STASH_GIT_USERNAME=foo
      - DRONE_STASH_GIT_PASSWORD=bar
      - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
      - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
+  volumes:
+     - /etc/bitbucket/key.pem:/etc/bitbucket/key.pem
```

Archivo de llave privada montado como variable de entorno:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
    - DRONE_OPEN=true
    - DRONE_HOST=${DRONE_HOST}
      - DRONE_STASH=true
      - DRONE_STASH_GIT_USERNAME=foo
      - DRONE_STASH_GIT_PASSWORD=bar
      - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA_STRING=contentOfPemKeyAsString
      - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
```

# Servicio de Cuentas

Drone usa `git+https` para clonar repositorios, sin embargo, el servidor Bitbucket no permite clonar repositorios usando un token oauth. Para trabajar con esta limitación, puedes crear un servicio de cuentas y proveer un nombre de usuario y contraseña a Drone. Este servicio de cuenta puede ser usado para clonar repositorios privados.

# Registration

Debes registrar tu aplicación con el servidor Bitbucket para poder generar las llave de cliente y llave secreta. Navega a tu configuración de cuenta y escoge Aplicaciones del menú, luego Registrar nueva aplicación. Ahora copia y pega el valor de texto de `/etc/bitbucket/key.pub` en el campo `Llave pública` en la parte siguiente del registro.

Por favor usa http://drone.mycompany.com/authorize como URL de callback de Autorización.

# Configuración

Esta es una lista completa de todas las opciones de configuración. Por favor nota que estas opciones usan valores por defecto y deberían funcionar con la mayoría de instalaciones.

DRONE_STASH=true
: Habilita el manejador de Servidor Bitbucket (Stash)
: Set to true to enable the Bitbucket Server (Stash) driver.

DRONE_STASH_URL
: Dirección del servidor de Bitbucket

DRONE_STASH_CONSUMER_KEY
: Llave de cliente de Servidor oauth1

DRONE_STASH_CONSUMER_RSA
: Archivo de Llave privada oath1 para el servidor de Bitbucket

DRONE_STASH_CONSUMER_RSA_STRING
: Contenido de Llave privada oath1 para el servidor de Bitbucket

DRONE_STASH_GIT_USERNAME
: Cuenta de usuario para clonar los repositorios.

DRONE_STASH_GIT_PASSWORD
: Contraseña de usuario para clonar los repositorios.
