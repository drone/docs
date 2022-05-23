---
date: 2000-01-01T00:00:00+00:00
title: DRONE_SETTINGS_AWS_AVAILABILITY_ZONE
author: tphoney
weight: 1
---

Optional value. The name of the specific AWS availability zone to use for the instance. If not specified, it will be randomly chosen from the region by AWS. This mimics how an EC2 instance is created on the AWS console specifying placement. IE: [here](https://docs.outscale.com/en/userguide/Launching-Instances-Using-AWS-CLI.html).

```bash
DRONE_SETTINGS_AWS_AVAILABILITY_ZONE=eu-west-2a
```
