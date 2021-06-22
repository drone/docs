+++
date = "2000-01-01T00:00:00+00:02"
title = "Template Info"
description = "Endpoint to find a template"
+++

Returns the template.
Please note this api requires write access to the repository,

```
GET /api/templates/{namespace}/{name}
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
