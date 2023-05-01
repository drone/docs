---
date: 2000-01-01T00:00:00+00:00
title: Exec Pipeline
title_in_card: Exec Pipeline Specification
author: bradrydzewski
weight: 4
toc: true
type: spec

description: |
  Exec pipeline specification document.
---

This document introduces the data structures that represent the _exec pipeline_. The exec pipeline is a continuous integration pipeline that executes commands directly on the host machine.

<a id="the-resource-interface"></a>

# The `Resource` interface

The [`Resource`](#the-resource-interface) interface is implemented by all top-level objects, including the exec [`Pipeline`](#the-pipeline-object).

```typescript {linenos=table}
interface Resource {
  kind: string;
  type: string;
  name: string;
  concurrency: Concurrency;
  depends_on: string[];
}
```

<a id="the-kind-attribute"></a>

## The `kind` attribute

Defines the kind of resource, used to identify the resource implementation. This attribute is of type `string` and is required.

<a id="the-type-attribute"></a>

## The `type` attribute

Defines the type of resource, used to identify the resource implementation. This attribute is of type `string` and is required.

<a id="the-name-attribute"></a>

## The `name` attribute

The name of the resource. This value is required and must match `[a-zA-Z0-9_-]`. This value is displayed in the user interface (non-normative) and is used to identify the pipeline (non-normative).

<a id="the-concurrency-attribute"></a>

## The `concurrency` attribute

Defines the concurrency limits for the pipeline stage. This attribute is of type [`Concurrency`](#the-concurrency-object) and is optional.

<a id="the-depends_on-attribute"></a>

## The `depends_on` attribute

Defines a list of pipeline dependencies, used to defer execution of the pipeline until the named pipelines are in a completed state. This attribute is an array of type `string` and is optional.

<a id="the-pipeline-object"></a>

# The `Pipeline` object

The [`Pipeline`](#the-pipeline-object) is the top-level object used to represent the exec pipeline. The [`Pipeline`](#the-pipeline-object) object implements the [`Resource`](#the-resource-interface) interface.

```typescript {linenos=table}
class Pipeline : Resource {
  kind:      string;
  type:      string;
  name:      string
  platform:  Platform;
  workspace: Workspace;
  clone:     Clone;
  steps:     Step[];
  trigger:   Conditions;
}
```

<a id="the-kind-attribute"></a>

## The `kind` attribute

The kind of resource. This value must be set to `pipeline`.

<a id="the-type-attribute"></a>

## The `type` attribute

The type of resource. This value must be set to `exec`.

<a id="the-platform-section"></a>

## The `platform` section

The target operating system and architecture on which the pipeline must execute. This attribute is of type [`Platform`](#the-platform-object) and is recommended. If empty, the default operating system and architecture may be `linux` and `amd64` respectively.

<a id="the-workspace-section"></a>

## The `workspace` section

The working directory where the source code is cloned and the default working directory for each pipeline step. This attribute is of type [`Workspace`](#the-workspace-object) and is optional.

<a id="the-clone-section"></a>

## The `clone` section

Defines the pipeline clone behavior and can be used to disable automatic cloning. This attribute is of type [`Clone`](#the-clone-object) and is optional.

<a id="the-steps-section"></a>

## The `steps` section

Defines the pipeline steps. This attribute is an array of type [`Step`](#the-step-object) and is required. The array must not be empty and the order of the array must be retained.

<a id="the-trigger-section"></a>

## The `trigger` section

The conditions used to determine whether or not the pipeline should be skipped. This attribute is of type [`Conditions`](#the-conditions-object) and is optional.

<a id="the-platform-object"></a>

# The `Platform` object

The [`Platform`](#the-platform-object) object defines the target os and architecture for the pipeline and is used to route the pipeline to the correct instance (non-normative).

```typescript {linenos=table}
class Platform {
  os:      OS;
  arch:    Arch;
  variant: string;
  version: string;
}
```

<a id="the-os-attribute"></a>

## The `os` attribute

Defines the target operating system. The attribute is an enumeration of type `OS` and is recommended. If empty the operating system may default to Linux.

<a id="the-arch-attribute"></a>

## The `arch` attribute

Defines the target architecture. The attribute is an enumeration of type `Arch` and is recommended. If empty the architecture may default to amd64.

<a id="the-variant-attribute"></a>

## The `variant` attribute

Defines the architecture variant. This is most commonly used in conjunction with the arm architecture (non-normative) and can be used to differentiate between armv7, armv8, and so on (non-normative).

<a id="the-version-attribute"></a>

## The `version` attribute

Defines the operating system version. This is most commonly used in conjunction with the windows operating system (non-normative) and can be used to differentiate between 1809, 1903, and so on (non-normative).

<a id="the-clone-object"></a>

# The `Clone` object

The [`Clone`](#the-clone-object) object defines the clone behavior for the pipeline.

```typescript {linenos=table}
class Clone {
  depth:   number;
  disable: boolean;
}
```

<a id="the-depth-attribute"></a>

## The `depth` attribute

Configures the clone depth. This is an optional `number` value. If empty the full repository may be cloned (non-normative).

<a id="the-disable-attribute"></a>

## The `disable` attribute

Disables cloning the repository. This is an optional `boolean` value. It can be useful when you need to disable implement your own custom clone logic (non-normative).

<a id="the-step-object"></a>

# The `Step` object

The `Step` object defines a pipeline step.

```typescript {linenos=table}
class Step {
  name:        string;
  failure:     Failure;
  commands:    string[];
  environment: [string, string];
  when:        Conditions;
  depends_on:  string[];
}
```

<a id="the-name-attribute"></a>

## The `name` attribute

The name of the step. This value is required and must match [a-zA-Z0-9_-]. This value is displayed in the user interface (non-normative) and is used to identify the step (non-normative).

<a id="the-failure-attribute"></a>

## The `failure` attribute

Defines how the system handles failure. The default value is `always` indicating a failed step always fails the overall pipeline. A value of `ignore` indicates the failure is ignored. This attribute is of enumeration `Failure` and is optional.

<a id="the-commands-attribute"></a>

## The `commands` attribute

Defines a list of shell commands executed on the host machine. This attribute is an array of type `string` and is required.

<a id="the-environment-attribute"></a>

## The `environment` attribute

Defines a list of environment variables scoped to the pipeline step. This attribute is of type `[string, string]` and is optional.

<a id="the-when-section"></a>

## The `when` section

The conditions used to determine whether or not the step should be skipped. This attribute is of type [`Conditions`](#the-conditions-object) and is optional.

<a id="the-depends_on-attribute"></a>

## The `depends_on` attribute

Defines a list of steps dependencies, used to defer step execution until the named steps are in a completed state. This attribute is of type `string` and is optional.

<a id="the-conditions-object"></a>

# The `Conditions` object

The [`Conditions`](#the-conditions-object) object defines a set of conditions. If any condition evaluates to true its parent object is skipped.

```typescript {linenos=table}
class Conditions {
  action:   Constraint | string[];
  branch:   Constraint | string[];
  cron:     Constraint | string[];
  event:    Constraint | Event[];
  instance: Constraint | string[];
  ref:      Constraint | string[];
  repo:     Constraint | string[];
  status:   Constraint | Status[];
  target:   Constraint | string[];
}
```

<a id="the-action-attribute"></a>

## The `action` attribute

Defines matching criteria based on the build action. The build action is synonymous with a webhook action (non-normative). This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-branch-attribute"></a>

## The `branch` attribute

Defines matching criteria based on the git branch. This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-cron-attribute"></a>

## The `cron` attribute

Defines matching criteria based on the cron job that triggered the build. This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-event-attribute"></a>

## The `event` attribute

Defines matching criteria based on the build event. The build event is synonymous with a webhook event (non-normative). This attribute is of type [`Constraint`](#the-constraint-object) or an array of type [`Event`](#the-event-enum) and is optional.

<a id="the-instance-attribute"></a>

## The `instance` attribute

Defines matching criteria based on the instance hostname. This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-ref-attribute"></a>

## The `ref` attribute

Defines matching criteria based on the git reference. This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-repo-attribute"></a>

## The `repo` attribute

Defines matching criteria based on the repository name. This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-status-attribute"></a>

## The `status` attribute

Defines matching criteria based on the pipeline status. This attribute is of type [`Constraint`](#the-constraint-object) or an array of type [`Status`](#the-status-enum) and is optional.

<a id="the-target-attribute"></a>

## The `target` attribute

Defines matching criteria based on the target environment. The target environment is typically defined by a promote or rollback event (non-normative). This attribute is of type [`Constraint`](#the-constraint-object) or an array of type `string` and is optional.

<a id="the-constraint-object"></a>

# The `Constraint` object

The [`Constraint`](#the-constraint-object) object defines pattern matching criteria. If the pattern matching evaluates to false, the parent object is skipped.

```typescript {linenos=table}
class Constraint {
  exclude: string[];
  include: string[];
}
```

<a id="the-include-attribute"></a>

## The `include` attribute

List of matching patterns. If no pattern is a match, the parent object is skipped. This attribute is an array of type `string` and is optional.

<a id="the-exclude-attribute"></a>

## The `exclude` attribute

List of matching patterns. If any pattern is a match, the parent object is skipped. This attribute is an array of type `string` and is optional.

<a id="the-concurrency-object"></a>

# The `Concurrency` object

The [`Concurrency`](#the-concurrency-object) object defines the concurrency limits for the named pipeline.

```typescript {linenos=table}
class Concurrency {
  limit: number;
}
```

<a id="the-workspace-object"></a>

# The `Workspace` object

The [`Workspace`](#the-workspace-object) object defines the path to which the source code is cloned (non-normative) and the default working directory for each pipeline step (non-normative).

```typescript {linenos=table}
class Workspace {
  path: string;
}
```

# Enums

<a id="the-event-enum"></a>

## The `Event` enum

The `Event` enum provides a list of pipeline events. This value represents the event that triggered the pipeline.

```typescript {linenos=table}
enum Event {
  cron,
  custom,
  promote,
  pull_request,
  push,
  rollback,
  tag,
}
```

<a id="the-status-enum"></a>

## The `Status` enum

The `Status` enum provides a list of pipeline statuses. The default pipeline state is `success`, even if the pipeline is still running.

```typescript {linenos=table}
enum Status {
  failure,
  success,
}
```

<a id="the-failure-enum"></a>

## The `Failure` enum

The `Failure` enum defines a list of failure behaviors. The value `always` indicates a failure will fail the parent process. The value `ignore` indicates the failure is silently ignored.

```typescript {linenos=table}
enum Failure {
  always,
  ignore,
}
```

<a id="the-os-enum"></a>

## The `OS` enum

The `OS` enum provides a list of supported operating systems.

```typescript {linenos=table}
enum OS {
  darwin,
  dragonfly,
  freebsd,
  linux,
  netbsd,
  openbsd,
  solaris,
  windows,
}
```

<a id="the-arch-enum"></a>

## The `Arch` enum

The `Arch` enum provides a list of supported chip architectures.

```typescript {linenos=table}
enum Arch {
  386,
  amd64,
  arm64,
  arm,
}
```
