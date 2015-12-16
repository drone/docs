+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "User"
menu = "api"
weight = 5
toc = true
+++


# Gets a user

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
		
	



# Updates a user

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
		
	

	



# Get user repos

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
		
	

	



