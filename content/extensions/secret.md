---
date: 2000-01-01T00:00:00+00:00
title: Secret Extension
title_in_sidebar: Secrets
author: bradrydzewski
weight: 10
toc: true
aliases:
- /extend/secrets
- /setup-global-secrets-endpoint/
---

You can use secret extensions to provide your pipeline with secrets from a custom, third party source. For example, we [created](https://github.com/drone/drone-vault) a reference secret extension to source secrets from Vault.

Here are some reference extensions:

* [Vault Secrets Extension](https://github.com/drone/drone-vault)
* [Amazon Secrets Extension](https://github.com/drone/drone-amazon-secrets)
* [Kubernetes Secrets Extension](https://github.com/drone/drone-kubernetes-secrets)

# Configuration

You can register a secret extension with your runners by providing the following configuration parameters:

* `DRONE_SECRET_PLUGIN_ENDPOINT`
  : Provides the endpoint used to make http requests to a secret extension.

* `DRONE_SECRET_PLUGIN_SECRET`
  : Provides the secret used to authenticate http requests to the extension. This token is shared between the server and extension.

<div class="alert">
Secret extensions are registered with runners as opposed to the central server. This design is intentional. For example, runners in different regions or data centers may have access to different secrets.
</div>

# How it Works

You can define an [external secret]({{< relref "secret/external/_index.md" >}}) resource in your Yaml configuration file. When you define an external secret, the runner makes an HTTP POST request to the secret extension to retrieve the external secret.

Example Yaml with external secret:

```yaml  {linenos=table}
kind: secret
name: username
get:
  path: secrets/data/docker
  name: username

---
kind: secret
name: password
get:
  path: secrets/data/docker
  name: password
```

# Request

The secret extensions receives an HTTP request to return a named secret. The JSON-encoded request body includes the name of the secret being requested, as well as the repository and build information.

Request Body definition:

```typescript  {linenos=table}
class Request {
    name: string;
    path: string;
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

The secret extension should respond to the request with a 200 response code, and the secret encoded in JSON format. If access to the requested secret cannot be granted, the extension should return a 204 status code.

Response definition:

```typescript  {linenos=table}
class Secret {
  name: string
  data: string
}
```

Example response:

```
{ "password": "correct-horse-battery-staple" }
```


# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Starter Project

If you are interested in creating a secret extension we recommend using our [starter project](https://github.com/drone/boilr-secret) as a base to jumpstart development.
