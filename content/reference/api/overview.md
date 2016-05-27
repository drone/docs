+++
date = "2015-12-05T16:00:21-08:00"
title = "API overview"
weight = 1
toc = true

[menu.main]
	parent="api"
+++

# Overview

Drone provides a comprehensive remote API for interacting with the Drone server. This section of the documents provides instructions for authenticating and using the remote API.

# Libraries

Drone provides the following official libraries for integrating with the remote API:

Language | Repository
---------|----------------------------------------------------------------------
Go       | https://github.com/drone/drone-go
Node     | https://github.com/drone/drone-node

# Authentication

Drone uses tokens for authentication. You can retrieve your user token from your user profile screen in the Drone user interface. You can provide your token using the http `Authorization` header:

```
Authorization: Bearer AKIAIOSFODNN7EXAMPLE
```

Or using the `access_token` query parameters:

```
http://drone.mycompany.com/api/user?access_token=AKIAIOSFODNN7EXAMPLE
```
