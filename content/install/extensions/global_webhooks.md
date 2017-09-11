+++
date = "2017-04-15T14:39:04+02:00"
title = "Global Webhooks"
url = "configure-system-webhooks"

[menu.install]
  identifier = "configure-system-webhooks"
  parent = "install_enterprise"
  weight = 8
+++

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise expansion pack provides support for global webhooks for internal system events. Notification are sent as HTTP POST requests to an endpoint of your choosing.

Example server configuration:

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_HOOK=http://localhost:4567/payload
```

Webhook requests are triggered for the following event types:

Name        | Description
------------|-------------
`user`      | any time a user is created or deleted
`repo`      | any time a repository is created or updated
`build`     | any time a build is created or updated

Webhook requests contain the following special headers:

Header          | Description
----------------|-------------
`X-Drone-Event` | name of the event that triggered this delivery

Example delivery:

```nohighlight
POST /payload HTTP/1.1

Host: localhost:4567
Content-Type: application/json
Content-Length: 100
X-Drone-Event: user

{
  "action": "created",
  "user": {
    "id": 1,
    "login": "johnsmith",
    "active": true
  }
}
```
