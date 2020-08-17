---
date: 2000-01-01T00:00:00+00:00
title: Kubernetes Secrets
author: bradrydzewski
weight: 61
description: |
  Install the Kubernetes secret extension.

aliases:
- /extend/secrets/kubernetes
- /extend/secrets/kubernetes/install
- /extend/secrets/kubernetes/config-in-kubernetes
- /extend/secrets/kubernetes/config-in-drone
---

The Kubernetes Secret resource secures, stores, and controls access to tokens, passwords, certificates, and other secrets in modern computing. The Kubernetes Secrets extension provides your pipeline with access to Kubernetes secrets.

# Installation

1. Create a shared secret key.
   ```
   $ openssl rand -hex 16
   bea26a2221fd8090ea38720fc445eca6
   ```

2. Deploy the secret extension in the same Pod as your Kubernetes runner. _This manifest is for illustrative purpose only and is not intended to be used as-is._ 
    ```
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: drone
      labels:
        app.kubernetes.io/name: drone
    spec:
      replicas: 1
      template:
        spec:
          containers:
          - name: secrets
            image: drone/kubernetes-secrets:latest
            ports:
            - containerPort: 3000
            env:
            - name: SECRET_KEY
              value: bea26a2221fd8090ea38720fc445eca6
    ```

3. Update your Drone runner configuration to include the extension address and the shared secret key.
   ```
   DRONE_SECRET_PLUGIN_ENDPOINT=http://...:3000
   DRONE_SECRET_PLUGIN_TOKEN=bea26a2221fd8090ea38720fc445eca6
   ```

# Verification

You can verify the extension is configured and is processing requests using the command line utility.

1. Provide the command line utility with the extension endpoint and secret key.
   ```
   $ export DRONE_SECRET_ENDPOINT=http://...:3000
   $ export DRONE_SECRET_SECRET=bea26a2221fd8090ea38720fc445eca6
   ```

2. Use the command line utility to retrieve an existing secret:
   ```
   $ drone plugins secret get docker username
   ```
