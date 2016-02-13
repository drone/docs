+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "HipChat"
description = "Send build status notifications via HipChat"
user = "drone-plugins"
repo = "drone-hipchat"
image = "plugins/drone-hipchat"
tags = ["chat", "messaging", "hipchat"]
categories = "notify"
draft = false
date = 2016-02-13T08:58:14Z
menu = ""
weight = 1

+++

Use this plugin for sending build status notifications via HipChat. You will
need to supply Drone with a HipChat authentication token. You can learn more
about authentication tokens [here](https://www.hipchat.com/docs/apiv2/auth). You
can override the default configuration with the following parameters:

* `url` - HipChat server URL, defaults to `https://api.hipchat.com`
* `auth_token` - HipChat API token
* `room_id_or_name` - ID or URL encoded name of the room
* `from` - A label to be shown, defaults to `drone`
* `notify` - Whether this message should trigger a user notification (change the
  tab color, play a sound, notify mobile phones, etc). Each recipient's
  notification preferences are taken into account, defaults to false

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  hipchat:
    auth_token: xxxxxxxxxxxxxxx
    room_id_or_name: 1234567
    notify: true
```

### Custom Messages

In some cases you may want to customize the body of the HipChat message you can
use custom templates. For the use case we expose the following additional
parameters:

* `template` - A handlebars template to create a custom payload body. For more
  details take a look at the [docs](http://handlebarsjs.com/).

Example configuration that generate a custom message:

```yaml
notify:
  hipchat:
    auth_token: xxxxxxxxxxxxxxx
    room_id_or_name: 1234567
    from: drone
    notify: true
    template: >
      <strong>{{ uppercasefirst build.status }}</strong> <a href=\"{{ system.link_url }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}\">{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.started_at build.finished_at }} </br> - {{ build.message }}
```

