+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Secrets"
url = "release-0.6.0-secrets"
+++

Previously secrets were interpolated in the yaml using `${variable}` syntax. In the past we have attempted alternate approaches to limit secret expose to specific images without success. This release adopts the docker-compose syntax for exposing secrets to your pipeline steps.

The pipeline steps request access to named secrets. If access is granted, secrets are loaded as environment variables. Support for loading secrets from a temporary file are planned.

Example configuration:

```diff
pipeline:
  notify:
    image: plugins/slack
    channel: drone
-   webhook: ${slack_webhook}
+   secrets: [ slack_webhook ]
```

Secret interpolation is being removed in favor of this docker-compose syntax. This implementation improves upon prior attempts and eliminates previously encountered [design flaws](https://github.com/drone/drone/issues/1888).

# Commands

The command line utility can be used to manage secrets. These commands have slightly changed, using explicit flags for providing the secret name and value.

```text
drone secret add <repo> \
  --image=<image> \
  --name=<secret> \
  --value=<value>
```

Example command:

```text
drone secret add octocat/hello-world \
  --image=plugins/slack \
  --name=slack_webhook \
  --value=https://hooks.slack.com/services/...
```

# Matching

Prior implementations had very literal matching logic. The matching logic has been simplified to match the literal image name and ignore tags. Wildcards are no longer supported. Please update your secrets accordingly.

```diff
drone secret update octocat/hello-world \
- --image=plugins/slack:* \
+ --image=plugins/slack \
  --name=slack_webhook \
  --value=https://hooks.slack.com/services/...
```

# Plugins

Secrets are passed to pipeline steps (including plugins) as environment variables. Plugin authors must update their plugins to source secrets from specific environment variables.

Example slack plugin changes:

```diff
cli.StringFlag{
  Name:   "webhook",
  Usage:  "slack webhook url",
- EnvVar: "PLUGIN_WEBHOOK",
+ EnvVar: "SLACK_WEBHOOK,PLUGIN_WEBHOOK",
},
```

Example yaml configuration:

```diff
pipeline:
  notify:
    image: plugins/slack
    channel: drone
+   secrets: [ slack_webhook ]
```

Please note that plugin authors need to document the names of the environment variables from which they will source these values. Users will need to create secrets using these exact environment variable names.

# Global and Organization Secrets

Global and organization secrets are removed from Drone core and are delegated to plugins, using the new pluggable [secret backend](https://github.com/drone/drone/issues/1997). Commands for managing global and organization secrets have been removed accordingly.

```diff
drone secret add <repo> <name> <value>
- drone org secret add <org> <name> <value>
- drone global secret add <name> <value>
```

This decision was made because organization secrets were not fully integrated and only supported GitHub (GitLab, Gitea, Gogs, Bitbucket, and Stash were not fully supported). There were also conflicting requirements from organizations regarding access control and storage, including desire to use third party secret management tools (e.g. Vault).

This change reflects the fact that Drone is moving toward a more focused core, where advanced functionality and customization is delegated to plugins. With pluggable secret backends, organizations have fine grained control over where secrets are stored and which builds have access to those secrets.
