+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Repos"
menu = "api"
weight = 5
toc = true
+++


# Get a repo

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
		
	

	



# Activates a repo

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
		
	

	

	

	

	

	



# Updates a repo

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
		
	

	

	



# Delete a repo

Permanently deletes a repository. It cannot be undone.

```
DELETE /api/repos/{owner}/{name}
```


	

	



	

	

	

	



# Encrypt repo secrets

Encryptes a YAML file with secret environment variables for secure public storage.

```
POST /api/repos/{owner}/{name}/encrypt
```


	

	



	

	

	



