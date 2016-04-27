+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Plugins"
weight = 27
menu = "usage"
toc = true
+++

# Overview

Deployments are triggered at the end of a successful build. Note that deployment steps are not executed for pull requests or failed builds.

Example Heroku deployment using the Heroku plugin:

```yaml
---
deploy:
  heroku:
    app: petstore
```

# Plugins

Drone does not have any builtin deployment capabilities. Drone instead offers a large library of plugins for deploying to [Heroku](/plugins/heroku/), [Tutum](/plugins/tutum/), [Amazon](/plugins/aws_codedeploy/) and more. Please see the [plugin documentation](/plugins/) for a complete list.


# Conditions

Use the `when` block to limit step execution:

```yaml
script:

  heroku:
    app: mycompany.com
    when:
      branch: master
```

Execute a step if the branch is `master` or `develop`:

```yaml
when:
  branch: [master, develop]
```

Execute a step if the branch is starts with `prefix/*`:

```yaml
when:
  branch: prefix/*
```

Execute a step if the build event is a `tag`:

```yaml
when:
  event: tag
```

Execute a step if the build event matches one of the values:

```yaml
when:
  event: [push, pull_request, tag, deployment]
```

Execute a step if the build event matches one of the values:

```yaml
when:
  event: [push, pull_request, tag, deployment]
```

Execute a step for a specific platform:

```yaml
when:
  platform: linux/amd64
```

Execute a step for a specific platform using wildcards:

```yaml
when:
  platform:  [ linux/*, windows/amd64 ]
```

Execute a step when the build status changes:

```yaml
when:
  status: changed
```

Execute a step when the build is passing or failing:

```yaml
when:
  status:  [ failure, success ]
```

Execute a step for deployment events matching the target environment:

```yaml
when:
  environment: production
  event: deployment
```


<!-- Define the same deployment step multiple times, using different configuration based on branch:

```yaml
script:

  staging:
    image: heroku
    app: test.mycompany.com
    when:
      branch: develop

  production:
    image: heroku
    app: mycompany.com
    when:
      branch: master
``` -->
