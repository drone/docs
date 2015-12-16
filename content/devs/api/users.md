+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Users"
menu = "api"
weight = 5
toc = true
+++


# Get all users

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
		
	



# Get a user

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
		
	

	



# Create a user

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
		
	

	



# Update a user

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
		
	

	



# Delete a user

Deletes the user with the specified login name.

```
DELETE /api/users/{login}
```


	



	

	

	

	



