+++
date = "2017-04-15T14:39:04+02:00"
title = "Reference Environment"
url = "environment-reference"

[menu.usage]
  weight = 1
  parent = "usage_reference"
  identifier = "environment-reference"
+++

This is the reference list of all environment variables available to your build environment. These are injected into your build and plugins containers, at runtime.

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
`DRONE_BRANCH`                     | commit branch
`DRONE_COMMIT`                     | commit sha
`DRONE_TAG`                        | commit tag
`DRONE_PULL_REQUEST`               | pull request number
`DRONE_DEPLOY_TO`                  | deployment target (ie production)
