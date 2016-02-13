+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Slack Blame"
description = "Send build status blame messages via Slack"
user = "drone-plugins"
repo = "drone-slack-blame"
image = "plugins/drone-slack"
tags = ["chat", "messaging"]
categories = "notify"
draft = false
date = 2016-02-13T08:59:04Z
menu = ""
weight = 1

+++

Use the Slack Blame plugin to send a message to a Slack channel or through
direct message when a build completes. You will need to supply Drone with an
access token to the Slack API. You can create a new access token here:
https://api.slack.com/web




## Overview

This plugin is responsible for sending build notifications to developers when
they break the build, either directly through slackbot messages, or to the
designated channel. This lets developers get feedback on their own feature
branches without bombarding the main development channel as well as notifying
the channel on important branches such as release or master.





The following parameters are used to configure the notification

* **token** - the access token
* **channel** - the channel to post to (if present messages will also be posted
  to the channel)
* **success** - the options for a successful build
  * **icon** - an emoji or image url for the bot
  * **username** - the username for a successful build (defaults to "drone")
  * **message** - the message for a successful build
  * **image_attachments** - a list of image attachments to append (only one will
    be randomly selected per build)
* **failure** - the options for a build failure
  * **icon** - an emoji or image url for the bot
  * **username** - the username for a failed build (defaults to "drone")
  * **message** - the message for a failed build
  * **image_attachments** - a list of image attachments to append (only one will
    be randomly selected per build)

The following is a simple Slack Blame configuration in your .drone.yml file:

```yaml
notify:
  slack_blame:
    token: "xxxx"
    channel: "#dev"
    success:
      username: "Happy Keanu (on behalf of Drone)"
      icon: ":happy_keanu:"
      message: "The build is fixed!"
      image_attachments:
        - "http://i.imgur.com/TP4PIxc.jpg"
    failure:
      username: "Sad Keanu (on behalf of Drone)"
      icon: ":sad_keanu:"
      message: "The build is broken!"
      image_attachments:
        - "http://cdn.meme.am/instances/51000361.jpg"
```

