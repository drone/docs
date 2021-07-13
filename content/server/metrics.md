---
date: 2000-01-01T00:00:00+00:00
title: Metrics
author: bradrydzewski
weight: 21
toc: true
aliases:
- /configure-prometheus/
- /installation/metrics/
- /administration/server/metrics
description: |
  Configure server metrics.
---

Drone publishes and exposes metrics that can be consumed by Prometheus at the standard `/metrics` endpoint. Access to the metrics endpoint is restricted and requires an authorization token.

# Configuration

1. Create a machine user:

    ```
    $ drone user add prometheus --admin --machine
    ```

2. Drone will generate a 32-byte authorization token for this user:

    ```
    Successfully added user prometheus
    Generated account token fe8c402a51e6629aa1f43a4234afee81
    ```

3. Configure the prometheus scraper:

    ```
    global:
        scrape_interval: 60s

        scrape_configs:
          - job_name: 'drone'
            bearer_token: fe8c402a51e6629aa1f43a4234afee81

            static_configs:
              - targets: ['domain.com']
    ```

# Drone Metrics

Drone collects performance metrics exposed by the Go runtime, including memory, compute, garbage collection and more. These default metrics are augmented with the following Drone metrics:

- drone_build_count
  : total number of builds executed by the system.
- drone_user_count
  : total number of user accounts.
- drone_repo_count
  : total number of activated repositories.
- drone_pending_builds
  : total number of pending builds.
- drone_pending_jobs
  : total number of pending jobs. A single build can have one or many jobs, where a job represents a single pipeline in a multi-pipeline yaml.
- drone_running_builds
  : total number of running builds.
- drone_running_jobs
  : total number of running jobs. A single build can have one or many jobs, where a job represents a single pipeline in a multi-pipeline yaml.

## Custom Metrics

Drone collects a limited set of metrics by design. We are not currently accepting proposals to collect additional metrics in Drone core, however, we do provide a [starter project](https://github.com/drone/boilr-metrics) that you can use to create your own custom metrics provider to gather additional metrics.
