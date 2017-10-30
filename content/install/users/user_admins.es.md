+++
date = "2017-04-15T14:39:04+02:00"
title = "Usuarios Administradores"
url = "es/user-admins"

[menu.install]
  weight = 3
  identifier = "user-admins-es"
  parent = "install_access"
+++

Puedes conceder privilegios administrativos a usuarios proveyendo una lista separada por comas de usuarios, usando la variable de ambiente designada.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_ADMIN=janedoe,johnsmith
```

Por favor nota que los nombre de usuarios son sensibles entre mayúsculas y minúsculas, y deben coincidir exactamente con los valores retornados con tu sistema de control de versiones (Ej Github).
