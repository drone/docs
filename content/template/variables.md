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
Must be prefixed with 'build.'

Jsonnet example usage:
```
local event = std.extVar("build.event");
local action = std.extVar("build.action");
```

Starlark example usage:
```
{
 "event": ctx.build.event,
 "action": ctx.build.action
}
```

Variable	| Description
------------|------------
event	      | Provides the event that triggered the pipeline execution.
action	      | Provides the action that triggered the pipeline execution. Use this value to differentiate between the pull request being opened vs synchronized.
environment	  | Provides the target deployment environment for the running build.
link	      | Provides a link the git commit or object in the source control management system.
branch	      | Provides the target branch for the push or pull request. This value may be empty for tag events.
source	      | Provides the source branch for the pull request. This value may be empty for certain source control management providers.
before	      | Provides the git commit sha before the patch is applied. This may be used in conjunction with the after commit sha to create a diff.
after	      | Provides the git commit sha after the patch is applied. This may be used in conjunction with the before commit sha to create a diff.
target	      | Provides the target branch for the push or pull request. This value may be empty for tag events.
ref	          | Provides the target branch for the push or pull request. This value may be empty for tag events.
commit	      | Provides the git commit sha after the patch is applied. This may be used in conjunction with the before commit sha to create a diff.
title	      | Provides the title of the commit. The full first line of the message.
message	      | Provides the commit message for the current running build.
source_repo	  | Provides the source repository name of the pull request.
author_login  | Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username).
author_name   | Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate.
author_email  | Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate.
author_avatar | Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub).
sender	      | Provides the event sender’s login name

# Repo
Must be prefixed with 'repo.'

Jsonnet example usage:
```
local uid = std.extVar("repo.uid");
local name = std.extVar("repo.name");
 ```

Starlark example usage:
```
{
 "uid": ctx.repo.uid,
 "name": ctx.repo.name
}
```

Variable	| Description
------------|------------
uid	         | Provides the UUID of the repository.
name	     | Provides the full repository name for the current running build.
namespace	 | Provides the repository namespace for the current running build. The namespace is an alias for the source control management account that owns the repository.
slug	     |
git_http_url | Provides the git+http url that should be used to clone the repository.
git_ssh_url	 | Provides the git+ssh url that should be used to clone the repository.
link	     | Provides the repository link for the current running build.
branch	     | Provides the default repository branch for the current running build.
config	     | Provides the configuration path for the repository.
private	     | Provides a boolean flag that indicates whether or not the repository is private or public.
visibility	 | Provides the repository visibility level for the current running build.
active	     | Provides a boolean flag that indicates whether or not the repository is active.
trusted	     | Provides a boolean flag that indicates whether or not the repository is trusted.
protected	 | Provides a boolean flag that indicates whether or not the repository is protected.
ignore_forks | Provides a boolean flag that indicates whether or not to ignore forks.
ignore_pull_requests  | Provides a boolean flag that indicates whether or not to ignore pull requests.