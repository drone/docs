---
date: 2000-01-01T00:00:00+00:00
title: Migration
author: tphoney
weight: 1
---

<div class="alert">
The VM runner is in the Release Candidate phase.
</div>

## Migrating from RC4 to latest

### .env file

- Removed all default runner settings for specific runners eg DRONE_SETTINGS_AWS_ACCESS_KEY_SECRET, DRONE_SETTINGS_AWS_ACCESS_KEY_ID , DRONE_SETTINGS_AWS_ACCESS_KEY_SECRET, DRONE_SETTINGS_AWS_REGION.
- DRONE_SETTINGS_LITE_ENGINE_PATH becomes DRONE_LITE_ENGINE_PATH and has a default.
- DRONE_SETTINGS_REUSE_POOL becomes DRONE_REUSE_POOL.

### Pool file

- .drone_pool.yml  was renamed from to pool.yml
- Syntax has changed significantly [example](https://github.com/drone-runners/drone-runner-aws/blob/master/pool_example.yml)
- Security / account info is done per instance, there is no global settings for credentials or region.

### Build file

- `type: aws` is now `type: vm`, this is to accommodate the new cloud / VM providers.
