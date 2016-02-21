+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Webhook"
description = "Send build status notifications via Webhook"
user = "drone-plugins"
repo = "drone-webhook"
image = "plugins/drone-webhook"
tags = ["notify", "webhook", "rest", "json", "hook"]
categories = "notify"
draft = false
date = 2016-02-21T08:35:14Z
menu = ""
weight = 1

+++

Use the Webhook plugin to notify services via Webhook when a build completes.
You will need to supply Drone with outgoing Webhook URLs. You can override the
default configuration with the following parameters:

* `urls` - JSON payloads are sent to each URL
* `method` - HTTP request method. Defaults to `POST`
* `header` - HTTP request header map

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  webhook:
    urls:
      - https://your.webhook/...
      - https://your.other.webhook/...
    header:
      Authorization: pa55word
```

### Custom Body

In some cases you may want to submit a custom payload in the body of your hook.
For the use case we expose the following additional parameters:

* `template` - Handlebars template to create a custom payload body. See [docs](http://handlebarsjs.com/)
* `content_type` - HTTP request content type

Example configuration that generate a custom Yaml payload:

```yaml
notify:
  webhook:
    urls:
      - https://your.webhook/...
      - https://your.other.webhook/...
    content_type: application/yaml
    template: >
      repo: {{repo.full_name}}
      build: {{build.number}}
      commit: {{build.commit}}
```

### Basic Authentication

> It is important to note that with HTTP Basic Authentication the provided
> username and password are not encrypted.

In some cases your webhook may need to authenticate with another service. You
can set the basic `Authentication` header with a username and password. For
these use cases we expose the following additional parameters:

* `auth` - Sets the request's `Authorization` header to use HTTP Basic Authentication with the provided username and password below
  * `username` - The username as a string
  * `password` - The password as a string

Example configuration to include HTTP Basic Authentication:

```yaml
notify:
  webhook:
    method: POST
    auth:
      username: $$USERNAME
      password: $$PASSWORD
    urls:
      - https://tower.example.com/...
```

### Debugging Webhooks

> If you have private variables that are encrypted and hidden in `.drone.sec`,
> remember that the `debug` flag may print out those sensitive values. Please
> use `debug: true` wisely.

In some cases complicated webhooks may need debugging to ensure `urls`,
`template`, `auth` and more a properly configured. For these use cases we expose
the following `debug` parameter:

* `debug` - If `true` it will print out each URL request and response information

Example configuration to include the `debug` parameter:

```yaml
notify:
  webhook:
    debug: true
    method: POST
    auth:
      username: $$TOWER_USER
      password: $$TOWER_PASS
    urls:
      - http://tower.example.com/api/v1/job_templates/44/launch/
      - http://tower.example.com/api/v1/job_templates/45/launch/
      content_type: application/json
      template: '{"name": "project.deploy","extra_vars": "{\"env\": \"dev\",\"git_branch\": \"{{ build.branch }}\",\"hipchat_token\": \"$$HIPCHAT_TOKEN\"}"}'
```

Example of a debug print result:

```
[debug] Webhook 1
  URL: http://tower.example.com/api/v1/job_templates/44/launch/
  METHOD: POST
  HEADERS: map[Content-Type:[application/json] Authorization:[Basic EMfNB3fakB8EMfNB3fakB8==]]
  REQUEST BODY: {"name": "project.deploy","extra_vars": "{\"env\": \"dev\",\"git_branch\": \"develop\",\"hipchat_token\": \"h1pchatT0k3n\"}"}
  RESPONSE STATUS: 202 ACCEPTED
  RESPONSE BODY: {"job": 236}

[debug] Webhook 2
  URL: http://tower.example.com/api/v1/job_templates/45/launch/
  METHOD: POST
  HEADERS: map[Content-Type:[application/json] Authorization:[Basic EMfNB3fakB8EMfNB3fakB8==]]
  REQUEST BODY: {"name": "project.deploy","extra_vars": "{\"env\": \"dev\",\"git_branch\": \"develop\",\"hipchat_token\": \"h1pchatT0k3n\"}"}
  RESPONSE STATUS: 202 ACCEPTED
  RESPONSE BODY: {"job": 406}
```

