---
date: 2017-04-15T14:39:04+02:00
title: "Secrets Not Working"
url: secrets-not-working
---

The purpose of this guide is to help troubleshoot common root causes for secrets not working as expected. Before you proceed please make sure you have read the official secret documentation and double check your configuration.

The most common root cause is forgetting to include the `secrets` block in your pipeline step configuration.

```diff
pipeline:
  publish:
    image: plugins/docker
    repo: octocat/hello-world
+   secrets: [ docker_username , docker_password ]
```

The second most common root cause is using incorrect variable names. Plugins look for specific variable names. It is therefore required that you provide secrets to the plugin using the expected name.

```diff
pipeline:
  publish:
    image: plugins/docker
    repo: octocat/hello-world
-   secrets: [ username , password ]
+   secrets: [ docker_username , docker_password ]
```

Please consult the respective plugin documentation, or contact the plugin author, if you are unsure which variable names to use.

# Variable Expansion

Please note that variable expansion of secrets is not supported. The following yaml configuration __will not work__.

```diff
pipeline:
  publish:
    image: plugins/docker
    repo: octocat/hello-world
    build_args:
-     - npm_username=${npm_username}
-     - npm_password=${npm_password}
    secrets: [ docker_username , docker_password, npm_username, npm_password ]
```

# Variables with Newlines

If your secret includes newlines or special characters we recommend creating these from a file. This ensures newlines and special characters are preserved.

```diff
drone secret add \
  -repository octocat/hello-world \
  -name ssh_key \
- -value $(cat /root/ssh/id_rsa)
+ -value @/root/ssh/id_rsa
```

When writing secrets to file using `echo` please ensure you are using the correct shell syntax and quoting the variable to ensure newlines and special characters are preserved.


```diff
pipeline:
  build:
    image: golang
     commands:
-      - echo -n $SSH_KEY > /root/.ssh/id_rsa
+      - echo -n "$SSH_KEY" > /root/.ssh/id_rsa
```

# Getting Help

If you continue to experience issues you can engage the community for support. Please include the  following information in your support request:

* the full yaml configuration file
* the relevant logs from your build
* the results of `drone secrets ls <repo>`
* the results of `drone repo info <repo>`
