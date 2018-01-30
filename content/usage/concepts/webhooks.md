+++
date = "2017-04-15T14:39:04+02:00"
title = "Webhooks"
url = "hooks"

[menu.usage]
  weight = 1
  identifier = "hooks"
  parent = "usage_concepts"
+++

When you activate your repository Drone automatically add webhooks to your version control system (e.g. GitHub). There is no manual configuration required.

Webhooks are used to trigger pipeline executions. When you push code to your repository, open a pull request, or create a tag, your version control system will automatically send a webhook to Drone which will in turn trigger pipeline execution.

# Required Permissions

The user who enables a repo in Drone must have `Admin` rights on that repo, so that Drone can add the webhook.  

Note that manually creating webhooks yourself is not possible. This is because webhooks are signed using a per-repository secret key which is not exposed to end users. 

<!-- # Recreate Webhooks

Drone provides the ability to recreate webhooks, in case they were accidentally removed or altered, using the command line utility.

```text
drone repo repair <repo>
drone repo repair octocat/hello-world
``` -->

# Skip Commits

Drone gives the ability to skip individual commits by adding `[CI SKIP]` to the commit message. Note this is case-insensitive.

```diff
git commit -m "updated README [CI SKIP]"
```

# Skip Branches

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
