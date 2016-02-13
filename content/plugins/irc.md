+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "IRC"
description = "Send build status notifications via IRC"
user = "drone-plugins"
repo = "drone-irc"
image = "plugins/drone-irc"
tags = ["chat", "messaging", "irc"]
categories = "notify"
draft = false
date = 2016-02-13T08:58:18Z
menu = ""
weight = 1

+++

Use this plugin for sending build status notifications via IRC. You can override
the default configuration with the following parameters:

* `prefix` - Prefix for the sent notifications
* `nick` - Nickname used by the bot
* `channel` - Messages sent are posted here
* `recipient` - Alternatively you can send it to a specific user
* `server` - Connection information for the server
  * `host` - IRC server host to connect to
  * `port` - IRC server port, defaults to 6667
  * `password` - Password for IRC server, optional
  * `tls` - Enable TLS, defaults to false

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  irc:
    prefix: build
    nick: drone
    channel: my-channel
    server:
      host: chat.freenode.net
      port: 6667
```

