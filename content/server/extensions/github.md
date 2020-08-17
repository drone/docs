---
date: 2000-01-01T00:00:00+00:00
title: GitHub Teams
author: bradrydzewski
weight: 90
description: |
  Install the GitHub Teams admission extension.
---


Drone provides an official extension to limit system access based on GitHub organization and team membership. You can use this extension to configure access policies, such as:

* If user is organization member, grant access
* If user is organization admin, grant admin access
* If user is member of designated team, grant admin access (optional)
* Else deny access

<!-- If your yaml configuration files are growing increasing complex, you may want to give Starlark a shot. -->

# Configuration

* DRONE_GITHUB_ENDPOINT
  : Github API base URL. This is only required when integrating with GitHub Enterprise. The URL format should be `http(s)://[hostname]/api/v3`

* DRONE_GITHUB_TOKEN
  : GitHub API personal access token. This token must have adequate permissions to access organization and team endpoints.

* DRONE_GITHUB_ORG
  : Comma-separated lists of organizations. If defined, the user must be a member of at least one organization in the list.

* DRONE_GITHUB_TEAM
  : Comma-separated lists of teams. If defined, users that are members of this team are granted administrative access.

# Installation

1. Create a shared secret.

    ```
    $ openssl rand -hex 16
    bea26a2221fd8090ea38720fc445eca6
    ```

2. Download and run the extension.

    ```
    $ docker run -d \
    --publish=3000:3000 \
    --env=DRONE_DEBUG=true \
    --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
    --env=DRONE_GITHUB_TOKEN=3da541559918a808c2402bba5012f6c6 \
    --env=DRONE_GITHUB_ORG=acme \
    --env=DRONE_GITHUB_TEAM=admins \
    --restart=always \
    --name=admitter drone/drone-admit-members
    ```

3. Update your Drone server configuration to include the extension address and the shared secret.

    ```
    DRONE_ADMISSION_PLUGIN_ENDPOINT=http://1.2.3.4:3000
    DRONE_ADMISSION_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
    ```

# Verification

You can verify the extension is configured and is processing requests using the command line utility.

1. Provide the command line utility with the extension endpoint and secret.

    ```
    export DRONE_ADMISSION_ENDPOINT=http://localhost:3000
    export DRONE_ADMISSION_SECRET=bea26a2221fd8090ea38720fc445eca6
    ```

2. Use the command line utility to check if a user is admitted:

    ```
    $ drone plugins admit octocat
    ```

# Customization

This extension is considered a reference implementation of an admission controller, and has limited scope. You are encouraged to fork and customize this extension as needed. You can find the source code at [drone/drone-admit-members](https://github.com/drone/drone-admit-members).
