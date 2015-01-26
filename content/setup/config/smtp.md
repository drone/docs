+++
weight = 7
date = "2014-11-08T12:54:12-08:00"
title = "SMTP"

[menu.main]
parent = "Configure"
+++

You may configure Drone to send Email notifications. An SMTP server can be configured in the `/etc/drone/drone.toml` configuration file:

```toml
[smtp]
host = "smtp.foo.com"
port = "587"
from = "bar@foo.com"
user = "foo"
pass = "password"
```

## Environment Variables

You may also configure your SMTP server using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_SMTP_HOST="smtp.foo.com"
DRONE_SMTP_PORT="587"
DRONE_SMTP_FROM="bar@foo.com"
DRONE_SMTP_USER="foo"
DRONE_SMTP_PASS="password"
```

---

## MailGun

If you don't already have an SMTP server we recommend using [Mailgun](http://www.mailgun.com/). It is extremely simple to setup and free to use. Also, if you don't have a Domain Name they provide free sandbox accounts that require no DNS configuration.

![GitHub Setup](/img/mailgun_setup.png)

Example Mailgun Configuration:

```
[smtp]
host = "smtp.mailgun.org"
port = "587"
from = "builds@sandboxda39a3ee5e6b4b0d3255.mailgun.org"
user = "postmaster@sandboxda39a3ee5e6b4b0d3255.mailgun.org"
pass = "f14d48f99ceea309c4aff4f86d2d7da4"
```