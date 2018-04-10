+++
date = "2017-04-15T14:39:04+02:00"
title = "Setup with Ngrok"
url = "setup-with-ngrok"

[menu.install]
  identifier = "setup-with-ngrok"
  parent = "install_server"
  weight = 6
+++

After installing [ngrok](https://ngrok.com/), open a new console and run:

```
ngrok http 80
```

Set `DRONE_HOST` (for example in `docker-compose.yml`) to the ngrok url (usually xxx.ngrok.io) and start the server.
