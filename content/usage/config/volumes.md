+++
date = "2017-04-15T14:39:04+02:00"
title = "Volumes"
url = "docker-volumes"

[menu.usage]
  weight = 6
  identifier = "volumes"
  parent = "usage_config"
+++

Drone gives the ability to define Docker volumes in the Yaml. You can use this parameter to mount files or folders on the host machine into your containers.

{{% alert error %}}
Volumes are only available to trusted repositories and for security reasons should only be used in private environments.
{{% /alert %}}

```diff
pipeline:
  build:
    image: docker
    commands:
      - docker build --rm -t octocat/hello-world .
      - docker run --rm octocat/hello-world --test
      - docker push octocat/hello-world
      - docker rmi octocat/hello-world
    volumes:
+     - /var/run/docker.sock:/var/run/docker.sock
```

Please note that Drone mounts volumes on the host machine. This means you must use absolute paths when you configure volumes. Attempting to use relative paths will result in an error.

```diff
- volumes: [ ./certs:/etc/ssl/certs ]
+ volumes: [ /etc/ssl/certs:/etc/ssl/certs ]
```

## Global volume

You may want to provide the ``DRONE_VOLUME`` variable to your server configuration and case if you would like to mount a host directory into all containers.

Example:

```
DRONE_VOLUME=/tmp/drone-cache:/cache
```

This will mount `/tmp/drone-cache` on the host into all containers at path `/cache`. This includes plugin containers.

Pros:
- The volume will be available for all containers in the pipeline;

Cons:
- Multiple containers may write to the volume at the same time which will likely cause data corruption;
- Sensitive data in the volume will be exposed to any containers in the pipeline;
