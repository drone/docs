+++
weight = 3
date = "2014-11-08T12:54:06-08:00"
title = "Gitlab"

[menu.main]
parent = "Configure"
+++

You may configure Drone to integrate with GitLab (version 7.7 or higher). This can be configured in the `/etc/drone/drone.toml` configuration file:

```toml
[gitlab]
url = "https://gitlab.com"
client = "c0aaff74c060ff4a950d"
secret = "1ac1eae5ff1b490892f5546f837f306265032412"
skip_verify=false
open=false
```

Or a custom installation:

```toml
[gitlab]
url = "http://gitlab.drone.io"
client = "c0aaff74c060ff4a950d"
secret = "1ac1eae5ff1b490892f5546f837f306265032412"
skip_verify=false
open=false
```

## Environment Variables

You may also configure Gitlab using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_GITLAB_URL="https://gitlab.com"
DRONE_GITLAB_CLIENT="c0aaff74c060ff4a950d"
DRONE_GITLAB_SECRET="1ac1eae5ff1b490892f5546f837f306265032412"
```

## User Registration

User registration is closed by default and new accounts must be provisioned in the user interface. You may allow users to self-register with the following configuration flag:

```toml
[gitlab]
open = true
```

Please note this has security implications. This setting should only be enabled if you are running Drone behind a firewall.

## Self-Signed Certs

If your Gitlab installation uses a self-signed certificate you may need to instruct Drone to skip TLS verification. This is not recommended, but if you have no other choice you can include the following:

```toml
[gitlab]
skip_verify=true
```

---

## Generate Client and Secret

You must register your application with GitLab in order to generate a Client and Secret. 

1. If you have admin rigths on your GitLab instance, you can create application from admin panel ( *Admin Panel* > *Applications* > *New Application* )
	- Set Name field as *"Drone"*
	- Set Redirect URI field as `https://ci.yourhost.com/api/auth/gitlab.com`
2. If you not have admin rigths, you can create application from user profile ( *User profile* > *Edit Profile settings* > *Applications* > *New Application* )
	- Set Name field as *"Drone"*
	- Set Redirect URI field as `https://ci.yourhost.com/api/auth/gitlab.com`

Please use `/api/auth/gitlab.com` as the Authorization callback URL path.
