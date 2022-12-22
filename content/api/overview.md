---
date: 2000-01-01T00:00:00+00:05
title: Overview
weight: 1
aliases:
- /api-overview/
- /api-authentication/

related:
  items:
  - name: Build Endpoint
    path: builds/_index.md
  - name: Cron Endpoint
    path: cron/_index.md
  - name: Repository Endpoint
    path: repos/_index.md
  - name: Secrets Endpoint
    path: secrets/_index.md
  - name: User Endpoint
    path: user/_index.md
  - name: Users Endpoint
    path: users/_index.md

description: |
  Overview of the remote API.
---

Drone provides a comprehensive remote API for interacting with the Drone server. This section of the documents provides instructions for authenticating and using the remote API.

# Authorization

The remote API uses access tokens to authorize requests. You can retrieve an access token in the Drone user interface by navigating to your user profile.

Authorization to the API is performed using the HTTP Authorization header. Provide your token as the bearer token value.

* Example Header.

  ```
  Authorization: Bearer AKIAIOSFODNN7EXAMPLE
  ```

* Example Request.

  ```
  curl -X GET "http://localhost:8080/api/user" \
    -H "Authorization: Bearer AKIAIOSFODNN7EXAMPLE"
  ```


# Libraries

Drone provides the following official libraries for integrating with the remote API:

* [github.com/drone/drone-go](https://github.com/drone/drone-go)
* [github.com/drone/drone-js](https://github.com/drone/drone-js)
* [github.com/drone/drone-node](https://github.com/drone/drone-node)

Community libraries:

* [github.com/tinvaan/pydroneio](https://github.com/tinvaan/pydroneio)
* [github.com/amir1430/drone-dart](https://github.com/amir1430/drone-dart)

<!--
Language | Repository
-------- | ----------
Go       | https://github.com/drone/drone-go
-->
