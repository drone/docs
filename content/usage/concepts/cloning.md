+++
date = "2017-04-15T14:39:04+02:00"
title = "Cloning"
url = "cloning"

[menu.usage]
  weight = 3
  identifier = "cloning"
  parent = "usage_concepts"
+++

Drone automatically configures a default clone step if not explicitly defined. You can manually configure the clone step in your pipeline for customization:

```diff
+clone:
+  git:
+    image: plugins/git

pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test
```

Example configuration to override depth:

```diff
clone:
  git:
    image: plugins/git
+   depth: 50
```

Example configuration to use a custom clone plugin:

```diff
clone:
  git:
+   image: octocat/custom-git-plugin
```

Example configuration to clone Mercurial repository:

```diff
clone:
  hg:
+   image: pluins/hg
+   path: bitbucket.org/foo/bar
```
