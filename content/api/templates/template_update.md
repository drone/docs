+++
date = "2000-01-01T00:00:00+00:02"
title = "Template Update"
description = "Endpoint to update a template"
+++

Updates the specified repository secret.
Please note this api requires write access to the repository,
and the request parameter `{secret}` is not the secret's id but secret name.

```
PATCH /api/templates/{namespace}/{name}
```

Example Request Body:

```json {linenos=table}
{
  "data": "updated_data"
}
```

Example Response Body:

```json {linenos=table}
{ 
    "id": 1,
    "name": "plugin.starlark",
    "namespace": "octocat",
    "data": "updated_data"
}
```
