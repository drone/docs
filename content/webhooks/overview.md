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

# Authorization / HTTP Signature Verification

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

## Implementation Guide 

When a drone webhook request payload comes in, first inspect the `signature` http request header. This header is used conjunction with other headers for performing signature verification.

```
Date: "Fri, 15 Jul 2022 18:47:25 GMT"
Digest: "SHA-256=wyFE2yWKPBpOLHuIVBHf4oD21wY4yINZZzoyR9jB6xo="
Signature: keyId="hmac-key",algorithm="hmac-sha256",signature="ObOcdsOSyYMy+0DDlg6X1naqPYY0qe59OrHmjv6Hav0=",headers="date digest"
```

The `Signature` header contains the instructions for how to validate the incoming request. The `headers` field specifies which HTTP headers to use for computing the final signature. 

To perform the computation, header names and values must be concatenated with a newline, then hashed with the specified `algorithm` and `hmac-key` and finally this result must be base 64 encoded. The result is then compared to the `signature` field from the `Signature` header and if the values match the request is considered valid.

Example ruby code:

```ruby
require  'openssl'
require 'base64' 

key = "a34999ae0599f579eca8582058b46eee" # generated above with `openssl -rand`

# extracted from Signature header and signature field
provided_signature = "ObOcdsOSyYMy+0DDlg6X1naqPYY0qe59OrHmjv6Hav0=" 

# concatenate Date and Digest field with a newline
signature_string = "date: Fri, 15 Jul 2022 18:47:25 GMT\ndigest: SHA-256=wyFE2yWKPBpOLHuIVBHf4oD21wY4yINZZzoyR9jB6xo="

# hash and encode
digest = OpenSSL::HMAC.digest(OpenSSL::Digest::SHA256.new, key, signature_string)
computed_signature = Base64.strict_encode64(digest)

# compare computed value with the value provided in the Signature http header
provided_signature == computed_signature
```

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
