+++
title = "API Reference"
menu = "developers"
weight = 5
toc = true
class = "api"
+++


# Repos


## <span class="operation operation-GET">GET</span> <span>Get a repo</span>

Retrieves the details of a repository.

```
GET /api/repos/{owner}/{name}
```










Example Response Body:

```js
{
  "id": 1,
  "scm": "git",
  "owner": "octocat",
  "name": "hello-world",
  "full_name": "octocat/hello-world",
  "avatar_url": "https://avatars.githubusercontent.com/u/2181346?v=3",
  "link_url": "https://github.com/octocat/hello-world",
  "clone_url": "https://github.com/octocat/hello-world.git",
  "default_branch": "master",
  "timeout": 60,
  "private": false,
  "trusted": false,
  "allow_pr": true,
  "allow_push": true,
  "allow_deploys": false,
  "allow_tags": false
}

```







## <span class="operation operation-POST">POST</span> <span>Activates a repo</span>

Activates a repository.

```
POST /api/repos/{owner}/{name}
```










Example Response Body:

```js
{
  "id": 1,
  "scm": "git",
  "owner": "octocat",
  "name": "hello-world",
  "full_name": "octocat/hello-world",
  "avatar_url": "https://avatars.githubusercontent.com/u/2181346?v=3",
  "link_url": "https://github.com/octocat/hello-world",
  "clone_url": "https://github.com/octocat/hello-world.git",
  "default_branch": "master",
  "timeout": 60,
  "private": false,
  "trusted": false,
  "allow_pr": true,
  "allow_push": true,
  "allow_deploys": false,
  "allow_tags": false
}

```















## <span class="operation operation-PATCH">PATCH</span> <span>Updates a repo</span>

Updates the specified repository.

```
PATCH /api/repos/{owner}/{name}
```








Example Request Body:

```js
{
  "timeout": 60,
  "private": false,
  "trusted": false,
  "allow_pr": true,
  "allow_push": true,
  "allow_deploys": false,
  "allow_tags": false
}

```







Example Response Body:

```js
{
  "id": 1,
  "scm": "git",
  "owner": "octocat",
  "name": "hello-world",
  "full_name": "octocat/hello-world",
  "avatar_url": "https://avatars.githubusercontent.com/u/2181346?v=3",
  "link_url": "https://github.com/octocat/hello-world",
  "clone_url": "https://github.com/octocat/hello-world.git",
  "default_branch": "master",
  "timeout": 60,
  "private": false,
  "trusted": false,
  "allow_pr": true,
  "allow_push": true,
  "allow_deploys": false,
  "allow_tags": false
}

```









## <span class="operation operation-DELETE">DELETE</span> <span>Delete a repo</span>

Permanently deletes a repository. It cannot be undone.

```
DELETE /api/repos/{owner}/{name}
```


















## <span class="operation operation-POST">POST</span> <span>Encrypt repo secrets</span>

Encryptes a Yaml file with secret environment variables for secure public storage.

```
POST /api/repos/{owner}/{name}/encrypt
```

















# Builds


## <span class="operation operation-GET">GET</span> <span>Get recent builds</span>

Returns recent builds for the repository based on name.

```
GET /api/repos/{owner}/{name}/builds
```










Example Response Body:

```js
[{
  "id": 1,
  "number": 1,
  "event": "push",
  "status": "success",
  "created_at": 1443677151,
  "enqueued_at": 1443677151,
  "started_at": 1443677151,
  "finished_at": 1443677255,
  "commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
  "branch": "master",
  "ref": "refs/heads/master",
  "remote": "https://github.com/octocat/hello-world.git",
  "message": "New line at end of file. --Signed off by Spaceghost",
  "timestamp": 1443677255,
  "author": "Spaceghost",
  "author_avatar": "https://avatars0.githubusercontent.com/u/251370?v=3",
  "author_email": "octocat@github.com",
  "link_url": "https://github.com/octocat/hello-world/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
  "jobs": [
    {
      "id": 1,
      "number": 1,
      "status": "success",
      "enqueued_at": 1443677151,
      "started_at": 1443677151,
      "finished_at": 1443677255,
      "exit_code": 0,
      "environment": { "GO_VERSION": "1.4" }
    },
    {
      "id": 2,
      "number": 2,
      "status": "success",
      "enqueued_at": 1443677151,
      "started_at": 1443677151,
      "finished_at": 1443677255,
      "exit_code": 0,
      "environment": { "GO_VERSION": "1.5" }
    }
  ]
}
]
```







## <span class="operation operation-GET">GET</span> <span>Get the latest build</span>

Returns the latest repository build.

```
GET /api/repos/{owner}/{name}/builds/{number}
```














Example Response Body:

