---
date: 2000-01-01T00:00:00+00:00
title: Badges
author: bradrydzewski
weight: 1000
separator: true
aliases:
- /badges/
- /ccmenu/
description: |
  Configure pipeline status badges.
---

Drone has built-in support for rendering badges to display your repository build status. You can generate the badge URL from the repository _Settings_ screen in the user interface, or you can build the URL manually.

* The badge URL format:
  ```
  http(s)://[hostname]/api/badges/[namespace]/[name]/status.svg(?ref=[ref])
  ```

* Example URL using the default branch:
  ```
  https://drone.acme.com/api/badges/octocat/hello-world/status.svg
  ```

* Example URL using a custom branch:
  ```
  https://drone.acme.com/api/badges/octocat/hello-world/status.svg?ref=refs/heads/master
  ```

# Shields

[Shields](https://shields.io) is a third party service that provides customizable informational images as badges. Shields can be used as an alternative to Drone's built-in badges when more customization is required.

{{< link-to "https://shields.io/badges/drone" >}}
Create Custom Badges
{{< / link-to >}}
