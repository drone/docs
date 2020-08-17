---
date: 2000-01-01T00:00:00+00:00
title: Headers
author: bradrydzewski
weight: 21
description: |
  Configure server headers.
---

Drone provides a number of configuration options that can be used to secure http requests. Here is a list of available parameters:


```
DRONE_HTTP_ALLOWED_HOSTS
DRONE_HTTP_PROXY_HEADERS
DRONE_HTTP_SSL_REDIRECT
DRONE_HTTP_SSL_TEMPORARY_REDIRECT
DRONE_HTTP_SSL_HOST
DRONE_HTTP_SSL_PROXY_HEADERS
DRONE_HTTP_STS_SECONDS
DRONE_HTTP_STS_INCLUDE_SUBDOMAINS
DRONE_HTTP_STS_PRELOAD
DRONE_HTTP_STS_FORCE_HEADER
DRONE_HTTP_BROWSER_XSS_FILTER
DRONE_HTTP_FRAME_DENY
DRONE_HTTP_CONTENT_TYPE_NO_SNIFF
DRONE_HTTP_CONTENT_SECURITY_POLICY
DRONE_HTTP_REFERRER_POLICY
```

_This document is still a work-in-progress. You can read more about each of these parameters at [github.com/unrolled/secure](https://github.com/unrolled/secure)._

# Recommended Headers

We recommend setting these configuration parameters. _Please note these parameters should only be used when SSL is enabled._

```
DRONE_HTTP_SSL_REDIRECT=true
DRONE_HTTP_SSL_TEMPORARY_REDIRECT=true
DRONE_HTTP_SSL_HOST=drone.company.com
DRONE_HTTP_STS_SECONDS=315360000
```