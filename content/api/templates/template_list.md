+++
date = "2000-01-01T00:00:00+00:02"
title = "Template List"
description = "Endpoint to list organization templates"
+++

Returns the organization template list.
Please note this api requires write access to the repository.

```
GET /api/templates/{namespace}
```

Example Response Body:

```json {linenos=table}
[{
	"id": 1,
	"name": "plugin.starlark",
	"namespace": "octocat",
	"data": "data"
}]
```
