+++
date = "2016-07-12T16:00:21-08:00"
draft = false
title = "Getting Help"
weight = 1

[menu.main]
	parent="support"
+++

# I am seeing unexpected or potentially buggy behavior

Start out by collecting whatever details you can. Log output, build output, your Drone version number, which source service you are pointed at (Github, Bitbucket, etc), which database Drone is using, etc. From there, try your luck in the [Gitter room](https://gitter.im/drone/drone).

If you don't receive a quick answer or would like to post something more in-depth for review, head on over to our [Discussion Site](https://discuss.drone.io/c/help). Make sure to share the details you have gathered (output, screenshots, versions, etc).

# I've found a bug in Drone!

*Before continuing, it's very important to make the distinction between Drone and Drone plugins*. If your issue regards an external service, you are probably dealing with a plugin. We ask that you don't file issues in our issue tracker for plugins unless it's one of the core ones that we maintain under the [drone-plugins](https://github.com/drone-plugins) org on Github.

Proceeding with the assumption that you have found a Drone bug (and not a plugin bug), stop by the [Gitter room](https://gitter.im/drone/drone) to make sure sure that you aren't running into a known issue. If it is determined that you have found a bug, your next step is our [issue tracker](https://github.com/drone/drone/issues). Please take a moment to read our issue tracker template and follow the steps.

# I've found a bug in a plugin

The vast majority of Drone plugins are community-maintained. You'll want to visit your plugin maintainer's issue tracker and describe what's going on.

# I'd like to request a feature

Stop by the [Gitter room](https://gitter.im/drone/drone) to get the discussion started. Describe what you are trying to do and what is currently standing in your way. We may be able to suggest an alternative path that is already possible.

If your request fits into the roadmap, it'll be worked into a future release. You may be asked to provide an initial pull request if the feature isn't high priority or on the current roadmap.
