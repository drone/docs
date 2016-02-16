+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Secrets"
weight = 3
menu = "usage"
toc = true
+++

> WARNING: drone does not prevent you from inadvertently exposing your secrets to the world. Please read this section of the documentation carefully to ensure you follow best practices and avoid exposing sensitive data.


# Overview

Drone lets you store secret variables in an encrypted `.drone.sec` file in the root of your repository. This is useful when your build requires sensitive information that should not be stored in plaintext in your yaml file. This document assumes you have installed the Drone [command line tools](/devs/cli).

Start with a plaintext YAML file that defines your secrets. For demonstration purposes let's assume this file is stored on disk and named `secrets.yml`. Secrets are defined in the `environment` section of this file:

```yaml
---
environment:
  HEROKU_TOKEN: pa$$word
```

Reference secrets in your `.drone.yml` file using the `$$` notation:

```yaml
---
deploy:
  heroku:
    app: petstore
    token: $$HEROKU_TOKEN
```

Encrypt the `secrets.yml` file and generate a `.drone.sec` file:

```
drone secure --repo octocat/hello-world --in secrets.yml
```

Commit the encrypted `.drone.sec` file to the root of your repository:

```
git add .drone.sec
git commit -m "added .drone.sec file"
```

# Interpolation

Secrets are injected directly into the YAML at runtime using the `$$` notation:

```yaml
---
deploy:
  heroku:
    app: foo
    token: $$HEROKU_TOKEN
```

Note that secrets are not automatically exposed to your build environment as
environment variables. For security purposes, you be must explicit when and
where to inject secrets.

This example demonstrates explicitly injecting a secret as an environment variable:

```yaml
---
build:
  image: golang
  commands:
    - go build
    - go test
  environment:
    - PRIVATE_KEY=$$PRIVATE_KEY
```

Secrets that are multiple lines or contain special characters should be escaped to
avoid YAML parsing errors post-interpolation. You can escape strings by wrapping
the variable in quotes as shown below:

```yaml
---
build:
  image: golang
  commands:
    - go build
    - go test
  environment:
    - PRIVATE_KEY="$$PRIVATE_KEY"
```

# Checksums

Drone automatically calculates and stores a checksum of your `.drone.yml` file inside your `.drone.sec` file. Secrets are not injected into your build if the checksum cannot be verified. This means you must re-generate your `.drone.sec` file every time you `.drone.yml` changes (yes security is inconvenient).

Invalid checksums result in the following error at the top of your build logs:

```
Unable to validate YAML checksum.
2ca66eb7be89f31afdebb197174abfa6dd14866ecbf9e552f44be5bd3244d08a
```

# Pull Requests

Drone injects secrets for all build types, including pull requests, as long as the checksum validation passes. You should avoid injecting secrets into the build section of the YAML and inadvertently exposing your secrets to malicious pull requests.

**This is bad** because a malicious pull request could gain access to your secrets:

```yaml
---
build:
  image: node
  environment:
    PASSWORD: $$PASSWORD
  commands:
    - npm install
    - npm run tests
    - npm run integration
```

**This is better** because the build is split into sections. Secrets are injected into a section of the build that is not executed for pull requests:

```yaml
---
build:
  unit_tests:
    image: node
    commands:
      - npm install
      - npm run tests

  integration_tests:
    image: node
    environment:
      PASSWORD: $$PASSWORD
    commands:
      - npm run integration
    when:
      event: push
```


# Global Secrets

Global secrets are stored in the `PLUGIN_PARAMS` environment variable declared in your Drone server configuration file. Global secrets are passed to every single build and should therefore only be used in trusted environments with trusted developers.

Example global secret declaration:

```bash
PLUGIN_PARAMS=SMTP_PASSWORD=foo SLACK_TOKEN=bar
```

Example injecting global secrets into your `.drone.yml` file:

```yaml
---
deploy:
  slack:
    channel: foo
    token: $$SLACK_TOKEN
```


# Common Issues

Secrets are not injected into your build if the checksum cannot be validated. This happens when you change your `.drone.yml` file without re-generating a `.drone.sec` file resulting in the following error message at the top of your build logs:

```
Unable to validate YAML checksum.
2ca66eb7be89f31afdebb197174abfa6dd14866ecbf9e552f44be5bd3244d08a
```
