---
date: 2000-01-01T00:00:00+00:00
title: Cookies
author: bradrydzewski
weight: 21
toc: true
description: |
  Configure server cookies.
---

Drone uses cookies to authenticate client requests and maintain the user session. This document provides resources for customizing cookie creation.

# Signature

Drone signs authentication cookies using a secret key that is randomly generated when the server starts. Every time you restart the server a new key is generated, which means existing sessions are terminated. We can avoid terminating existing sessions on restart by providing the server with a static secret key.

1. Create a secret key:
   ```
   $ openssl rand -hex 16
   bea26a2221fd8090ea38720fc445eca6
   ```

2. Provide the server with the secret key:
   ```
   $ docker run \
     -e DRONE_COOKIE_SECRET=bea26a2221fd8090ea38720fc445eca6
   ```

# Expiration

Drone creates cookies with a default timeout of 30 days (_720h_). When the cookie expires the user sessions is terminated and the user will be required to login. You can customize the cookie timeout through the environment.

```
DRONE_COOKIE_TIMEOUT=720h
```