```js
{
  "id": 1,
  "number": 1,
  "event": "push",
  "status": "success",
  "created_at": 1443677151,
  "enqueued_at": 1443677151,
  "started_at": 1443677151,
  "finished_at": 1443677255,
  "commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
  "branch": "master",
  "ref": "refs/heads/master",
  "remote": "https://github.com/octocat/hello-world.git",
  "message": "New line at end of file. --Signed off by Spaceghost",
  "timestamp": 1443677255,
  "author": "Spaceghost",
  "author_avatar": "https://avatars0.githubusercontent.com/u/251370?v=3",
  "author_email": "octocat@github.com",
  "link_url": "https://github.com/octocat/hello-world/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
  "jobs": [
    {
      "id": 1,
      "number": 1,
      "status": "success",
      "enqueued_at": 1443677151,
      "started_at": 1443677151,
      "finished_at": 1443677255,
      "exit_code": 0,
      "environment": { "GO_VERSION": "1.4" }
    },
    {
      "id": 2,
      "number": 2,
      "status": "success",
      "enqueued_at": 1443677151,
      "started_at": 1443677151,
      "finished_at": 1443677255,
      "exit_code": 0,
      "environment": { "GO_VERSION": "1.5" }
    }
  ]
}

```







## <span class="operation operation-POST">POST</span> <span>Restart a build</span>

Restart the a build by number.

```
POST /api/repos/{owner}/{name}/builds/{number}
```












Example Response Body:

```js
{
  "id": 1,
  "number": 1,
  "event": "push",
  "status": "success",
  "created_at": 1443677151,
  "enqueued_at": 1443677151,
  "started_at": 1443677151,
  "finished_at": 1443677255,
  "commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
  "branch": "master",
  "ref": "refs/heads/master",
  "remote": "https://github.com/octocat/hello-world.git",
  "message": "New line at end of file. --Signed off by Spaceghost",
  "timestamp": 1443677255,
  "author": "Spaceghost",
  "author_avatar": "https://avatars0.githubusercontent.com/u/251370?v=3",
  "author_email": "octocat@github.com",
  "link_url": "https://github.com/octocat/hello-world/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
  "jobs": [
    {
      "id": 1,
      "number": 1,
      "status": "success",
      "enqueued_at": 1443677151,
      "started_at": 1443677151,
      "finished_at": 1443677255,
      "exit_code": 0,
      "environment": { "GO_VERSION": "1.4" }
    },
    {
      "id": 2,
      "number": 2,
      "status": "success",
      "enqueued_at": 1443677151,
      "started_at": 1443677151,
      "finished_at": 1443677255,
      "exit_code": 0,
      "environment": { "GO_VERSION": "1.5" }
    }
  ]
}

```









## <span class="operation operation-GET">GET</span> <span>Get build logs</span>

Returns the build logs for a specific job (build step).

```
GET /api/repos/{owner}/{name}/logs/{number}/{job}
```


















## <span class="operation operation-DELETE">DELETE</span> <span>Cancel a Job</span>

Cancel the a build job by number.

```
DELETE /api/repos/{owner}/{name}/logs/{number}/{job}
```





















# User


## <span class="operation operation-GET">GET</span> <span>Gets a user</span>

Returns the currently authenticated user.

```
GET /api/user
```






Example Response Body:

```js
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}

```





## <span class="operation operation-PATCH">PATCH</span> <span>Updates a user</span>

Updates the currently authenticated user.

```
PATCH /api/user
```








Example Response Body:

```js
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}

```







## <span class="operation operation-GET">GET</span> <span>Get user repos</span>

Retrieve the currently authenticated User's Repository list


```
GET /api/user/repos
```






Example Response Body:

```js
[{
  "id": 1,
  "scm": "git",
  "owner": "octocat",
  "name": "hello-world",
  "full_name": "octocat/hello-world",
  "avatar_url": "https://avatars.githubusercontent.com/u/2181346?v=3",
  "link_url": "https://github.com/octocat/hello-world",
  "clone_url": "https://github.com/octocat/hello-world.git",
  "default_branch": "master",
  "timeout": 60,
  "private": false,
  "trusted": false,
  "allow_pr": true,
  "allow_push": true,
  "allow_deploys": false,
  "allow_tags": false
}
]
```








# Users


## <span class="operation operation-GET">GET</span> <span>Get all users</span>

Returns all registered, active users in the system.

```
GET /api/users
```






Example Response Body:

```js
[{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}
]
```





## <span class="operation operation-GET">GET</span> <span>Get a user</span>

Returns a user with the specified login name.

```
GET /api/users/{login}
```








Example Response Body:

```js
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}

```







## <span class="operation operation-POST">POST</span> <span>Create a user</span>

Creates a new user account with the specified external login.

```
POST /api/users/{login}
```








Example Response Body:

```js
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}

```







## <span class="operation operation-PATCH">PATCH</span> <span>Update a user</span>

Updates an existing user account.

```
PATCH /api/users/{login}
```






Example Request Body:

```js
{
  "email": "octocat@github.com",
  "admin": false,
  "active": true
}

```







Example Response Body:

```js
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}

```







## <span class="operation operation-DELETE">DELETE</span> <span>Delete a user</span>

Deletes the user with the specified login name.

```
DELETE /api/users/{login}
```
