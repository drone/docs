+++
date = "2017-04-15T14:39:04+02:00"
title = "Pipelines"
url = "pipelines"

[menu.usage]
  weight = 4
  identifier = "pipelines"
  parent = "usage_concepts"
+++

The pipeline section defines a list of steps to build, test and deploy your code. Pipeline steps are executed serially, in the order in which they are defined. If a step returns a non-zero exit code, the pipeline immediately aborts and returns a failure status.

Example pipeline:

```yaml
pipeline:
  backend:
    image: golang
    commands:
      - go build
      - go test
  frontend:
    image: node
    commands:
      - npm install
      - npm run test
      - npm run build
```

In the above example we define two pipeline steps, `frontend` and `backend`. The names of these steps are completely arbitrary.

# Build Steps

Build steps are steps in your pipeline that execute arbitrary commands inside the specified docker container. The commands are executed using the workspace as the working directory.

```diff
pipeline:
  backend:
    image: golang
    commands:
+     - go build
+     - go test
```

There is no magic here. The above commands are converted to a simple shell script. The commands in the above example are roughly converted to the below script:

```diff
#!/bin/sh
set -e

go build
go test
```

The above shell script is then executed as the docker entrypoint. The below docker command is an (incomplete) example of how the script is executed:

```
docker run --entrypoint=build.sh golang
```

{{% alert info %}}
Please note that only build steps can define commands. You cannot use commands with plugins or services.
{{% /alert %}}

# Parallel Execution

Drone supports parallel step execution for same-machine fan-in and fan-out. Parallel steps are configured using the `group` attribute. This instructs the pipeline runner to execute the named group in parallel.

Example parallel configuration:

```diff
pipeline:
  backend:
+   group: build
    image: golang
    commands:
      - go build
      - go test
  frontend:
+   group: build
    image: node
    commands:
      - npm install
      - npm run test
      - npm run build
  publish:
    image: plugins/docker
    repo: octocat/hello-world
```

In the above example, the `frontend` and `backend` steps are executed in parallel. The pipeline runner will not execute the `publish` step until the group completes.

# Conditional Pipeline Execution

Drone gives the ability to skip commits based on the target branch. The below example will skip a commit when the target branch is not master.

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

+branches: master
```

Please see the pipeline conditions [documentation]({{< ref "usage/config/pipeline-conditions.md" >}}) for more options and details.

# Conditional Step Execution

Drone gives the ability to conditionally limit the execution of steps at runtime. The below example limits execution of Slack plugin steps based on branch:

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
+   when:
+     branch: master
```

Please see the step conditions [documentation]({{< ref "usage/config/step-conditions.md" >}}) for more options and details.

# Failure Execution

Drone uses the container exit code to determine the success or failure status of a build. Non-zero exit codes fail the build and cause the pipeline to immediately exit.

There are use cases for executing pipeline steps on failure, such as sending notifications for failed builds. Use the status constraint to override the default behavior and execute steps even when the build status is failure:

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
+   when:
+     status: [ success, failure ]
```

Please see the step conditions [documentation]({{< ref "usage/config/step-conditions.md" >}}) for more options and details.
