---
date: 2000-01-01T00:00:00+00:00
title: Certificates
author: bradrydzewski
weight: 20
separator: true
aliases:
- /installation/security/ssl/
- /installation/security/ssl-lets-encrypt/
- /configure-ssl/
- /configure-lets-encrypt/
- /administration/server/ssl
- /administration/server/ssl-with-lets-encrypt

description: |
  Configure server security.
---

Drone supports native SSL configuration by mounting certificates into the server container. _If your server is public you should consider using Lets Encrypt._

1. Mount your certificate and key into the server container:

    ```
    $ docker run \
    -v /etc/certs/drone.company.com/server.crt:/etc/certs/drone.company.com/server.crt \
    -v /etc/certs/drone.company.com/server.key:/etc/certs/drone.company.com/server.key
    ```

2. Configure the path to your certificate and key:

    ```
    $ docker run \
    -e DRONE_TLS_CERT=/etc/certs/drone.company.com/server.crt \
    -e DRONE_TLS_KEY=/etc/certs/drone.company.com/server.key
    ```

3. Expose the standard http and https ports:

    ```
    $ docker run \
    -p 80:80 \
    -p 443:443
    ```

# Lets Encrypt

Drone supports automated SSL configuration and updates using Let's Encrypt. You can enable Let’s encrypt with the following flag:

1. Enable Lets Encrypt with the following parameter:
    ```
    DRONE_TLS_AUTOCERT=true
    ```

2. Ensure the desired hostname is configured:

    ```
    DRONE_SERVER_HOST=domain.com
    DRONE_SERVER_PROTO=https
    ```

3. Expose the standard http and https ports:

    ```
    docker run \
    -p 80:80 \
    -p 443:443
    ```

4. Mount the certificate cache to the host:

    ```
    docker run \
    -v /var/lib/drone:/data
    ```

## Certificate Cache

Drone caches generated certificates on disk at `/data/golang-autocert`. This prevents the system from re-requesting certificates on restart. It is best practice to bind mount the `/data` directory to the host.

## Certificate Upgrades

Drone uses the official Go acme library which will handle certificate upgrades. There should be no additional configuration or management required.