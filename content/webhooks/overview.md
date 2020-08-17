---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
toc: true

aliases:
- /configure-system-webhooks

related:
  items:
  - name: Starter Project
    path: ./tutorial.md
  - name: Sample Payloads
    path: ./examples.md
---

System webhooks can be used to send an http request to a designated endpoint every time a system event occurs. Example system events:

* User is created
* User is deleted
* Repository is activated
* Repository is de-activated
* Build is created
* Build is updated or completed

# Configuration

1. Create a shared secret used to sign the http request. _This shared secret is to be provided to both the Drone server and your webhook receiver._
   ```
   $ openssl rand -hex 16
   bea26a2221fd8090ea38720fc445eca6
   ```

2. Update your Drone server configuration to include the webhook endpoint and the shared secret.
   ```
   DRONE_WEBHOOK_ENDPOINT=http://...
   DRONE_WEBHOOK_SECRET=bea26a2221fd8090ea38720fc445eca6
   ```

# Request

Example webhook request headers:

```http {linenos=table}
POST / HTTP/1.1
Host: www.nowhere123.com
Date: Wed, 01 Jan 2020 00:00:00 GMT
Content-Type: application/json
Content-Length: 358
X-Drone-Event: user
Authorization: hmac-key",algorithm="hmac-sha256",signature="QrS16+RlRsFjXn5IVW8tWz+3ZRAypjpNgzehEuvJksk=
```

Example webhook request body:

```json  {linenos=table}
{
    "action": "created",
    "user": {
        "id": 1,
        "login": "octocat",
        "email": "octocat@github.com",
        "active": true,
        "admin": false,
        "machine": false,
        "syncing": false,
        "synced": 1564096971,
        "created": 1564096971,
        "updated": 1564096971,
        "last_login": 1564096971
    }
}
```

# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Events

The http request includes a header variable to determine he type of event. The following is a list of event types:

```
X-Drone-Event: user
X-Drone-Event: repo
X-Drone-Event: build
```

# Actions

The payload includes an action field, which can be used to determine the type of action that was taken against the object. In the below example we see the user object was created.

```json {linenos=table}
{
    "action": "created",
    "user": {
        "id": 1,
        "login": "octocat",
        "email": "octocat@github.com",
        "active": true,
        "admin": false,
        "machine": false,
        "syncing": false,
        "synced": 1564096971,
        "created": 1564096971,
        "updated": 1564096971,
        "last_login": 1564096971
    }
}
```
