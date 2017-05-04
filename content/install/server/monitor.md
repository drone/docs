+++
date = "2017-04-15T14:39:04+02:00"
title = "Configure Prometheus"
url = "configure-prometheus"

[menu.install]
  identifier = "configure-prometheus"
  parent = "install_server"
  weight = 7
+++

{{% alert warn %}}
This feature is only available in the [Enterprise Edition](https://drone.io/enterprise/)
{{% /alert %}}


Drone is compatible with Prometheus and exposes a `/metrics` endpoint. Please note that access to the metrics endpoint is restricted and requires an authorization token with administrative privileges.

```nohighlight
global:
  scrape_interval: 60s

scrape_configs:
  - job_name: 'drone'
    bearer_token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

    static_configs:
       - targets: ['drone.domain.com']
```
