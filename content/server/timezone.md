---
date: 2000-01-01T00:00:00+00:00
title: Time Zone
author: tphoney
weight: 21
description: |
  Configure server timezone.
---

When running the Drone server image, the timezone can be set with an environment variable `TZ={Area/Location}` using a valid [TZ database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) EG:

{{< highlight bash "linenos=table,hl_lines=3" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --env=TZ=Europe/London
  --env=DRONE_GITHUB_CLIENT_ID=your-id \
  --env=DRONE_GITHUB_CLIENT_SECRET=super-duper-secret \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_SERVER_HOST=drone.company.com \
  --env=DRONE_SERVER_PROTO=https \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:2
{{< / highlight >}}

For more information on tzdata see [Wikipedia](https://en.wikipedia.org/wiki/Tz_database)
