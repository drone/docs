+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Builds"
menu = "api"
weight = 5
toc = true
+++


# Get recent builds

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
		
	

	



# Get the latest build

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
		
	

	



# Restart a build

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
		
	

	

	



# Get build logs

Returns the build logs for a specific job (build step).

```
GET /api/repos/{owner}/{name}/logs/{number}/{job}
```


	

	

	

	



	

	



# Cancel a Job

Cancel the a build job by number.

```
DELETE /api/repos/{owner}/{name}/logs/{number}/{job}
```


	

	

	

	



	

	

	



