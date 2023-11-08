---
date: 2000-01-01T00:00:00+00:00
title: Gitness vs Drone
author: bradrydzewski
weight: 10
separator: true
toc: false
disable_toc: true
---



# What is the future of Drone?
We recently [announced](https://www.linkedin.com/feed/update/urn:li:activity:7112103002798268416/) the next major version of Drone – rebranded as [Gitness](https://gitness.com). With this change we are adding native Source Control Management capabilities, which includes support for essential features like code hosting, pull requests, code reviews and more.

This is a major evolution of Drone, from Continuous Integration to a full fledged Developer Platform. Learn more at [gitness.com](https://gitness.com).

# Will I be forced to use Gitness for code hosting?
No. We plan to support Gitness pipelines for all major code hosting providers, including GitHub, GitLab, Bitbucket, Bitbucket Server, Gitea, Gogs, and Azure Devops.

# Will Drone pipelines work with Gitness?
Yes. We plan to make Gitness fully backward compatible with Drone.

# Will Drone extensions work with Gitness?
Yes. We are planning to continue support for extensions. We are also evaluating native support for popular extensions, such as the [paths changed](https://github.com/meltwater/drone-convert-pathschanged) extension.

# Will Gitness maintain API compatibility with Drone?
We will investigate creating a compatibility layer.

# Will upgrading to Gitness require a complex migration?
No. We plan to provide a seamless upgrade experience. It should ideally be as simple as upgrading to a new version of Drone.

# When should I upgrade from Drone to Gitness?
Gitness is currently available in Beta. We recommend upgrading to Gitness once a stable release is published. The requirement for publishing a stable release is full backward compatibility with Drone 2.x and a seamless upgrade experience.

# Will this impact my existing Drone Enterprise contract?
No. There will be no impact to the enterprise offering. You will still require an enterprise plan to use enterprise features.

# Where is the Drone source code?
Gitness development is taking place in the main branch. Drone 2.x development and maintenance is taking place in the [drone](https://github.com/harness/gitness/tree/drone) branch.

# Why is Gitness hosted on GitHub?
We use Gitness to develop Gitness. We also mirror our code on GitHub. We do this because we want to build a community which means we need to meet developers where they exist today. We are therefore happy to accept pull requests and feature requests via GitHub.
