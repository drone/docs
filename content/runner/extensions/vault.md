---
date: 2000-01-01T00:00:00+00:00
title: Vault Secrets
author: bradrydzewski
weight: 60
description: |
  Install the Vault secret management extension.


aliases:
- /extend/secrets/vault
- /extend/secrets/vault/install
- /extend/secrets/vault/config-in-vault
- /extend/secrets/vault/config-in-drone
- /setup-vault-integration
---

Vault secures, stores, and tightly controls access to tokens, passwords, certificates, keys, and other secrets in modern computing. The Vault extension provides your pipeline with access to Vault secrets.

# Installation

1. Create a shared secret.
   ```
   $ openssl rand -hex 16
   bea26a2221fd8090ea38720fc445eca6
   ```

2. Download and run the extension.
   ```
   $ docker run -d \
     --publish=3000:3000 \
     --env=DRONE_DEBUG=true \
     --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
     --env=VAULT_ADDR=... \
     --env=VAULT_TOKEN=... \
     --restart=always \
     --name=secrets \
     drone/vault
   ```

3. Update your Drone runner configuration to include the extension address and the shared secret.
   ```
   DRONE_SECRET_PLUGIN_ENDPOINT=http://1.2.3.4:3000
   DRONE_SECRET_PLUGIN_TOKEN=bea26a2221fd8090ea38720fc445eca6
   ```

# Verification

You can verify the extension is configured and is processing requests using the command line utility.

1. Provide the command line utility with the extension endpoint and secret.
   ```
   $ export DRONE_SECRET_ENDPOINT=http://1.2.3.4:3000
   $ export DRONE_SECRET_SECRET=bea26a2221fd8090ea38720fc445eca6
   ```

2. Use the command line utility to retrieve the secret:
   ```
   $ drone plugins secret get secrets/data/docker username
   ```

