+++
weight = 17
date = "2014-11-08T12:54:06-08:00"
title = "Get the Commit Output"
draft = false

path = "/api/repos/:repo/branches/:branch/commits/:commit/console"
method = "GET"
resource = "commits"
+++

Gets the build output (stdout) for the commit matching the specified repository, branch and sha.

### Example Request: 

```bash
curl -X GET https://drone.io/api/repos/github.com/foo/bar/branches/master/commits/d5e5797/console
```

### Example Response (Plain Text, 200 OK):

```nohighlight
$ git clone github.com/drone/drone
$ make deps
$ make test
go vet ./...
go test -cover -short ./...
ok  	/plugin/condition	0.003s	coverage: 94.4% of statements
ok  	/plugin/deploy	0.004s	coverage: 85.1% of statements
?   	/plugin/deploy/cloudfoundry	[no test files]
ok  	/plugin/deploy/deis	0.003s	coverage: 90.0% of statements
ok  	/plugin/deploy/git	0.003s	coverage: 92.3% of statements
ok  	/plugin/deploy/heroku	0.004s	coverage: 90.9% of statements
ok  	/plugin/deploy/modulus	0.004s	coverage: 88.9% of statements
ok  	/plugin/deploy/nodejitsu	0.003s	coverage: 87.5% of statements
ok  	/plugin/deploy/tsuru	0.008s	coverage: 90.0% of statements
exit 0
```