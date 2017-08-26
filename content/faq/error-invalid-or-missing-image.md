---
date: 2017-04-15T14:39:04+02:00
title: "Error: Invalid or missing Image"
url: error-invalid-or-missing-image
---

This error message occurs when your yaml is malformed and does not properly define an image for a pipeline step or service.

You must define an image for every pipeline step:

```diff
pipeline:
  slack:
+   image: plugins/slack
    room: foo
```

You must define an image for every service:

```diff
services:
  mysql:
+   image: mysql
```

You may not nest steps in your pipeline:

```diff
pipeline:
- notify:
-   slack:
-     image: plugins/slack
-     room: foo
+ slack:
+   image: plugins/slack
+   room: foo
```
