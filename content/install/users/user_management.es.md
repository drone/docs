+++
date = "2017-04-15T14:39:04+02:00"
title = "Manejo de usuarios"
url = "es/user-management"

[menu.install]
  weight = 2
  identifier = "user-management-es"
  parent = "install_access"
+++

Puedes manejar el registro de usuarios manualmente usando la utilidad de línea de comandos. Por favor mira la documentación de la utilidad de línea de comandos para instalar y configurar la utilidad de línea de comandos.

Usa el comando `ls` para listar todos los usuarios activos:

```nohighlight
drone user ls
```

Usa el comando `add` para agregar usuarios al sistema por nombre de usuario:

```nohighlight
drone user add octocat
```

Usa el comando `rm` para eliminar usuarios del sistema por nombre de usuario:

```nohighlight
drone user rm octocat
```

Por favor nota que solamente los administradores pueden manejar usuarios. En el ejemplo que mostraremos a continuación configuraremos varios administradores separados por coma, usando la variable de ambiente correspondiente.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_ADMIN=janedoe,johnsmith
```
