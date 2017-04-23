+++
date = "2000-01-01T00:00:00+00:04"
title = "CLI Authentication"
url = "cli-authentication"

[menu.cli]
  weight = 2
  identifier = "cli-authentication"
  parent = "cli_overview"
+++

You will need to provide the Drone command line tools with your server address and personal access token. You can retrieve a Drone personal access token from your user profile screen. Click the show token button.

![user token](/images/drone_user_token.png)


You can provide the server address and token using environment variables:

```nohighlight
export DRONE_SERVER=http://drone.mycompany.com
export DRONE_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

You can confirm the above configuration by running the below command:

```nohighlight
$ drone info
User: octocat
Email: octocat@github.com
```
