+++
date = "2000-01-01T00:00:00+00:02"
title = "Current User Repos"
+++

Returns the currently authenticated user's repository list.

```
GET /api/user/repos
```

Include the latest build for each active repository.

```
GET /api/user/repos?latest=true
```