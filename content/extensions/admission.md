---
date: 2000-01-01T00:00:00+00:00
title: Admission Extension
title_in_sidebar: Admission
author: bradrydzewski
weight: 10
separator: true
toc: true
---

You can use an admission extension to apply custom access control logic to your Drone instance. For example, we [created](https://github.com/drone/drone-admit-members) a reference admission extension that grants system access based on organization and team membership.

# Configuration

You can register an admission extension by providing the following configuration parameters to the Drone server:

* `DRONE_ADMISSION_PLUGIN_ENDPOINT`
  : Provides the endpoint used to make http requests to an admission extension.

* `DRONE_ADMISSION_PLUGIN_SECRET`
  : Provides the token used to authenticate http requests to the admission extension. This token is shared between the server and extension.

# How it Works

The server makes an HTTP post to the admission extension during the login flow. The admission extension is expected to grant or deny the user access to the system, and can also (optionally) grant the user administrative access to the system.

# Request

The admission extension receives an HTTP request to authorize the user. The request body includes the User details in JSON format and the Event type. The Event differentiates between existing user logins and new user registrations.

Request Body definition:

```typescript  {linenos=table}
class Request {
  event: Event;
  user: User;
}
```

```typescript  {linenos=table}
class User {
  id: int64;
  login: string;
  email: string;
  admin: boolean;
}
```

```typescript  {linenos=table}
enum Event {
  login,
  register
}
```

Example Request Body:

```json  {linenos=table}
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "admin": false
}
```

# Response

The admission extension should respond with one of the following:

* `403` response code indicating the user is not authorized to login
* `204` response code indicating the user is granted access
* `200` response code, with a JSON encoded User object, to grant the user access and to grant or revoke administrative privileges.

Response definition:

```typescript  {linenos=table}
class User {
  login: string;
  admin: boolean;
}
```

Example response:

```json  {linenos=table}
{
    "login": "octocat",
    "admin": false
}
```


# Authorization

The http request is signed per the [http signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-10) draft specification use the shared secret. The receiver should use the signature to verify the authenticity and integrity of the webhook.

# Starter Project

If you are interested in creating an admission extension we recommend using our [starter project](https://github.com/drone/boilr-admission) as a base to jumpstart development.
