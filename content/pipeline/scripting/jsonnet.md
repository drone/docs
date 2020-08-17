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
In order to use Jsonnet configuration files your system administrator must enable Jsonnet support for your Drone server.
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