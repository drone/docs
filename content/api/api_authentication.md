+++
date = "2000-01-01T00:00:00+00:04"
title = "API Authentication"
url = "api-authentication"

[menu.api]
  weight = 2
  identifier = "api-authentication"
  parent = "api_overview"
+++

Drone uses tokens for authentication. You can retrieve your user token from your user profile screen in the Drone user interface. You can provide your token using the http `Authorization` header:

```
Authorization: Bearer AKIAIOSFODNN7EXAMPLE
```

Or using the `access_token` query parameters:

```
http://drone.mycompany.com/api/user?access_token=AKIAIOSFODNN7EXAMPLE
```
