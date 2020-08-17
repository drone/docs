---
date: 2000-01-01T00:00:00+00:00
title: Configuration Extension
title_in_sidebar: Configuration
author: bradrydzewski
weight: 10
toc: true
aliases:
- /extend/config
- /extend/config/jsonnet
---

You can use a configuration extension to override the process of fetching the configuration file (e.g. .drone.yml). This can be used to return default or global configurations for projects where none exists. _For other use cases you should consider using a [conversion]({{< relref "conversion.md" >}}) extension_.

# Configuration

You can register a validation extension by providing the following configuration parameters to the Drone server:

* `DRONE_YAML_ENDPOINT`
  : Provides the endpoint used to make http requests to an extension.

* `DRONE_YAML_SECRET`
  : Provides the token used to authenticate http requests to the extension. This token is shared between the server and extension.

# How it Works

The server makes an HTTP post to the configuration extension to fetch the configuration file. If the configuration extension returns a 204 status code the system will fallback to the configuration file in the repository.

# Request

The configuration extension receives an HTTP request to return a configuration file. The request body includes the Repository and Build details in JSON format.

Request Body definition:

```typescript  {linenos=table}
class Request {
  repo: Repo;
  build: Build
}
```

Repository definition:

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

Build definition:

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

The configuration extension should respond to the request with a 200 response code and the raw configuration file. If the configuration extension returns a 204 status code the system will fallback to the configuration file in the repository.

Response definition:

```typescript  {linenos=table}
class Config {
    data: string;
}
```

# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Starter Project

If you are interested in creating a configuration extension we recommend using our [starter project](https://github.com/drone/boilr-config) as a base to jumpstart development.
