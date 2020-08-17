---
date: 2000-01-01T00:00:00+00:00
title: drone encrypt
author: bradrydzewski
weight: 4
separator: true
---

The command encrypts a secret. The resulting string can be added to your pipeline. See [Encrypted Secrets]({{< relref "/secret/encrypted.md" >}}) for complete instructions.

```
NAME:
   drone encrypt - encrypt a secret

USAGE:
   drone encrypt [command options] <repo/name> <string>
```

Example usage:

```
drone encrypt octocat/hello-world "correct-horse-batter-staple"
```

Example usage, reads the secret from a file:

```
drone encrypt octocat/hello-world @/path/to/secret.txt
```