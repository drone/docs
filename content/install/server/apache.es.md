+++
date = "2017-04-15T14:39:04+02:00"
title = "Configuración con Apache"
url = "setup-with-apache"

[menu.install]
  identifier = "setup-with-apache"
  parent = "install_server"
  weight = 5
+++

Ésta guía provee una vista corta para instalar un servidor de Drone detrás de un servidor Apache2. Ésta es una configuración de ejemplo:

```nohighlight
ProxyPreserveHost On

RequestHeader set X-Forwarded-Proto "https"

ProxyPass /ws/ ws://localhost:8000/ws/
ProxyPassReverse /ws/ ws://localhost:8000/ws/

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```

Debes tener los siguientes módulos de Apache instalados:

```nohighlight
a2enmod proxy
a2enmod proxy_http
a2enmod proxy_wstunnel
```

Debes configurar Apache con `X-Forwarded-Proto` cuando se use https.

```diff
ProxyPreserveHost On

+RequestHeader set X-Forwarded-Proto "https"

ProxyPass /ws/ ws://localhost:8000/ws/
ProxyPassReverse /ws/ ws://localhost:8000/ws/

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```

Debes configurar Apache para habilitar actualizaciones de sockets web.

```diff
ProxyPreserveHost On

RequestHeader set X-Forwarded-Proto "https"

+ProxyPass /ws/ ws://localhost:8000/ws/
+ProxyPassReverse /ws/ ws://localhost:8000/ws/

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```
