+++
date = "2017-04-15T14:39:04+02:00"
title = "Configurar Lets Encrypt"
url = "es/configure-lets-encrypt"

[menu.install]
  identifier = "configure-lets-encrypt-es"
  parent = "install_server"
  weight = 9
+++

Drone soporta la configuración y actualización automatizada de SSL usando Let's Encrypt. Puedes habilitar Let's Encrypt haciendo las siguientes modificaciones a la configuración de tu servidor:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
+     - 80:80
+     - 443:443
      - 9000:9000
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
+     - DRONE_LETS_ENCRYPT=true
```

Nota que Drone usa el hostname de la variable de entorno `DRONE_HOST`cuando se soliciten los certificados. Por ejemplo, para `DRONE_HOST=https://foo.com` el certificado solicitado será para el dominio `foo.com`.

<!-- Once enabled you can visit your website at both the http and the https address. There are no immediate plans to redirect from http to https, but it may be considered for a future release. -->

# Caché de certificados

Drone almacenará los certificados al siguiente directorio:

```
/var/lib/drone/golang-autocert
```

# Actualizaciones de certificados

Drone usa la librería oficial de Go ACME que maneja las actualizaciones de certificados. No debería requerirse configuración o manejo adicional.
