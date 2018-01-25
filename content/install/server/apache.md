+++
date = "2017-04-15T14:39:04+02:00"
title = "Setup with Apache"
url = "setup-with-apache"

[menu.install]
  identifier = "setup-with-apache"
  parent = "install_server"
  weight = 5
+++

This guide provides a brief overview for installing Drone server behind the Apache2 webserver. This is an example configuration:

```nohighlight
ProxyPreserveHost On

RequestHeader set X-Forwarded-Proto "https"

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```

You must have the below Apache modules installed.

```nohighlight
a2enmod proxy
a2enmod proxy_http
```

You must configure Apache to set `X-Forwarded-Proto` when using https.

```diff
ProxyPreserveHost On

+RequestHeader set X-Forwarded-Proto "https"

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```
