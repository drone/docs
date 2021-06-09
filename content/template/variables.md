---
date: 2000-01-01T00:00:00+00:00
title: Variables
author: eoinmcafee
weight: 20
disable_toc: true
aliases:
- /configure/templates/variables/ 
description: a list and description of variables available for use
---
The following variables are available for use for both starlark & Jsonnet templates:

# Build

Jsonnet example usage:
```
local event = std.extVar("build.event");
local action = std.extVar("build.action");

please note they must be prefixed with 'build.'
```

Starlark example usage:
```
{
 "event": ctx.build.event,
 "action": ctx.build.action
}
please note they must be prefixed with 'build.'
```

Variable	| Description
------------|------------
event	      | 
action	      |
environment	  |
link	      |
branch	      |
source	      |
before	      |
after	      |
target	      |
ref	          |
commit	      |
title	      |
message	      |
source_repo	  |
author_login  |
author_email  |
author_avatar |
sender	      |

# Repo

Jsonnet example usage:
```
local uid = std.extVar("repo.uid");
local name = std.extVar("repo.name");

please note they must be prefixed with 'repo.'
```

Starlark example usage:
```
{
 "uid": ctx.repo.uid,
 "name": ctx.repo.name
}
please note they must be prefixed with 'repo.'
```

Variable	| Description
------------|------------
uid	         | 
name	     |
namespace	 |
slug	     |
git_http_url |
git_ssh_url	 |
link	     |
branch	     |
config	     |
private	     |
visibility	 |
active	     |
trusted	     |
protected	 |
ignore_forks |
ignore_pull_requests  |