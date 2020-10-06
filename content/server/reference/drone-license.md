---
date: 2000-01-01T00:00:00+00:00
title: DRONE_LICENSE
author: bradrydzewski
aliases:
- /reference/server/drone-license/
- /installation/reference/drone-license/
---

Optional string value provides the filepath of the Drone Enterprise license key. This is used to unlock the Drone Enterprise edition.

```
DRONE_LICENSE=/etc/drone.key
```

If you are running the Drone server in a Docker container you will need to mount the license key as a volume:

```
$ docker run \
  --volume=/path/on/host/drone.key:/etc/drone.key
```

If you are running the Drone server using docker-compose or Kubernetes or you have configured Drone using Yaml, you can provide the server with the license key as an environment variable:

```
DRONE_LICENSE: |
  -----BEGIN LICENSE KEY-----
  Thjh7sTA1VDE4OjM2tpmQQZCyRd43M1ODI1OVoiLCJkYXQiSukU/Y
  -----END LICENSE KEY-----
```

# Common Problems

The most common root cause for license errors is failing to mount the license key, or providing an old or expired license key. The debug server logs provide useful information that can help troubleshoot license issues.

Error indicates license file was not properly mounted:

```json {linenos=table,linenostart=1,hl_lines=[2]}
{
  "error": "open /tmp/LICENSE.txt: no such file or directory",
  "level": "fatal",
  "msg": "main: invalid or expired license",
  "time": "2020-10-01T12:04:39-04:00"
}
```

Error indicates license environment variable was empty and was not properly configured, resulting in a trial license being issued. This will result in _License Expired_ errors if you have exceeded the trial build limit.

```json {linenos=table,linenostart=1,hl_lines=[4]}
{
  "build.limit": 5000,
  "expires": "0001-01-01T00:00:00Z",
  "kind": "trial",
  "level": "debug",
  "msg": "main: license loaded",
  "repo.limit": 0,
  "time": "2020-10-01T12:05:32-04:00",
  "user.limit": 0
}
```

Example log entry indicating the license was successfully loaded:

```json {linenos=table,linenostart=1,hl_lines=[4, 6, 8]}
{
  "build.limit": 0,
  "expires": "2020-11-24T23:26:34.67261883Z",
  "kind": "standard",
  "level": "debug",
  "msg": "main: license loaded",
  "time": "2019-12-02T10:38:21-08:00",
  "user.limit": 50
}
```

Please double-check the user limits to ensure they align with your expectations. If you continue to experience _License Exceeded_ errors you should [follow our guide]({{< ref "/enterprise/usage.md" >}}) to make sure your usage does not exceed the license limits.