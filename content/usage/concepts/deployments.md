+++
date = "2017-04-15T14:39:04+02:00"
title = "Promoting"
url = "promoting-builds"

[menu.usage]
  weight = 8
  identifier = "deployment"
  parent = "usage_concepts"
+++

Drone provides the ability to promote individual commits or tags (e.g. promote to production). When you promote a commit or tag it triggers a new pipeline execution with event type `deployment`. You can use the event type and target environment to limit step execution.

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

  publish:
    image: plugins/docker
    registry: registry.heroku.com
    repo: registry.heroku.com/my-staging-app/web
    when:
+     event: deployment
+     environment: staging

  publish_to_prod:
    image: plugins/docker
    registry: registry.heroku.com
    repo: registry.heroku.com/my-production-app/web
    when:
+     event: deployment
+     environment: production
```

The above example demonstrates how we can configure pipeline steps to only execute when the deployment matches a specific target environment.


# Trigger Deployments

Deployments are triggered from the command line utility. They are triggered from an existing build. This is conceptually similar to promoting builds.

```text
drone deploy <repo> <build> <environment>
```

```text
For drone cli version 1.0 or above
drone build promote <repo> <build> <environment>
```

Promote the specified build number to your staging environment:

```text
drone deploy octocat/hello-world 24 staging
```
```text
For drone cli version 1.0 or above
drone build promote octocat/hello-world 24 staging
```

Promote the specified build number to your production environment:

```text
drone deploy octocat/hello-world 24 production
```
```text
For drone cli version 1.0 or above
drone build promote octocat/hello-world 24 production
```
