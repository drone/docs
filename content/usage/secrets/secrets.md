---
title: Manage Secrets
url: manage-secrets

menu:
  usage:
    weight: 3
    identifier: manage_secrets
    parent: usage_secrets
---

Drone provides the ability to store named parameters external to the Yaml configuration file, in a central secret store. Individual steps in the yaml can request access to these named parameters at runtime.

Example configuration using named secrets:

```diff
pipeline:
  docker:
    image: plugins/docker
    repo: octocat/hello-world
    tags: latest
-   username: <username>
-   password: <password>
+   secrets: [ docker_username, docker_password ]
```

The secrets in the above are exposed to the plugin as uppercase environment variables. The plugin will look for the named environment variables at runtime. The variable names are therefore important.

# Add Secrets

Secrets are added to the Drone secret store using the command line utility:

```diff
drone secret add \
  --repository <registry> \
  --image <image> \
  --name <name> \
  --value <value>
```

Example command:

```diff
drone secret add \
  -repository octocat/hello-world \
  -image plugins/docker \
  -name docker_username \
  -value <value>
```

In the above examples we see `docker_username` and `docker_password` secrets. The names of these secret variables are dictated by the plugin. You should therefore consult the individual plugin documentation to determine the exact secret name that should be used.

# Alternate Names

There may be scenarios where you are required to store secrets using alternate names. You can map the alternate secret name to the expected name using the below syntax:

```diff
pipeline:
  docker:
    image: plugins/docker
    repo: octocat/hello-world
    tags: latest
+   secrets:
+     - source: docker_prod_password
+       target: docker_password
```

# Pull Requests

{{% alert error %}}
Expose secrets to pull requests with caution.
{{% /alert %}}

Secrets are not exposed to pull requests by default. You can override this behavior by creating the secret and enabling the `pull_request` event type.

```diff
drone secret add \
  -repository octocat/hello-world \
  -image plugins/docker \
+ -event pull_request \
+ -event push \
+ -event tag \
  -name docker_username \
  -value <value>
```

Please be careful when exposing secrets to pull requests. If your repository is open source and accepts pull requests your secrets are not safe. A bad actor can submit a malicious pull request that exposes your secrets. The recommended approach to secure secrets is to enable gated builds.

# Examples

Create the secret using default settings. The secret will be available to all images in your pipeline, and will be available to all push, tag, and deployment events (not pull request events).

```diff
drone secret add \
  -repository octocat/hello-world \
  -name aws_access_key_id \
  -value <value>
```

Create the secret and limit to a single image:

```diff
drone secret add \
  -repository octocat/hello-world \
+ -image plugins/s3 \
  -name aws_access_key_id \
  -value <value>
```

Create the secrets and limit to a set of images:

```diff
drone secret add \
  -repository octocat/hello-world \
+ -image plugins/s3 \
+ -image peloton/drone-ecs \
  -name aws_access_key_id \
  -value <value>
```

Create the secret and enable for multiple hook events:

```diff
drone secret add \
  -repository octocat/hello-world \
  -image plugins/s3 \
+ -event pull_request \
+ -event push \
+ -event tag \
  -name aws_access_key_id \
  -value <value>
```

# Example From File

Loading secrets from file using curl `@` syntax. This is the recommended approach for loading secrets from file to preserve newlines:

```diff
drone secret add \
  -repository octocat/hello-world \
  -name ssh_key \
+ -value @/root/ssh/id_rsa
```

# Example in Commands

Secrets are exposed to your pipeline steps and plugins as uppercase environment variables and can therefore be referenced in the commands section of your pipeline.

```diff
pipeline:
  docker:
    image: docker
    commands:
+     - echo $DOCKER_USERNAME
+     - echo $DOCKER_PASSWORD
    secrets: [ docker_username, docker_password ]
```

Please note parameter expressions are subject to pre-processing. When using secrets in parameter expressions they should be escaped.

```diff
pipeline:
  docker:
    image: docker
    commands:
-     - echo ${DOCKER_USERNAME}
-     - echo ${DOCKER_PASSWORD}
+     - echo $${DOCKER_USERNAME}
+     - echo $${DOCKER_PASSWORD}
    secrets: [ docker_username, docker_password ]
```
