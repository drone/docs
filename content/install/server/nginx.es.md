+++
date = "2017-04-15T14:39:04+02:00"
title = "Configuración con Nginx"
url = "es/setup-with-nginx"

[menu.install]
  identifier = "setup-with-nginx-es"
  parent = "install_server"
  weight = 3
+++

Esta guía provee un vistazo básico para instalar un servidor de Drone detrás de un servidor web Nginx. Para obtener opciones de configuración más avanzadas por favor consulta la [documentación oficial](https://www.nginx.com/resources/admin-guide/) de Nginx.

Configuración de ejemplo:

```nginx
server {
    listen 80;
    server_name drone.example.com;

    location / {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Host $http_host;

        proxy_pass http://127.0.0.1:8000;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_buffering off;

        chunked_transfer_encoding off;
    }
}
```
