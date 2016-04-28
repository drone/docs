+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Environment"
weight = 3
menu = "usage"
toc = true
+++

# Environment

Environment variables available to your build and plugin containers:

NAME                               | DESC
-----------------------------------|--------------------------------------------
`CI=drone`                         | environment is drone
`DRONE=true`                       | environment is drone
`DRONE_ARCH`                       | environment architecture (linux/amd64)
`DRONE_REPO`                       | repository full name
`DRONE_REPO_OWNER`                 | repository owner
`DRONE_REPO_NAME`                  | repository name
`DRONE_REPO_SCM`                   | repository scm (git)
`DRONE_REPO_LINK`                  | repository link
`DRONE_REPO_AVATAR`                | repository avatar
`DRONE_REPO_BRANCH`                | repository default branch (master)
`DRONE_REPO_PRIVATE`               | repository is private
`DRONE_REPO_TRUSTED`               | repository is trusted
`DRONE_REMOTE_URL`                 | repository clone url
`DRONE_COMMIT_SHA`                 | commit sha
`DRONE_COMMIT_REF`                 | commit ref
`DRONE_COMMIT_BRANCH`              | commit branch
`DRONE_COMMIT_LINK`                | commit link in remote
`DRONE_COMMIT_MESSAGE`             | commit message
`DRONE_COMMIT_AUTHOR`              | commit author username
`DRONE_COMMIT_AUTHOR_EMAIL`        | commit author email address
`DRONE_COMMIT_AUTHOR_AVATAR`       | commit author avatar
`DRONE_BUILD_NUMBER`               | build number
`DRONE_BUILD_EVENT`                | build event (push, pull_request, tag)
`DRONE_BUILD_STATUS`               | build status (success, failure)
`DRONE_BUILD_LINK`                 | build result link
`DRONE_BUILD_CREATED`              | build created unix timestamp
`DRONE_BUILD_STARTED`              | build started unix timestamp
`DRONE_BUILD_FINISHED`             | build finished unix timestamp
`DRONE_PREV_BUILD_STATUS`          | prior build status
`DRONE_PREV_BUILD_NUMBER`          | prior build number
`DRONE_PREV_COMMIT_SHA`            | prior build commit sha
`DRONE_YAML_SIGNED`                | yaml is signed
`DRONE_YAML_VERIFIED`              | yaml is signed and verified
`DRONE_BRANCH`                     | commit branch
`DRONE_COMMIT`                     | commit sha
`DRONE_TAG`                        | commit tag
`DRONE_PULL_REQUEST`               | pull request number
`DRONE_DEPLOY_TO`                  | deployment target (ie production)


# String Interpolation

Environment variables are interpolated in the yaml using the `${VARIABLE}` syntax.

```yaml
  s3:
    source: archive.tar.gz
    target: archive_${DRONE_COMMIT}.tar.gz
```

# String Operations

Environment variable interpolation supports emulated bash string operations:


OPERATION             | DESC
----------------------|---------------------------------------------------------
`${param}`            | parameter substitution
`"${param}"`          | parameter substitution with escaping
`${param:pos}`        | parameter substitution with substring
`${param:pos:len}`    | parameter substitution with substring
`${param=default}`    | parameter substitution with default
`${param##prefix}`    | parameter substitution with prefix removal
`${param%%suffix}`    | parameter substitution with suffix removal
`${param/old/new}`    | parameter substitution with find and replace

Example operation to shorten the git sha value:

```
${DRONE_COMMIT:0:8}
```

Example operation to remove the `v` from git tag value `v1.0.0`:

```
${DRONE_TAG##v}
```
