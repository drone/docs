---
date: 2000-01-01T00:00:00+00:00
title: Authentication
author: bradrydzewski
toc: true
weight: 1
description: |
  Describes the authentication process for cloning code
---

If your repository is private or requires authentication to clone, Drone injects the credentials into your pipeline environment. Drone uses the oauth2 token associated with the repository owner as the clone credentials.

# Security Consideration

The clone credentials are injected into to your pipeline environment which means a malicious collaborator could submit a patch (to your pipeline configuration, unit tests, etc) that attempts to expose these credentials during pipeline execution. Drone provides options, described below, to mitigate these risks.

## Machine Accounts

It is best practice to create a [machine user](https://docs.github.com/en/developers/overview/managing-deploy-keys#machine-users) with the exact permissions and access required to clone your repository and its dependencies. Login to Drone using the machine user account and activate your repositories. The machine account credentials will be injected into your pipeline environment, eliminating risk of exposing credentials linked to a user account.

## Global Machine Accounts

_In the previous section we described creating a [machine user](https://docs.github.com/en/developers/overview/managing-deploy-keys#machine-users). The downside to this approach is you have to login as the machine user to activate repositories. This section describes how to use machine account credentials regardless of the account that activated the repository._


If you configure a global machine account, it forces Drone to always use the global machine account credentials to clone all private repositories, regardless of the account that activated the repository. Please use the following links to configure global credentials:

{{< link "/server/reference/drone-git-username" >}}
{{< link "/server/reference/drone-git-password" >}}
