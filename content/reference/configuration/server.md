+++
date = "2015-12-05T16:00:21-08:00"
title = "Server Config"
weight = 1

[menu.main]
	parent="configuration"
+++

Below is a complete reference for all Drone Server settings. These can all be passed via environment variables, as shown in the *NAME* column.

NAME                        | DESC
----------------------------|--------------------------------------------------------
`DRONE_SECRET`              | Shared secret. Your Drone Agents must match this value.
`DRONE_DEBUG`               | Set to `true` to see verbose debugging log output.
`DRONE_SERVER_ADDR`         | HTTP server bind address. (default: `:8000`)
`DRONE_SERVER_CERT`         | Optional. Full path to SSL certificate for HTTPS.
`DRONE_SERVER_KEY`          | Optional. Full path to private key for HTTPS.
`DRONE_ADMIN`               | Comma-separated of usernames to grant admin privs to. See the Installation ([Users](../../../installation/users)) documentation.
`DRONE_ORGS`                | If specified, only members of the orgs in this comma-separated list will be able to login. 
`DRONE_OPEN`                | If `true`, user registration is open to everyone.
`DRONE_YAML`                | Changes the default build config file name. (default: `.drone.yml`)
`DATABASE_DRIVER`           | One of `sqlite3`, `postgres`, or `mysql`.
`DATABASE_CONFIG`           | See the Installation ([Database](../../../installation/database)) documentation. 
