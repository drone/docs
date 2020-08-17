---
date: 2000-01-01T00:00:00+00:00
title: Promotions
author: bradrydzewski
weight: 150
toc: true

related:
  items:
  - name: Install the Command Line Tools
    path: /cli/install.md
  - name: Get the Parent Build Number
    path: /pipeline/environment/reference/drone-build-parent.md
  - name: Get the Target Environment
    path: /pipeline/environment/reference/drone-deploy-to.md

aliases:
- /promoting-builds/
- /user-guide/pipeline/promotion
---

You can use the promotion feature to deploy code to a target environment, for example,  _Promote build number 42 to Production_. Here are some of the benefits of using promotions:

* Create repeatable deployments
* Create an audit trail
* Reduce human error
* Segregation of duties
* Revoke developer access to server environments

# How it Works

When you promote a build, Drone creates a new build that inherits its values from the build being promoted. This new build's event type is set to _promote_. You can reference the event type and target environment in your pipeline configuration.

Limit pipeline execution by target environment:

```yaml {linenos="table",hl_lines=["17-21"]}
kind: pipeline
type: docker
name: default

steps:
- name: test
  image: node
  commands:
  - npm install
  - npm run test
  - npm run bundle

- name: deploy
  image: plugins/ssh
  settings: ...

trigger:
  event:
  - promote
  target:
  - production
```

Limit pipeline step execution by target environment:

```yaml {linenos="table",hl_lines=["16-20"]}
kind: pipeline
type: docker
name: deploy

steps:
- name: test
  image: node
  commands:
  - npm install
  - npm run test
  - npm run bundle

- name: deploy
  image: plugins/ssh
  settings: ...
  when:
    event:
    - promote
    target:
    - production
```

# How To Promote

You can promote builds using the command line utility:

* Use the build [promote]({{< relref "/cli/build/drone-build-promote.md" >}}) command:
    ```
    $ drone build promote <repo> <number> <environment>
    ```
* Example promotes build 42 to staging:
    ```
    $ drone build promote octocat/hello-world 42 staging
    ```
* Example promotes build 42 to production:
    ```
    $ drone build promote octocat/hello-world 42 production
    ```

_In the above examples we reference staging and production environments, however, the environment is a freeform field. You can use any environment name._
