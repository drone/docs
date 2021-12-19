---
date: 2000-01-01T00:00:00+00:00
title: Enterprise
title_in_header: Pings
author: bradrydzewski
weight: 20000
hidden: true
toc: true
---

Drone periodically sends a ping to drone.io to help our product and customer teams. It sends only the high-level data below. It never sends code, repository names, or any other proprietary data.

Example payload:

```json {linenos=table}
{
    "installed_by": "jane@acme.com",
    "installed_at": "2021-01-01T00:00:00.000Z",
    "ip": "1.2.3.4",
    "hostname": "drone.acme.com",
    "provider": "github",
    "license": "trial",
    "version": "2.0.0",
    "user_count": 25,
    "repo_count": 50,
    "build_count": 2500,
}
```

Overview of payload attributes:

* `installed_by`
  : email address of the initial site installer (or if deleted, the first active site admin) collected from the user registration screen
* `installed_at`
  : installation date
* `hostname`
  : instance hostname
* `ip`
  : instance ip address
* `version`
  : version string
* `license`
  : license string (trial, enterprise)
* `provider`
  : which code host is being used (github, gitlab, bitbucket stash, gitea)
* `user_count`
  : aggregate count of the number of users
* `repo_count`
  : aggregate count of the number of repositories
* `build_count`
  : aggregate count of the number of builds

# Connections

Drone only connects to drone.io for the ping described above. There are no other automatic external connections to drone.io (or any other site on the internet).

# Privacy

Drone collects the email address of the initial site installer (or if deleted, the first active site admin), to know who to contact regarding sales, support, product updates, security updates, policy updates, and license compliance. The use of this email address is governed by our [Privacy Policy](https://harness.io/privacy/). If you would like us to stop processing your data, or if you have any other questions or requests concerning your data, please contact [privacy@harness.io](mailto:privacy@harness.io).  For more information on how we process your data, please visit our [privacy policy](https://harness.io/privacy/).


<!--

# Disable

Drone supports disabling this ping by setting the DRONE_DATADOG_ENABLED environment variable to false.
-->




