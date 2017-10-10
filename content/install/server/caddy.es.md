+++
date = "2017-04-15T14:39:04+02:00"
title = "Configuración con Caddy"
url = "setup-with-caddy"

[menu.install]
  identifier = "setup-with-caddy"
  parent = "install_server"
  weight = 4
+++

Esta guía provee un vistazo corto a la instalación de un servidor de Drone detrás de un servidor web Caddy. Éste es un archivo de configuración de proxy en caddy:

```nohighlight
drone.mycompany.com {
    gzip {
        not /stream/
    }
    proxy / localhost:8000 {
        websocket
        transparent
    }
}
```

Debe deshabilitar la compresión gzip para datos contínuos de lo contrario las actualizaciones en vivo no serán instantáneas:

```diff
drone.mycomopany.com {
+   gzip {
+       not /stream/
+   }
    proxy / localhost:8000 {
        websocket
        transparent
    }
}
```

Debes configurar el proxy para habilitar la actualización de sockets web:

```diff
drone.mycompany.com {
    gzip {
        not /stream/
    }
    proxy / localhost:8000 {
+       websocket
        transparent
    }
}
```

Debes configurar el proxy para incluir cabeceras `X-Forwarded` usando la directiva `transparent`:

```diff
drone.mycompany.com {
    gzip {
        not /stream/
    }
    proxy / localhost:8000 {
        websocket
+       transparent
    }
}
```
