---
date: 2017-04-15T14:39:04+02:00
title: "Error: No Space Left on Device"
url: error-no-space-left-on-device
---

This error is encountered when your host machine runs out of disk space. Please note that this is not considered a drone bug, but can certainly be a side-effect of a very active build server.

```nohighlight
no space left on device
```

The most common root cause for this error message is that your docker daemon is caching a large number of images and has exhausted disk space. We recommend regularly [pruning](https://docs.docker.com/engine/reference/commandline/image_prune/) cached images.

```nohighlight
docker images prune
```

# Automated Tooling

There are tools you can use to automatically purge old and unused images. We recommend running [docker-custodian](https://github.com/Yelp/docker-custodian) alongside your agent.

```yaml
version: '2'

services:
  docker-custodian:
    image: yelp/docker-custodian
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    restart: always
    command: [ dcgc, --max-image-age, 30days ]
```

Note the above configuration is a sample and should be tuned for your environment, with [pattern matching](https://github.com/Yelp/docker-custodian#prevent-images-from-being-removed) configured to prevent pruning commonly used images.
