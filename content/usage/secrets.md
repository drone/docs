+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Secrets"
weight = 4
toc = true


[menu.main]
	parent="usage"
+++

# Overview

Drone gives the ability to safely inject passwords, tokens, keys and sensitive information into your build as environment variables. Secrets are managed from the command line and stored in the central Drone server.

Example command adds the secret for the Heroku plugin:

```
drone secret add --image=heroku \
    octocat/hello-world HEROKU_TOKEN f1d2d2f924e986a
```

Example command removes the secret:

```
drone secret rm octocat/hello-world HEROKU_TOKEN
```

Example command signs the Yaml file:

```
drone sign octocat/hello-world
```

# Signature

Drone does not expose secrets to your build unless the Yaml file is signed and verified. You can sign the Yaml using the command line utility and committing the `.drone.yml.sig` file to your repository.

Any time you update your repository's `.drone.yml`, you must re-sign it. This prevents someone from opening a PR that prints your secrets.

Example command to sign your Yaml:

```
drone sign octocat/hello-world
```

# Skip Verification

Drone gives the option to skip signature verification. This disables import security checks and should only be used in trusted environments where your source code is not publicly accessible.

Example command adds secrets without signature verification:

```
drone secret add --image slack --skip-verify --insecure \
    octocat/hello-world SLACK_TOKEN f1d2d2f924e986a
```

# Limit Images

Drone requires an `--image` option to limit secrets to specific images or plugins. This limits the possible attack surface when using untrusted or public images in your build process.

Example command adds secrets to specific images:

```
drone secret add --image s3 --image ecr \
    octocat/hello-world AWS_CLIENT_SECRET f1d2d2f924e986a
```

# Limit Events

Drone gives the option to limit secrets to specific events, including `push`, `pull_request`, `tag` and `deployment`.

Example command adds secret to specific events:

```
drone secret add --image slack --event push --event tag \
    octocat/hello-world SLACK_TOKEN f1d2d2f924e986a
```

# Globbing

The secret implementation allows patterns and prefixes. It uses the glob package for matching. For example:

- `--image='*'` would match all images
- `--image=foo/*:latest` would match images with tag `latest` from the `foo` organization
- `--image=foo/*:*` would match images of any tag from the `foo` organization

# Multiline Secrets

You can add a multiline secret (or any secret really) by reading it from a local file.  For example:

```
drone secret add --image=heroku \
    octocat/hello-world HEROKU_TOKEN @/file/path/to/secret
```

# Pull Requests

Drone gives the option to expose secrets to pull requests, however, if your repository is public you should understand the potential risks. If your repository is public an attacker could submit a pull request that attempts to expose your secrets.

This attack vector is mitigated by signing your Yaml and only exposing secrets to plugins that do not execute arbitrary code. The official notification plugins meet this criteria and can be used safely with pull requests.

Example command adds secrets for push and pull request events:

```
drone secret add --image slack --event push --event pull_request \
    octocat/hello-world SLACK_TOKEN f1d2d2f924e986a
```
