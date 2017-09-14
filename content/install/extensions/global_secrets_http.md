+++
date = "2017-04-15T14:39:04+02:00"
title = "Global Secrets Endpoint"
url = "setup-global-secrets-endpoint"

[menu.install]
  parent = "install_enterprise"
  identifier = "setup-global-secrets-endpoint"
  weight = 2
+++

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise expansion pack provides support for global secret variables, sourced from an http endpoint. This gives you the ability to create your own secret server and provide secrets to your builds using highly customized logic.

Example server configuration:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    restart: always
    environment:
+     DRONE_SECRET_ENDPOINT=http://localhost:4567/some/endpoint
```

Example request:

```nohighlight
POST /some/endpoint HTTP/1.1

Host: localhost:4567
Content-Type: application/json

{
  "action": "created",
  "repo": {
    "id": 1,
    "owner": "octocat",
    "name": "hello-world",
    "full_name": "octocat/hello-world",
    "clone_url": "https://github.com/octocat/hello-world.git",
    "default_branch": "master"
  },
  "build": {
    "id": 1,
    "number": 1,
    "event": "push",
    "status": "pending"
    "commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
    "branch": "master",
    "ref": "refs/heads/master",
    "author": "Spaceghost",
    "author_email": "octocat@github.com"
  }
}
```

Example response:

```json
[
  {
    "name": "docker_username",
    "value": "****",
    "event": [ "push", "tag" ]
  },
  {
    "name": "docker_password",
    "value": "****",
    "event": [ "push", "tag" ]
  }
]
```

# Request

The drone sever will make an http POST request at runtime to the configured endpoint. The request body will contain the repository and build details.

Example request body:

```
{
  "action": "created",
  "repo": {
    "id": 1,
    "owner": "octocat",
    "name": "hello-world",
    "full_name": "octocat/hello-world",
    "clone_url": "https://github.com/octocat/hello-world.git",
    "default_branch": "master"
  },
  "build": {
    "id": 1,
    "number": 1,
    "event": "push",
    "status": "pending",
    "commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
    "branch": "master",
    "ref": "refs/heads/master",
    "author": "Spaceghost",
    "author_email": "octocat@github.com"
  }
}
```

# Response

The endpoint must response with an array of secret objects, json encoded in the response body. Example response body:

```json
[
  {
    "name": "docker_username",
    "value": "****",
    "event": [ "push", "tag" ]
  },
  {
    "name": "docker_password",
    "value": "****",
    "event": [ "push", "tag" ]
  }
]
```

If no secrets are being returned the endpoint must respond with an empty array.

# Precedence

Please note that secrets stored at the repository level will always take precedence over global secrets.
