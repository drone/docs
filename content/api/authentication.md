+++
weight = -2
date = "2014-11-08T12:54:12-08:00"
title = "Authentication"

[menu.api]
parent = "API Documentation"
+++

Restricted resources require authentication. You can authorize a request by
providing your `access_token` as a query parameter. The token is located on
your account settings screen.

```bash
curl "https://drone.io/api/user?access_token=AUTH-TOKEN"
```

Requests that require authentication may return 401 Unauthorized or 404 Not Found
instead of 403 Forbidden, in some places. This is to prevent the accidental leakage
of private repositories to unauthorized users.

## From the Command Line

Many of the API resources can be accessed using the `drone` command line utility.
You can authenticate the command line utility by providing your `token` and `server`
(optional) as environment variables:

```bash
DRONE_TOKEN="AUTH-TOKEN"
DRONE_SERVER="https://drone.io"
drone whoami

batman
```

Or you can provide as command line arguments:

```bash
drone --token="AUTH-TOKEN" --server="https://drone.io" whoami

batman
```
