+++
weight = 4
date = "2014-11-08T12:53:55-08:00"
title = "Bitbucket"

[menu.main]
parent = "Configure"
+++

You may configure Drone to integrate with Bitbucket. This can be configured in the `/etc/drone/drone.toml` configuration file:

```toml
[bitbucket]
client = "c0aaff74c060ff4a950d"
secret = "1ac1eae5ff1b490892f5546f837f306265032412"
```

## Environment Variables

You may also configure Bitbucket using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_BITBUCKET_CLIENT="c0aaff74c060ff4a950d"
DRONE_BITBUCKET_SECRET="1ac1eae5ff1b490892f5546f837f306265032412"
```

---

## Generate Client and Secret

You must register your application with Bitbucket in order to generate a Client and Secret. Navigate to your account settings and choose OAuth from the menu, and click Add Consumer.

Please use `/api/auth/bitbucket.org` as the Authorization callback URL path:

![Bitbucket Setup](/img/bitbucket_setup.png)

Copy the generated Client (key) and Secret into your Drone configuration file:

![Bitbucket Setup](/img/bitbucket_setup_2.png)