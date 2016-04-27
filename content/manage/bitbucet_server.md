+++
date = "2016-04-20T16:00:21-08:00"
draft = false
title = "Bitbucket Server"
weight = 25
menu = "installation"
toc = true
break = true
+++

# Overview

Drone comes with experimental support for Bitbucket Server (Atlassian Stash) and should be considered highly unstable. Now that you've been warned, and if you aren't scared away, you can enable Bitbucket Server using the following required environment variables:

```bash
DRONE_STASH_URL=http://stash.mycompany.com
DRONE_STASH_GIT_USERNAME=...
DRONE_STASH_GIT_PASSWORD=...
DRONE_STASH_CONSUMER_KEY=...
DRONE_STASH_CONSUMER_RSA=/path/to/key.pem
```

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration value that should work for the majority of installations.

NAME                        | DESC
----------------------------|--------------------------------------------------------
`DRONE_STASH_URL`           | bitbucket server address
`DRONE_STASH_CONSUMER_KEY`  | oauth1 consumer key
`DRONE_STASH_CONSUMER_RSA`  | oauth1 private key file path
`DRONE_STASH_GIT_USERNAME`  | service account username
`DRONE_STASH_GIT_PASSWORD`  | service account password

# Private Key File

Please note that the private key file needs to be mounted into your Drone container at runtime as a volume:

```
docker run --volume=/etc/bitbucket/key.pem:/etc/bitbucket/key.pem
```

Provide drone with the path to the key file inside the container:

```
DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
```

# Service Account

Drone users `git+https` to clone repositories, however, Bitbucket Server does not currently support cloning repositories with oauth token. To work around this limitation, you must create a service account and provide the username and password to Drone. This service account will be used to authenticate and clone private repositories.

# Registration

You must register your application with Bitbucket Server in order to generate a consumer key and key file. Navigate to your account settings and choose Applications from the menu, and click Register new application.

Please use http://drone.mycompany.com/authorize as the Authorization callback URL.
