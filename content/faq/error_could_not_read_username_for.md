---
date: 2017-04-15T14:39:04+02:00
title: "Fatal: could not read Username for"
url: error-could-not-read-username-for
---

This error message happens when attempting to clone a repository that requires authentication, but authentication credentials are not available to the pipeline.

```nohighlight
fatal: could not read Username for 'https://...': No such device or address
exit status 128
```

Note that when properly configured, you should never need to provide drone with credentials to clone your repository. This is something Drone handles automatically.

# Private Mode

The most common root cause for this error message is when your version control server requires authentication to clone __public__ repositories, usually known as private mode. You will need to configure Drone accordingly.

```nohighlight
DRONE_GITHUB_PRIVATE_MODE=true
DRONE_GITLAB_PRIVATE_MODE=true
DRONE_GOGS_PRIVATE_MODE=true
```

Note that after making this change you will need to run the below command. This will instruct Drone to treat the repository as if it were private and to require authentication.

```nohighlight
drone repo repair <owner/repo>
```

# Public to Private

The second most common root cause is when you change your repository from public to private. Note that after making this change you will need to run the below command. This will instruct Drone to treat the repository as if it were private and to require authentication.

```nohighlight
drone repo repair <owner/repo>
```
