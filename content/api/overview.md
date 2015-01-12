+++
weight = -3
date = "2014-11-08T12:54:12-08:00"
title = "Overview"

[menu.api]
parent = "API Documentation"
+++

Welcome to the Drone API documentation. This is the same API used by the Drone web interface (Angular),
so everything the web ui is able to do can also be accomplished via the API.

## Errors

Drone uses conventional HTTP response codes to indicate success or failure of an API request.
In general, codes in the 2xx range indicate success, codes in the 4xx range indicate an error
that resulted from the provided information (e.g. a required parameter was missing),
and codes in the 5xx range indicate an unexpected error from the server.

## CORS

The API supports Cross Origin Resource Sharing (CORS) for AJAX requests so that you can
interact with the API from client-side web applications.

## Clients

Here is a list of available API client libraries for working with the Drone API: 

* [Go](https://github.com/drone/drone/tree/master/client) (official)