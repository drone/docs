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
date = 2016-02-13T08:58:49Z
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

### Custom Templates

In some cases you may want to customize the look and feel of the email message
so you can use custom templates. For the use case we expose the following
additional parameters, all of the accept a custom handlebars template, directly
provided as a string or as a remote URL which gets fetched and parsed:

* `subject` - A handlebars template to create a custom subject. For more
  details take a look at the [docs](http://handlebarsjs.com/). You can see the
  default template [here](https://github.com/drone-plugins/drone-email/blob/master/template.go#L4)
* `template` - A handlebars template to create a custom template. For more
  details take a look at the [docs](http://handlebarsjs.com/). You can see the
  default template [here](https://github.com/drone-plugins/drone-email/blob/master/template.go#L8-L292)

Example configuration that generate a custom email:

```yaml
notify:
  email:
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    recipients:
      - octocat@github.com
    subject: >
      [{{ build.status }}]
      {{ repo.owner }}/{{ repo.name }}
      ({{ build.branch }} - {{ truncate build.commit 8 }})
    template: >
      https://git.io/vgvPz
```

### Skip SSL verify

In some cases you may want to skip SSL verification, even if we discourage that
as it leads to an unsecure environment. Please use this option only within your
intranet and/or with truested resources. For this use case we expose the
following additional parameter:

* `skip_verify` - Skip verification of SSL certificates

Example configuration that skips SSL verification:

```yaml
notify:
  email:
    from: noreply@github.com
    host: smtp.mailgun.org
    username: octocat
    password: 12345
    skip_verify: true
    recipients:
      - octocat@github.com
```

