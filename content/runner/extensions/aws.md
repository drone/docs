---
date: 2000-01-01T00:00:00+00:00
title: AWS Secrets
author: bradrydzewski
weight: 50
aliases:
- /extend/secrets/amazon
- /extend/secrets/amazon/install
- /extend/secrets/amazon/config-in-aws
- /extend/secrets/amazon/config-in-drone

description: |
  Install the AWS secret management extension.
---

The AWS Secrets Manager secures, stores, and controls access to tokens, passwords, certificates, and other secrets in modern computing. The AWS Secrets Manager extension provides your pipeline with access to AWS secrets.

# Installation

1. Create a shared secret key.
   ```
   $ openssl rand -hex 16
   bea26a2221fd8090ea38720fc445eca6
   ```

2. Download and run the extension.
   ```
   $ docker run -d \
     --publish=3000:3000 \
     --env=DEBUG=true \
     --env=SECRET_KEY=bea26a2221fd8090ea38720fc445eca6 \
     --env=AWS_ACCESS_KEY_ID=... \
     --env=AWS_SECRET_ACCESS_KEY=... \
     --restart=always \
     --name=secrets \
     drone/amazon-secrets
   ```

3. Update your Drone runner configuration to include the extension address and the shared secret key.
   ```
   DRONE_SECRET_PLUGIN_ENDPOINT=http://1.2.3.4:3000
   DRONE_SECRET_PLUGIN_TOKEN=bea26a2221fd8090ea38720fc445eca6
   ```

# Verification

You can verify the extension is configured and is processing requests using the command line utility.

1. Provide the command line utility with the endpoint and secret key:
   ```
   $ export DRONE_SECRET_ENDPOINT=http://1.2.3.4:3000
   $ export DRONE_SECRET_SECRET=bea26a2221fd8090ea38720fc445eca6
   ```

2. Use the command line utility to retrieve an existing secret:
   ```
   $ drone plugins secret get prod/docker username
   ```

