---
date: 2000-01-01T00:00:00+00:00
title: Starlark
author: bradrydzewski
weight: 1
toc: true
aliases:
- /starlark/overview/
description: |
  Overview of Starlark scripting.
---

Drone supports Starlark scripting as an alternate to yaml configurations. Starlark is a dialect of Python intended for use as a configuration language. The language is developed by Google for the [Bazel](https://bazel.build/) build tool.

{{< link-to "https://github.com/google/starlark-go/blob/master/doc/spec.md" >}}
Starlark Language Specification
{{</ link-to >}}

Starlark is intended for projects with complex configurations that benefit from a true scripting language.

# Usage

<div class="alert">
In order to use Starlark configuration files your system administrator must <a href="{{< relref "server/reference/drone-starlark-enabled.md" >}}">enable</a> Starlark support for your Drone server.
</div>

You can use Starlark for an individual project by creating a `.drone.star` file in the root of your git repository. Then update your repository configuration file accordingly, from your repository settings screen.

# Structure

The Starlark script returns one or many pipeline objects from the script's main method. The pipeline object uses the same structure as a yaml pipeline. In fact, under the covers, the returned object is converted directly to a yaml document.

Example Starlark script:

{{< highlight text "linenos=table" >}}
def main(ctx):
  return {
    "kind": "pipeline",
    "name": "build",
    "steps": [
      {
        "name": "build",
        "image": "alpine",
        "commands": [
            "echo hello world"
        ]
      }
    ]
  }
{{< / highlight >}}

Equivalent Yaml configuration:

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

# Example

Here is an example configuration script that returns a pipeline configuration. _Please note the returned pipeline object uses the same structure as a pipeline defined in yaml._

{{< highlight text "linenos=table" >}}
def main(ctx):
  return {
    "kind": "pipeline",
    "name": "build",
    "steps": [
      {
        "name": "build",
        "image": "alpine",
        "commands": [
            "echo hello world"
        ]
      }
    ]
  }
{{< / highlight >}}

The main function can also return an array. In the below example we create two pipelines, one for amd64 and one for arm.

{{< highlight text "linenos=table,hl_lines=2-5" >}}
def main(ctx):
  return [
    step("amd64"),
    step("arm64"),
  ]

def step(arch):
  return {
    "kind": "pipeline",
    "name": "build-%s" % arch,
    "steps": [
      {
        "name": "build",
        "image": "alpine",
        "commands": [
            "echo hello world"
        ]
      }
    ]
  }
{{< / highlight >}}

The above example also demonstrates how we can use functions and scripting to build the configuration without repetition.

# Context

The main method of the Starlark script accepts a context parameter. The context parameter provides access to repository and build meta-data, which can be used to dynamically build your pipeline configuration.

{{< highlight text "linenos=table,hl_lines=2" >}}
def main(ctx):
  if ctx.build.event == "tag":
    return pipeline_for_release(ctx)
  else:
    return pipeline(ctx)
{{< / highlight >}}

_Please note that you can still create pipeline objects with `trigger` stanzas and pipeline steps with `when` stanzas._

# Tooling

You can automatically convert your Starlark configuration file to yaml using the command line tools. This can be useful for local testing.

```
$ drone starlark convert --stdout

kind: pipeline
name: default

steps:
- name: build
  image: alpine
  commands: [ echo hello world ]
```