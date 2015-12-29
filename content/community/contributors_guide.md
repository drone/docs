+++
date = "2015-12-28T16:00:21-08:00"
draft = false
title = "Contributor's Guide"
weight = 2
menu = "community"
toc = true
+++

If you'd like to contribute to Drone, there is much to be done, regardless of
 skillset. Please read onward and jump in if any of this looks interesting to
  you. If you'd like some guidance in getting started, please stop by our 
  [Development category](https://discuss.drone.io/c/development) on the Drone
   Discussion site. You can also introduce yourself in our [Gitter room]
   (https://gitter.im/drone/drone).

# Drone Core

The core Drone daemon is written in Go and developed in the open on GitHub.
Head over to our [drone/drone repository](https://github.com/drone/drone) 
there and take a look at the README for details on how to get started. Fork 
the project, make your changes, and send in a pull request.

# Drone Plugins

Since we push most of the build logic out into plugins, the vast majority of 
our work ends up being focused here. Everything from cloning the source 
before a build to sending a Slack notification afterwards is done with a 
plugin.
 
All of our official plugins reside in the [drone-plugins](https://github
.com/drone-plugins) GitHub organization. Feel free to look through what we 
have there (or on our [plugins list](../../plugins)) and fork/PR away.

If you have a plugin of your own that you'd like to contribute, toss a post 
up on our [Development category](https://discuss.drone.io/c/development) on 
the Drone Discussion site and let us know what you've got.

# Drone UI and Design

This is an area that we are always eager to get help with. If you'd like to 
contribute on the design or UI front, head over to the [drone/drone-ui]
(https://github.com/drone/drone-ui) project and take a look. 

# Drone Documentation

We place a very high value on having excellent documentation. No contribution
 is too small, and we'd love to have your input on what we do (and don't) 
 have. Head over to the [drone/docs](https://github.com/drone/docs) 
 repository and send pull requests our way. 

# Drone Demos, Tips, Tricks, and How-To's

For many people, learning is best done by example. Submissions in here are a 
great way for new contributors to get started. On the flipside, our 
experienced community members have a lot to offer in sharing what they've 
learned.

The [drone-demos](https://github.com/drone-demos) GitHub organization has a 
number of repositories containing example projects that build with Drone. We 
are *always* on the lookout for more of these. We'd suggest working on your 
example repo separately, then post to the [Development]
(https://discuss.drone.io/c/development) Discussion site category when you are 
ready. We can help you get it moved over to the drone-demos GitHub 
organization and/or added as an org member.
 
Another way to contribute example material is via the [Tips, Tricks, and 
How-To's](https://discuss.drone.io/c/how-tos) Discussion site category. Most 
of the stuff in drome-demos will end up with a topic here, but we also would 
like more bite-sized examples of how to do things. For example, "[Publishing 
Docker images to Quay.io](https://discuss.drone
.io/t/publishing-docker-images-to-quay-io/33/1)". These are short, concise 
examples that are still valuable to have on file.
