+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Gated Builds"
url = "release-0.6.0-gating"
+++

Drone adds __experimental__ support for gating your builds. You can enable gating for your repository in the user interface or with the following command:

```text
drone repo update <repo> --gated=true
```

The default gating logic is very basic. If the repository is public, and the hook is a pull request, and the user is not the repository admin, then the pull request is blocked pending approval. This default logic is intended to help mitigate bad actors sending harmful pull requests to open source projects.

# Customization

The default gating logic is not going to be useful for many installations. We fully expect most organizations will need to customize this logic. To enable customization we have made gating pluggable. This means you will be able to supply your own custom logic, tailored to your organization's specific needs.

Please note that we do not want to embed a complex approval engine inside of Drone. Nor do we want to embed your organizations custom approval logic in Drone. So please use plugins and please do not send pull requests that attempt to alter the default gating logic.

# What's Next

* Builtin logic should not require approval if pull request is sender has write access
* Ability to permanently add users to a repository whitelist
* Ability to permanently add users to a repository blacklist
