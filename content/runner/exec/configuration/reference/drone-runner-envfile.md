---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RUNNER_ENVFILE
author: bradrydzewski
weight: 1
---

Optional string. Provides a path to an environment file containing global environment variables that are injected into every pipeline step.

```
DRONE_RUNNER_ENVIRON=/path/to/file.env
```

The environment file is a text file that defines environment variables in key value format. Please see the envfile [documentation](https://github.com/joho/godotenv) for more details about the file format.

```
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```
