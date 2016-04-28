+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Github"
weight = 20
menu = "installation"
toc = true
+++

# Overview

Drone comes with built-in support for GitHub and GitHub Enterprise. To enable GitHub you should configure the driver using the following required environment variables:

```
DRONE_GITHUB=true
DRONE_GITHUB_CLIENT=...
DRONE_GITHUB_SECRET=...
```

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration value that should work for the majority of installations.

NAME                        | DESC
----------------------------|--------------------------------------------------------
`DRONE_GITHUB_URL`          | github server address, defaults to https://github.com
`DRONE_GITHUB_CLIENT`       | github oauth client id
`DRONE_GITHUB_SECRET`       | github oauth client secret
`DRONE_GITHUB_SCOPE`        | github oauth scopes
`DRONE_GITHUB_PRIVATE_MODE` | github enterprise is running in private mode
`DRONE_GITHUB_SKIP_VERIFY`  | github certificate is self-signed

# Registration

You must register your application with GitHub in order to generate a Client and Secret. Navigate to your account settings and choose Applications from the menu, and click Register new application.

Please use `http://drone.mycompany.com/authorize` as the Authorization callback URL.

<!-- # Permissions

You may have issues if your organization limits third party organizations:

![third_party_restrictions](https://cloud.githubusercontent.com/assets/2988/5803370/e8024542-9fcb-11e4-8dc5-1810c2281e27.png)

You may need to grant access to individual organizations during authorization:

![third_party_restrictions](https://cloud.githubusercontent.com/assets/865/5805312/5701e842-9fd3-11e4-8f7b-a2bad994eb0a.gif) -->
