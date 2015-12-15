+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Email"
description = "Send build status notifications via Email"
user = "drone-plugins"
repo = "drone-email"
image = "plugins/drone-email"
tags = ["email"]
categories = "notify"
draft = false
date = 2015-12-15T06:51:17Z
menu = ""
weight = 1

+++

Use this plugin for sending build status notifications via Email. You can
override the default configuration with the following parameters:

* `from` - Send notifications from this address
* `host` - SMTP server host
* `port` - SMTP server port, defaults to `587`
* `username` - SMTP username
* `password` - SMTP password
* `recipients` - List of recipients, defaults to commit email

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  email:
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    recipients:
      - octocat@github.com
```
