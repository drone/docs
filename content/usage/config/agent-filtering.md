+++
date = "2018-09-15T14:39:04+02:00"
title = "Filtering builds by agent"
url = "filtering-builds-by-agent"

[menu.usage]
  weight = 2
  identifier = "filtering-builds-by-agent"
  parent = "usage_config"
+++

The first step is to define a query used to filter which builds the agent will accept. The query expression is similar to sql. For example:

```
platform = 'linux/amd64' AND repo GLOB octocat/*
```

There are two fields available by default:

* `repo` is the repository name
* `platform` is the target platform for the build

The user can define additional fields in the `labels` section of the `.drone.yml` file. These fields can then be evaluated by the query expression:

```
pipeline:
  build:
    image: golang
    commands: go test

labels:
  - ram=32
  - cores=16
```

For example:

```
DRONE_FILTER=platform = 'linux/amd64' AND ram >= 32 AND cores >= 16
```

**Caution**: The query expression must not be quoted in double quotes.

The query expression string is passed to the agent using the `DRONE_FILTER` environment variable. The agent will only process builds when the expression evaluates to true. In the previous example, the agent would only process builds for linux, with user-defined labels specifying 32 GB ram or higher and 16 cores or higher.

# Additional examples

Mostly taken from the expression parser [test suite](https://github.com/drone/expr/blob/master/selector_test.go):

```
# only if ram is bigger or equal to 32
ram >= 32

# only execute for drone repository which must be private and is from git VCS
repo-name == 'drone' AND repo-private == true AND repo-vcs == 'git'

# only execute builds of repositories owned by user-login
repo-owner == user-login

# unix-like glob every platform starting on linux/
platform GLOB 'linux/*'

# negation of the previous
platform NOT GLOB 'linux/*'

# classical regex mapping
platform REGEXP 'linux/(.+)'

# argument is between 2 (inclusive)
argument BETWEEN 2 AND 4
```
