---
date: 2000-01-01T00:00:00+00:00
title: Jsonnet
author: bradrydzewski
weight: 1
aliases:
- /extend/config/jsonnet/
description: |
  Overview of Jsonnet scripting.
---

Drone supports Jsonnet scripting as an alternate to yaml configurations. Jsonnet is a data templating language that extends Json syntax, adding constructs for generating, translating and refining data.

{{< link-to "https://jsonnet.org/" >}}
Jsonnet Language Specification
{{</ link-to >}}

Jsonnet is intended for projects with complex configurations that benefit from advanced scripting capabilities and code re-use.

# Usage

<div class="alert">
In order to use Jsonnet configuration files your system administrator must <a href="{{< relref "server/reference/drone-jsonnet-enabled.md" >}}">enable</a> Jsonnet support for your Drone server.
</div>

You can use Jsonnet for an individual project by creating a `.drone.jsonnet` file in the root of your git repository. Then update your repository configuration file accordingly, from your repository settings screen.
<!--
# Structure

The Jsonnet script returns one or many pipeline objects from the script's main method. The pipeline object uses the same structure as a yaml pipeline. In fact, under the covers, the returned object is converted directly to a yaml document.

Example Starlark script:

{{< highlight text "linenos=table" >}}
def main(ctx):
  return {
    "kind": "pipeline",
    "name": "build",
    "steps": [
      {
        "name": "build"
        "image": "alpine"
        "commands": [
            "echo hello world"
        ]
      }
    ]
  }
{{< / highlight >}}

Equivalent Yaml configuartion:

{{< highlight text "linenos=table" >}}
---
kind: pipeline
name: build

steps:
- name: build
  image: alpine
  commands:
  - echo hello world
{{< / highlight >}}
-->
# Example

Here is an example script that returns a pipeline configuration. _Please note the returned pipeline object uses the same structure as a pipeline defined in yaml._

{{< highlight text "linenos=table" >}}
{
    "kind": "pipeline",
    "type": "docker",
    "name": "default",
    "steps": [
        {
            "name": "build",
            "image": "alpine",
            "commands": [
                "echo hello world",
            ]
        }
    ]
}
{{< / highlight >}}

# Tooling

You can automatically convert your Jsonnet configuration file to yaml using the command line tools. This can be useful for local testing.

```
$ drone jsonnet --stdout

kind: pipeline
name: default

steps:
- name: build
  image: alpine
  commands: [ echo hello world ]
```

# Common Problems

The below error may indicate that Drone does not recognize your configuration as a jsonnet file.  The most common root cause for this problem is when you forget to [enable jsonnet]({{< relref "server/reference/drone-jsonnet-enabled.md" >}}) in your Drone server settings.

```
yaml: line 1: mapping values are not allowed in this context
```

The second most common root cause for this issue is when you forget to rename your file with a jsonnet extension in your repository and in your repository settings screen in Drone (i.e. rename from .drone.yml to .drone.jsonnet). Drone assumes configuration files are written in yaml unless you explicitly use the jsonnet file extension.
