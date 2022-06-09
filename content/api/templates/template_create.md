+++
date = "2000-01-01T00:00:00+00:02"
title = "Template Create"
description = "Endpoint to add template"
+++

Create a new template.
Please note this api requires write access to the repository.

```
POST /api/templates/{namespace}
```

Example Request Body:

```json {linenos=table}
{
  "name": "name",
  "data": "data"
}
```

Example Response Body:

```json {linenos=table}
{
  "id": 1,
  "name": "plugin.starlark",
  "namespace": "octocat",
  "data": "data"
}
```
