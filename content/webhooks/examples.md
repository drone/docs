---
date: 2000-01-01T00:00:00+00:00
title: Example Payloads
title_in_sidebar: Examples
author: bradrydzewski
weight: 3
toc: true
---

This article provides reference webhook payloads. These payload examples are grouped with their corresponding event and action.

# User
## Created

```json  {linenos=table}
{
  "action": "created",
  "user": {
    "id": 1,
    "login": "octocat",
    "email": "octocat@github.com",
    "active": true,
    "admin": false,
    "machine": false,
    "syncing": false,
    "synced": 1564096971,
    "created": 1564096971,
    "updated": 1564096971,
    "last_login": 1564096971
  }
}
```

## Deleted

```json  {linenos=table}
{
  "action": "deleted",
  "user": {
    "id": 1,
    "login": "octocat",
    "email": "octocat@github.com",
    "active": true,
    "admin": false,
    "machine": false,
    "syncing": false,
    "synced": 1564096971,
    "created": 1564096971,
    "updated": 1564096971,
    "last_login": 1564096971
  }
}
```

# Repository
## Activated

```json  {linenos=table}
{
  "action": "activated",
  "repo": {
    "id": 42,
    "uid": "16607898",
    "user_id": 2,
    "namespace": "octocat",
    "name": "octocat",
    "slug": "octocat/hello-world",
    "scm": "git",
    "git_http_url": "https://github.com/octocat/hello-world.git",
    "git_ssh_url": "git@github.com:octocat/hello-world.git",
    "link": "https://github.com/octocat/hello-world",
    "default_branch": "master",
    "private": false,
    "visibility": "public",
    "active": true,
    "config_path": ".drone.yml",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pull_requests": false,
    "timeout": 60,
    "counter": 185,
    "synced": 1542160374,
    "created": 1542160374,
    "updated": 1542160374,
    "version": 187,
    "permissions": {
    "read": true,
    "write": true,
    "admin": true
  }
}
```

## Deactivated

```json  {linenos=table}
{
  "action": "activated",
  "repo": {
    "id": 42,
    "uid": "16607898",
    "user_id": 2,
    "namespace": "octocat",
    "name": "octocat",
    "slug": "octocat/hello-world",
    "scm": "git",
    "git_http_url": "https://github.com/octocat/hello-world.git",
    "git_ssh_url": "git@github.com:octocat/hello-world.git",
    "link": "https://github.com/octocat/hello-world",
    "default_branch": "master",
    "private": false,
    "visibility": "public",
    "active": false,
    "config_path": ".drone.yml",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pull_requests": false,
    "timeout": 60,
    "counter": 185,
    "synced": 1542160374,
    "created": 1542160374,
    "updated": 1542160374,
    "version": 187,
    "permissions": {
    "read": true,
    "write": true,
    "admin": true
  }
}
```

# Build
## Created

```json  {linenos=table}
{
  "action": "created",
  "repo": {
    "id": 42,
    "uid": "16607898",
    "user_id": 2,
    "namespace": "octocat",
    "name": "octocat",
    "slug": "octocat/hello-world",
    "scm": "git",
    "git_http_url": "https://github.com/octocat/hello-world.git",
    "git_ssh_url": "git@github.com:octocat/hello-world.git",
    "link": "https://github.com/octocat/hello-world",
    "default_branch": "master",
    "private": false,
    "visibility": "public",
    "active": false,
    "config_path": ".drone.yml",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pull_requests": false,
    "timeout": 60,
    "counter": 185,
    "synced": 1542160374,
    "created": 1542160374,
    "updated": 1542160374,
    "version": 187
  },
  "build": {
    "id": 100207,
    "repo_id": 296163,
    "number": 42,
    "status": "pending",
    "event": "pull_request",
    "action": "sync",
    "link": "https://github.com/octoat/hello-world/compare/e3320539a4c0...9fc1ad6ebf12",
    "message": "updated README",
    "before": "e3320539a4c03ccfda992641646deb67d8bf98f3",
    "after": "9fc1ad6ebf12462f3f9773003e26b4c6f54a772e",
    "ref": "refs/heads/master",
    "source_repo": "spaceghost/hello-world",
    "source": "develop",
    "target": "master",
    "author_login": "octocat",
    "author_name": "The Octocat",
    "author_email": "octocat@github.com",
    "author_avatar": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
    "sender": "bradrydzewski",
    "started": 1564085874,
    "finished": 1564086343,
    "created": 1564085874,
    "updated": 1564085874,
    "version": 1
  }
}
```

## Updated

