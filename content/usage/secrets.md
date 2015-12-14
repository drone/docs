+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Secrets"
weight = 3
menu = "usage"
toc = true
+++

# Overview


# Encryption

# Checksum

# Interpolation

Secrets are injected directly into the Yaml at runtime using the `$$` notation:

```
deploy:
  heroku:
    app: foo
    token: $$HEROKU_TOKEN
```

Note that secrets are not automatically exposed to your build environment as
environment variables. For security purposes, you be must explicit when and
where to inject secrets.

This example demonstrates explicitly injecting a secret as an environment variable:

```
build:
  image: golang
  commands:
    - go build
    - go test
  environment:
    - PRIVATE_KEY=$$PRIVATE_KEY
```

Secrets that are multiple lines or contain special characters should be escaped to
avoid Yaml parsing errors post-interpolation. You can escape strings by wrapping
the variable in quotes as shown below:

```
build:
  image: golang
  commands:
    - go build
    - go test
  environment:
    - PRIVATE_KEY="$$PRIVATE_KEY"
```

# Pull Requests

Secret variables are not injected into to the `build` section of the `.drone.yml` if your repository is public and the build is a pull request. This is for security purposes to prevent a malicious pull request from leaking your secrets.

# Globals

Global secrets are stored in the `PLUGIN_PARAMS` environment variable declared in your Drone server configuration file. Note that global secrets are passed to every single build and should not be considered secure.

Example global secret declaration:

```
PLUGIN_PARAMS=SMTP_PASSWORD=foo SLACK_TOKEN=bar
```

Example injecting global secrets into your `.drone.yml` file:

```
deploy:
  slack:
    channel: foo
    token: $$SLACK_TOKEN
```
