---
date: 2000-01-01T00:00:00+00:00
title: Cron
author: bradrydzewski
weight: 170
toc: true
aliases:
- /configure/cron/
- /questions/how-to-schedule-builds/
description: |
  Setup cron jobs to execute pipelines at scheduled intervals.
---

You can use cron jobs to execute pipelines on time-based schedules. You can create and manage cron jobs from the repository _Settings_ screen, or using the command line tools.

* Example creates a cron job named _hourly_ using a cron expression:
  ```
  $ drone cron add "octocat/hello-world" "hourly" "0 0 * * * *"
  ```

* Example creates a cron job named _deploy_ using a pre-defined schedule:
  ```
  $ drone cron add "octocat/hello-world" "deploy" @daily
  ```

<div class="alert alert-no-cloud">
Please note that Cron scheduling is disabled on Drone Cloud. This feature is only available when self-hosting.
</div>

# Expressions

The cron expression represents a set of times, using 6 space-separated fields, including seconds. 

<div class="alert alert-warn">
Most examples on the internet represent the times using 5 space-separated fields and exclude seconds. Excluding seconds will lead to unexpected behavior.
</div>

Field name   | Mandatory? | Allowed values  | Allowed characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

Use one of several pre-defined schedules in place of a cron expression.

Entry    | Expression  | Description
---------|-------------|------------
@yearly  | 0 0 0 1 1 * | Run once a year, midnight, Jan. 1st
@monthly | 0 0 0 1 * * | Run once a month, midnight, first of month
@weekly  | 0 0 0 * * 0 | Run once a week, midnight between Sat/Sun
@daily   | 0 0 0 * * * | Run once a day, midnight
@hourly  | 0 0 * * * * | Run once an hour, beginning of hour

## Timezones

The current implementation calculates the execution time based on UTC as opposed to local time.

# Scheduling

The cron scheduler is approximate and executes jobs in hourly batches. This means cron jobs can be processed up to one hour after the scheduled execution time, and jobs scheduled with high frequency can only execute once per hour. This prevents users from creating high frequency jobs and overloading the system, but as a result, means cron scheduling has less accurate timing.

You can reduce the one hour window and increase scheduler accuracy by changing the server's cron interval setting.

{{< link "/server/reference/drone-cron-interval" "Customize the Cron Interval" >}}

# Pipelines

The cron scheduler executes all matching pipelines and steps defined in your yaml configuration file. See the pipeline configuration documentation to limit cron execution of pipelines and pipeline steps:

{{< link "/pipeline/docker/syntax/trigger.md#by-cron" "Limit Pipeline Execution" >}}
{{< link "/pipeline/docker/syntax/conditions.md#by-cron" "Limit Pipeline Step Execution" >}}
