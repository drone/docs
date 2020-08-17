+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Delete"
description = "Endpoint to delete a repository secret"
+++

Deletes a repository secret.
Please note this api requires write access to the repository,
and the request parameter `{secret}` is not the secret's id but secret name.

```
DELETE /api/repos/{owner}/{repo}/secrets/{secret}
```
