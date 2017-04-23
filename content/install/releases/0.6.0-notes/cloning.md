+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Cloning Changes"
url = "release-0.6.0-cloning"
+++

# Overview

The clone section is removed from the pipeline and has its own section in the yaml. Please note this change is experimental and the naming could be adjusted in the future.

```diff
+clone:
+  git:
+    image: plugins/git
+    depth: 50

pipeline:
- clone:
-   image: plugins/git
-   depth: 50
```

Please note that if you are using a plugin to cache your git repository you will need to adjust your pipeline configuration accordingly

```diff
+clone:
+  from_cache:
+    image: drillster/drone-volume-cache
+    restore: true
+  git:
+    image: plugins/git
+    depth: 50

pipeline:
- restore-cache:
-   image: drillster/drone-volume-cache
-   restore: true

- clone:
-   image: plugins/git
-   depth: 50
```
