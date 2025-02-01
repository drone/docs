---
date: 2025-01-31T00:00:00+00:00
title: Harness Open Source vs Drone
author: bradrydzewski
weight: 10
separator: true
toc: false
disable_toc: true
---



# What is the future of Drone?
We recently [announced](https://www.harness.io/blog/introducing-harness-open-source) the next major version of Drone – rebranded as [Harness Open Source](https://developer.harness.io/docs/open-source/). With this change we are adding native Source Control Management capabilities, which includes support for essential features like code hosting, artifact registry, pre-configured development environments, and more.

This is a major evolution of Drone, from Continuous Integration to a full fledged Developer Platform. Learn more at [Harness Developer Hub](https://developer.harness.io/docs/open-source/).

# Can I still use Drone?
Yes. Harness Open Source is the next major version of Drone, however, Harness Open Source is still in Beta and under active development. You should continue using the latest stable version of Drone for production installations until a stable release of Harness Open Source is available.

# Will I be forced to use Harness Open Source for code hosting?
No. We plan to support Harness Open Source pipelines for all major code hosting providers, including GitHub, GitLab, Bitbucket, Bitbucket Server, Gitea, Gogs, and Azure Devops.

# Will Drone pipelines work with Harness Open Source?
Yes. We plan to make Harness Open Source fully backward compatible with Drone.

# Will Drone plugins work with Harness Open Source?
Yes.

# Will Drone extensions work with Harness Open Source?
Yes. We are planning to continue support for extensions. We are also evaluating native support for popular extensions, such as the [paths changed](https://github.com/meltwater/drone-convert-pathschanged) extension.

# Will Harness Open Source maintain API compatibility with Drone?
We will investigate creating a compatibility layer.

# Will upgrading to Harness Open Source require a complex migration?
No. We plan to provide a seamless upgrade experience. Because Harness Open Source is the next major version of Drone, upgrading to Harness Open Source should be no different than upgrading to a newer version of Drone.

# When should I upgrade from Drone to Harness Open Source?
Harness Open Source is currently available in Beta and should be used for preview purposes only. We will recommend upgrading to Harness Open Source once a stable release is published. The requirement for publishing a stable release is full backward compatibility with Drone 2.x and a seamless upgrade experience.

# Will this impact my existing Drone Enterprise contract?
No. There will be no impact to the enterprise offering. You will still require an enterprise plan to use enterprise features.

# Where is the Drone source code?
Harness Open Source development is taking place in the main branch. Drone 2.x development and maintenance is taking place in the [drone](https://github.com/harness/harness/tree/drone) branch.

# Why is Harness Open Source hosted on GitHub?
We use Harness Open Source to develop Harness Open Source. We also mirror our code on GitHub. We do this because we want to build a community which means we need to meet developers where they exist today. We are therefore happy to accept pull requests and feature requests via GitHub.