```json  {linenos=table}
{
  "action": "updated",
  "repo": {
    "id": 42,
    "uid": "16607898",
    "user_id": 2,
    "namespace": "octocat",
    "name": "octocat",
    "slug": "octocat/hello-world",
    "scm": "git",
    "git_http_url": "https://github.com/octocat/hello-world.git",
    "git_ssh_url": "git@github.com:octocat/hello-world.git",
    "link": "https://github.com/octocat/hello-world",
    "default_branch": "master",
    "private": false,
    "visibility": "public",
    "active": false,
    "config_path": ".drone.yml",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pull_requests": false,
    "timeout": 60,
    "counter": 185,
    "synced": 1542160374,
    "created": 1542160374,
    "updated": 1542160374,
    "version": 187
  },
  "build": {
    "id": 100207,
    "repo_id": 296163,
    "number": 42,
    "status": "started",
    "event": "pull_request",
    "action": "sync",
    "link": "https://github.com/octoat/hello-world/compare/e3320539a4c0...9fc1ad6ebf12",
    "message": "updated README",
    "before": "e3320539a4c03ccfda992641646deb67d8bf98f3",
    "after": "9fc1ad6ebf12462f3f9773003e26b4c6f54a772e",
    "ref": "refs/heads/master",
    "source_repo": "spaceghost/hello-world",
    "source": "develop",
    "target": "master",
    "author_login": "octocat",
    "author_name": "The Octocat",
    "author_email": "octocat@github.com",
    "author_avatar": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
    "sender": "bradrydzewski",
    "started": 1564085874,
    "finished": 1564086343,
    "created": 1564085874,
    "updated": 1564085874,
    "version": 3,
    "stages": [
      {
        "id": 199937,
        "repo_id": 296163,
        "build_id": 100207,
        "number": 1,
        "name": "default",
        "kind": "pipeline",
        "type": "docker",
        "status": "started",
        "errignore": false,
        "exit_code": 0,
        "machine": "15e89c0f84f1",
        "os": "linux",
        "arch": "amd64",
        "started": 1564085874,
        "stopped": 1564086343,
        "created": 1564085874,
        "updated": 1564086343,
        "version": 4,
        "on_success": true,
        "on_failure": false,
        "steps": [
          {
            "id": 575525,
            "step_id": 199937,
            "number": 1,
            "name": "clone",
            "status": "success",
            "exit_code": 0,
            "started": 1564085876,
            "stopped": 1564085901,
            "version": 4
          },
          {
            "id": 575526,
            "step_id": 199937,
            "number": 2,
            "name": "test",
            "status": "started",
            "exit_code": 0,
            "started": 1564085901,
            "stopped": 1564085994,
            "version": 4
          }
        ]
      }
    ]
  }
}
```

## Completed

```json  {linenos=table}
{
  "action": "completed",
  "repo": {
    "id": 42,
    "uid": "16607898",
    "user_id": 2,
    "namespace": "octocat",
    "name": "octocat",
    "slug": "octocat/hello-world",
    "scm": "git",
    "git_http_url": "https://github.com/octocat/hello-world.git",
    "git_ssh_url": "git@github.com:octocat/hello-world.git",
    "link": "https://github.com/octocat/hello-world",
    "default_branch": "master",
    "private": false,
    "visibility": "public",
    "active": false,
    "config_path": ".drone.yml",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pull_requests": false,
    "timeout": 60,
    "counter": 185,
    "synced": 1542160374,
    "created": 1542160374,
    "updated": 1542160374,
    "version": 187
  },
  "build": {
    "id": 100207,
    "repo_id": 296163,
    "number": 42,
    "status": "success",
    "event": "pull_request",
    "action": "sync",
    "link": "https://github.com/octoat/hello-world/compare/e3320539a4c0...9fc1ad6ebf12",
    "message": "updated README",
    "before": "e3320539a4c03ccfda992641646deb67d8bf98f3",
    "after": "9fc1ad6ebf12462f3f9773003e26b4c6f54a772e",
    "ref": "refs/heads/master",
    "source_repo": "spaceghost/hello-world",
    "source": "develop",
    "target": "master",
    "author_login": "octocat",
    "author_name": "The Octocat",
    "author_email": "octocat@github.com",
    "author_avatar": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
    "sender": "bradrydzewski",
    "started": 1564085874,
    "finished": 1564086343,
    "created": 1564085874,
    "updated": 1564085874,
    "version": 3,
    "stages": [
      {
        "id": 199937,
        "repo_id": 296163,
        "build_id": 100207,
        "number": 1,
        "name": "default",
        "kind": "pipeline",
        "type": "docker",
        "status": "success",
        "errignore": false,
        "exit_code": 0,
        "machine": "15e89c0f84f1",
        "os": "linux",
        "arch": "amd64",
        "started": 1564085874,
        "stopped": 1564086343,
        "created": 1564085874,
        "updated": 1564086343,
        "version": 4,
        "on_success": true,
        "on_failure": false,
        "steps": [
          {
            "id": 575525,
            "step_id": 199937,
            "number": 1,
            "name": "clone",
            "status": "success",
            "exit_code": 0,
            "started": 1564085876,
            "stopped": 1564085901,
            "version": 4
          },
          {
            "id": 575526,
            "step_id": 199937,
            "number": 2,
            "name": "test",
            "status": "success",
            "exit_code": 0,
            "started": 1564085901,
            "stopped": 1564085994,
            "version": 4
          }
        ]
      }
    ]
  }
}
```
