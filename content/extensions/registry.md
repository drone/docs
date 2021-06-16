---
date: 2000-01-01T00:00:00+00:00
title: Registry Extension
title_in_sidebar: Registry
author: bradrydzewski
weight: 10
toc: true
---

You can use registry extensions to provide your pipelines with docker login credentials. If your pipeline depends on private images, these credentials are used to authenticate with a remote registry and pull private images.

# Configuration

You can register an extension with your runners by providing the following configuration parameters:

* `DRONE_REGISTRY_PLUGIN_ENDPOINT`
  : Provides the endpoint used to make http requests to a registry extension.

* `DRONE_REGISTRY_PLUGIN_SECRET`
  : Provides the secret used to authenticate http requests to the extension. This token is shared between the server and extension.

# How it Works

The runner makes an HTTP POST request to retrieve a list of registry credentials. The runner matches the registry credentials with docker images in your yaml by comparing the registry address with the fully-qualified image url.

# Request

The registry extension receives an HTTP request to return a list of registry credentials. The JSON-encoded request body includes the repository and build information.

Request Body definition:

```typescript  {linenos=table}
class Request {
    repo: Repository;
    build: Build;
}
```

```typescript  {linenos=table}
class Repository {
    id: int64;
    uid: int64;
    user_id: int64;
    namespace: string;
    name: string;
    slug: string;
    scm: string;
    git_http_url: string;
    git_ssh_url: string;
    link: string;
    default_branch: string;
    private: boolean;
    visibility: string;
    active: boolean;
    config: string;
    trusted: boolean;
    protected: boolean;
    ignore_forks: boolean;
    ignore_pulls: boolean;
    cancel_pulls: boolean;
    timeout: int64;
    counter: int64;
    synced: int64;
    created: int64;
    updated: int64;
    version: int64;
}
```

```typescript  {linenos=table}
class Build {
    id: int64;
    repo_id: int64;
    number: int64;
    parent: int64;
    status: string;
    error: string
    event: string;
    action: string;
    link: string;
    timestamp: int64;
    title: string;
    message: string;
    before: string;
    after: string;
    ref: string;
    source_repo: string;
    source: string;
    target: string;
    author_login: string;
    author_name: string;
    author_email: string;
    author_avatar: string;
    sender: string;
    params: [string][string];
    cron: string;
    deploy_to: string;
    deploy_id: int64;
    started: int64;
    finished: int64;
    created: int64;
    updated: int64;
    version: int64;
}
```

# Response

The registry extension should respond to the request with a 200 response code and a list of registry credentials in JSON format.

Registry definition:

```typescript  {linenos=table}
class Registry {
  address: string
  username: string
  password: string
}
```

Example response:

```
[
  {
    "address": "docker.io",
    "username": "octocat",
    "password": "correct-horse-battery-staple"
  }
]
```

# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Starter Project

If you are interested in creating a registry extension we recommend using our [starter project](https://github.com/drone/boilr-registry) as a base to jumpstart development.
