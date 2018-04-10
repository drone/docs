+++
date = "2017-04-15T14:39:04+02:00"
title = "Configurar SSL"
url = "es/configure-ssl"

[menu.install]
  identifier = "configure-ssl-es"
  parent = "install_server"
  weight = 10
+++

Drone soporta configuración SSL montando certificados en tu contenedor. Nota que también existe soporte para [Let's Encrypt]({{< relref "lets_encrypt.es.md" >}}) automatizado.

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
+     - /etc/certs/drone.foo.com/server.crt:/etc/certs/drone.foo.com/server.crt
+     - /etc/certs/drone.foo.com/server.key:/etc/certs/drone.foo.com/server.key
    restart: always
    environment:
+     - DRONE_SERVER_CERT=/etc/certs/drone.foo.com/server.crt
+     - DRONE_SERVER_KEY=/etc/certs/drone.foo.com/server.key
```

Actualiza tu configuración para exponer los siguientes puertos:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
+     - 80:80
+     - 443:443
      - 9000:9000
```

Actualiza tu configuración para montar tu certificado y llave:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:80
      - 443:443
      - 9000:9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
+     - /etc/certs/drone.foo.com/server.crt:/etc/certs/drone.foo.com/server.crt
+     - /etc/certs/drone.foo.com/server.key:/etc/certs/drone.foo.com/server.key
```

Actualiza tu configuración para proveer las rutas de acceso de tu certificado y llave:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:80
      - 443:443
      - 9000:9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
      - /etc/certs/drone.foo.com/server.crt:/etc/certs/drone.foo.com/server.crt
      - /etc/certs/drone.foo.com/server.key:/etc/certs/drone.foo.com/server.key
    restart: always
    environment:
+     - DRONE_SERVER_CERT=/etc/certs/drone.foo.com/server.crt
+     - DRONE_SERVER_KEY=/etc/certs/drone.foo.com/server.key
```

# Cadena de certificados

El problema más común es proveer un certificado sin la cadena intermediaria.

> LoadX509KeyPair lee y analiza un par de llaves públicas/privadas de un par de archivos. Los archivos deben contener datos codificados en PEM. El archivo de certificado puede contener certificados intermedios después del certificado hoja resultando en una cadena de certificados.

# Errores de certificado

El soporte SSL es proveído usando la función [ListenAndServeTLS](https://golang.org/pkg/net/http/#ListenAndServeTLS) de la librería estándar de Go. Si recibes errores o advertencias de certificado por favor examina tu configuración a más detalle. Por favor no crees issues mencionando que SSL está roto.
