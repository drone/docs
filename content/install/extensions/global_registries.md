+++
date = "2017-04-15T14:39:04+02:00"
title = "Global Registries File"
url = "setup-global-registry-credentials"

[menu.install]
  parent = "install_enterprise"
  identifier = "setup-global-registry-credentials"
  weight = 3
+++

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise expansion pack provides support for global registry credentials, sourced from a yaml file on your server. You should mount the credentials file into your container and specify the path to the file in your configuration.

Example server configuration:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
+     - /etc/drone-registry.yml:/etc/drone-registry.yml
    restart: always
    environment:
+     DRONE_GLOBAL_REGISTRY=/etc/drone-registry.yml
```

Example registry credentials file:

```nohighlight
- address: docker.io
  username: octocat
  password: correct-horse-batter-staple
- address: gcr.io
  username: _json_key
  password: |
    {
      "private_key_id": "...",
      "private_key": "...",
      "client_email": "...",
      "client_id": "...",
      "type": "..."
    }
```

# Elastic Container Registry (ECR)

The global registry file supports elastic container registries. Temporary authentication credentials are automatically generated and refreshed using the GetAuthorizationToken endpoint. Note that you must provide your aws access key and secret, and configure the appropriate IAM roles.

Example registry credentials for an ECR repository:

```nohighlight
- address: 012345678910.dkr.ecr.us-east-1.amazonaws.com
  aws_access_key_id: a50d28f4dd477bc184fbd10b376de753
  aws_secret_access_key: bc5785d3ece6a9cdefa42eb99b58986f9095ff1c
```

# Restricting Access

Currently, global registry credentials do not support any attribute-based usage restriction (repo, images, events).

This is because registry credentials are internal-only to Drone, and unlike secrets, are never exposed to the build.
