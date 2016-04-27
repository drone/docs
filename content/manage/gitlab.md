+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Gitlab"
weight = 22
menu = "installation"
toc = true
+++

# Overview

Drone comes with built-in support for GitLab version 8.2 and higher. To enable Gitlab you should configure the driver using the following environment variables:

```bash
DRONE_GITLAB_URL=http://gitlab.mycompany.com
DRONE_GITLAB_CLIENT=...
DRONE_GITLAB_SECRET=...
```

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration value that should work for the majority of installations.

NAME                          | DESC
------------------------------|--------------------------------------------------------
`DRONE_GITLAB_URL`            | gitlab server address
`DRONE_GITLAB_CLIENT`
`DRONE_GITLAB_SECRET`
`DRONE_GITLAB_HIDE_ARCHIVES`  | gitlab repositories are excluded if archived
`DRONE_GITLAB_CLONE_AUTH`     | gitlab repositories clone with oauth token (false)
`DRONE_GITLAB_SKIP_VERIFY`    | gitlab certificate is self-signed

# Registration

You must register your application with GitLab in order to generate a Client and Secret. Navigate to your account settings and choose Applications from the menu, and click New Application.

Please use `http://drone.mycompany.com/authorize` as the Authorization callback URL.

# Known Issues

There are known issues using GitLab with an Apache reverse proxy. If you are using Apache without `AllowEncodedSlashes NoDecode` you should set the following environment variable:

```
DRONE_GITLAB_SEARCH=true
```
