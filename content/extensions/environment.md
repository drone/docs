---
date: 2000-01-01T00:00:00+00:00
title: Environment Extension
title_in_sidebar: Environment
author: bradrydzewski
weight: 10
toc: true
aliases:
- /configure-global-environment/
---


You can use environment extensions to provide your pipeline steps with custom environment variables. In addition, this extension can be used to inject default secrets or default plugin parameters into your pipelines _(secrets and plugin parameters are  passed to your pipeline as environment variables)_.

# Configuration

You can register an extension with your runners by providing the following configuration parameters:

* `DRONE_ENV_PLUGIN_ENDPOINT`
  : Provides the endpoint used to make http requests to the extension.

* `DRONE_ENV_PLUGIN_SECRET`
  : Provides the secret used to authenticate http requests to the extension. This token is shared between the server and extension.

# How it Works

The runner makes an HTTP POST request to retrieve a map of environment variables which are injected into each pipeline step.

# Request

The environment extension receives an HTTP request to return a map of environment variables. The JSON-encoded request body includes the repository and build information.

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

The environment extension should respond to the request with a 200 response code and a list of environment variables in JSON format.

Example response:

```json  {linenos=table}
[
  {
    "name": "keyb",
    "data": "valuea",
    "mask": false
  },
  {
    "name": "keyb",
    "data": "valueb",
    "mask": false
  }
]
```

# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Starter Project

If you are interested in creating an environment extension we recommend using our [starter project](https://github.com/drone/boilr-environ) as a base to jumpstart development.
