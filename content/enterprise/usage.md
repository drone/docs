---
date: 2000-01-01T00:00:00+00:00
title: Enterprise
title_in_header: Usage FAQ
author: bradrydzewski
weight: 20000
hidden: true
inline_toc: true
---

# How does Drone count users?

The system counts all registered users. A registered user is user that has signed-in to Drone which in turn creates an entry for the user in the Drone database. The user count is a simple count of user entries in the database.

# How does Drone count repositories?

The system counts all enabled repositories. A repository is enabled by clicking the _Enable_ button in the user interface which in turn configures the repository to accept webhooks and trigger builds. The system does not count disabled repositories that have not been _Enabled_.

# How can I see the repository count?

* Find the repository count using the Prometheus [metrics]({{< relref "server/metrics" >}}) endpoint
  ```
  # HELP drone_repo_count Total number of registered repositories.
  # TYPE drone_repo_count gauge
  drone_repo_count 42
  ```

* Find the repository count by querying the database:
  ```
  SELECT count(*) FROM repos WHERE repo_active = true
  ```

# How can I see the user count?

* Find the user count using the Prometheus [metrics]({{< relref "server/metrics" >}}) endpoint
  ```
  # HELP drone_user_count Total number of active users.
  # TYPE drone_user_count gauge
  drone_user_count 42
  ```

* Find the user count using the command line tools
  ```
  drone user ls | wc -l
  ```

* Find the user count by querying the database
  ```
  SELECT count(*) FROM users
  ```

# How can I reduce the user count?

You can remove user accounts using the [command line tools]({{< relref "cli/user/drone-user-rm" >}}). _Please note that if you remove a user account, and that user account activated a repository, it will need to be de-activated and then re-activated by another user._

* Example command to list users, including their last login
  ```
  drone user ls --format="{{ .Login }} {{ .LastLogin | time }}
  ```

* Example command to remove a user
  ```
  drone user rm octocat
  ```

# Do inactive users count against usage limits?

The system does not differentiate between active an inactive users. This is because a user may look inactive (may not have recently authenticated) but may still be triggering builds and viewing results outside of the user interface (using command line tools, github status, etc).

We recognize our usage calculations are naive and could be improved to better reflect actual usage. We are considering adopting an active user formula that matches our competitors:

> An active user is anyone who triggers a build to run on Drone. The following count as an individual user:
> 1. Commits from registered and unregistered users that trigger builds, including pull request merges, and dependabot.
> 2. Running or re-running builds from the website or command line tools.
> 3. Running builds using cron
> 4. Machine users

The above formula may significantly increase user counts for some customers and reduce user counts for others. If we decide to change our usage formula we will provide 90 days notice to current customers.

# What happens when I exceed usage limits?

If you exceed the your license limit you will need to [reduce](#how-can-i-reduce-the-user-count) your user count or upgrade your subscription accordingly. We can usually process upgrade requests same-day. The cost to upgrade is prorated based on the number of days remaining in your current subscription period. Please contact our [sales team](mailto:sales@drone.io) for assistance.

# What if my user counts are rapidly growing?

Organizations with high growth potential should consider an unlimited site license. This eliminates the need for license accounting. The tradeoff is the annual license fee may be expensive if your current user counts are low relative to the overall size of your organization. Please contact our [sales team](mailto:sales@drone.io) to learn more.
