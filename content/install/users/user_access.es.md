+++
date = "2017-04-15T14:39:04+02:00"
title = "Registro de usuarios"
url = "es/user-registration"

[menu.install]
  weight = 1
  identifier = "user-registration-es"
  parent = "install_access"
+++

Drone provee múltiples configuraciones para acceso abierto o limitado al sistema. En esta sección se describen diferentes opciones para el registro de usuarios y el acceso.

# Registro abierto

El registro abierto es recomendado únicamente para instalaciones seguras en una red privada. Esta configuración le permite a cualquier persona registrarse e ingresar en el sistema.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_OPEN= true
```

# Registro restringido

El registro restringido es la configuración recomendada. Esta configuración le permite a los miembros de una lista de organizaciones registrarse a si mismos para ingresar al sistema.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     DRONE_OPEN: true
+     DRONE_ORGS: dogpatch,dolores
```

# Registro cerrado

El registro cerrado está habilitado por defecto. El registro cerrado requiere que el administrador manualmente registro usuarios usando la utilidad de linea de comandos. Cuando usamos registro cerrado es __imperativo__ proveer una lista de administradores que pueden ingresar para manejar cuentas.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
-     - DRONE_OPEN=true
+     - DRONE_OPEN=false
+     - DRONE_ADMIN=janedoe,johnsmith
```

Puedes manualmente conceder acceso a los usuarios usando la utilidad de línea de comandos.

```nohighlight
drone user add <username>
```
