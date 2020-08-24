---
date: 2000-01-01T00:00:00+00:00
title: Enterprise
title_in_header: Enterprise FAQ
author: bradrydzewski
weight: 20000
hidden: true
inline_toc: true
---

# Is Drone Open Source?
There are two versions of Drone. The Open Source Edition is Free and Open Source software licensed under the Apache2 License. The Enterprise Edition is [source available](https://en.wikipedia.org/wiki/Source-available_software) and is licensed under the [Polyform Small Business](https://polyformproject.org/licenses/small-business/1.0.0/) License.

# Is Drone Enterprise Free for Individuals?
Yes

# Is Drone Enterprise Free for Students?
Yes

# Is Drone Enterprise Free for Startups?
The Enterprise Edition is free for organizations with under $1 million US dollars in annual gross revenue and less than $5 million in debt or equity funding.

# Is Drone Enterprise Free for Small Businesses?
The Enterprise Edition is free for organizations with under $1 million US dollars in annual gross revenue and less than $5 million in debt or equity funding.

# Is Drone Enterprise Free for Non-Profits?
The Enterprise Edition is free for organizations with under $1 million US dollars in annual gross revenue and less than $5 million in debt or equity funding.

# Is Drone Enterprise available for trial?
Yes. The Enterprise Edition includes a free trial for up to 5000 builds. You can download and use the Enterprise Edition by following the [installation instructions]({{< relref "server/overview.md" >}}) described in this documentation.

# How do I purchase an Enterprise License?
You can purchase an Enterprise License from our [website](https://drone.io/enterprise).

# What if my company cannot afford an Enterprise License?
The Enterprise Edition is free for organizations with under $1 million US dollars in annual gross revenue. There are affordable entry-level plans starting at $299 per month.

# Where do I download the Enterprise Edition?
You can download the Enterprise Edition from Dockerhub by following the [installation instructions]({{< relref "server/overview.md" >}}) described in this documentation.

# Where do I download the Open Source Edition?
There is no official distribution of the Open Source Edition. You can [build](https://github.com/drone/drone/blob/master/BUILDING_OSS) the Open Source Edition from source using the oss build tags. 

```
$ go build -tags "oss nolimit" github.com/drone/drone/cmd/drone-server
```

# How do I use the Enterprise Edition for Free?
You can build the Enterprise Edition from source without build limits if and only if your organization has under $1 million US dollars in annual gross revenue. You will otherwise need to [purchase](https://drone.io/enterprise) a commercial license.

```
$ go build -tags "nolimit" github.com/drone/drone/cmd/drone-server
```

_If an organization does not qualify for free use, removal of these limits violates the obligations and conditions to your license and is subject to penalty under US Federal and International copyright law. If you are unsure whether or not you qualify for free use you should consult your legal team._

# What is the difference between Open Source and Enterprise?

The Open Source Edition is fully capable of executing complex pipelines but is optimized to meet basic needs of individuals and small teams.

The Enterprise Edition:

* Supports distributed runners
* Supports kubernetes runners
* Supports organization secrets, vault secrets, etc
* Supports cron scheduling
* Supports scalable storage (postgres, mysql, s3)
* Supports autoscaling
* Supports [extensions]({{< relref "extensions/overview.md" >}})

The Open Source Edition:

* Lacks support for [extensions]({{< relref "extensions/overview.md" >}})
* Lacks support for distributed runners
* Lacks support for kubernetes runners
* Lacks support for organization secrets, vault secrets, etc
* Lacks support for cron scheduling
* Lacks support for autoscaling
* Limited to a single machine
* Limited to an embedded sqlite database
