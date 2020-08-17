---
date: 2000-01-01T00:00:00+00:00
title: Validation Extension
title_in_sidebar: Validation
author: bradrydzewski
weight: 10
toc: true
---

You can use an validation extension to enforce custom checks, verification and linting rules to .drone.yml configuration files.

# Configuration

You can register a validation extension by providing the following configuration parameters to the Drone server:

* `DRONE_VALIDATE_PLUGIN_ENDPOINT`
  : Provides the endpoint used to make http requests to an extension.

* `DRONE_VALIDATE_PLUGIN_SECRET`
  : Provides the token used to authenticate http requests to the extension. This token is shared between the server and extension.

# How it Works

The server makes an HTTP post to the validation extension before the yaml file is processed and before any pipelines are scheduled. The validation extension is expected to accept or reject the yaml file.

# Request

The validation extension receives an HTTP request to verify the yaml. The request body includes the Repository and Build details in JSON format, as well as the raw Yaml configuration file.

Request Body definition:

```typescript  {linenos=table}
class Request {
  config: Config;
  repo: Repo;
  build: Build
}
```

```typescript  {linenos=table}
class Config {
    data: string;
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

The validation extension should respond with one of the following:

* `200` response code indicating the yaml passed validation
* `400` response code indicating the yaml filed validation, with a JSON encoded error messaging indicating why validation failed.

```
{
    "message": "cannot use image from external registry"
}
```

# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Starter Project

If you are interested in creating an validation extension we recommend using our [starter project](https://github.com/drone/boilr-validate) as a base to jumpstart development.
