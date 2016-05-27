+++
date = "2015-12-05T16:00:21-08:00"
title = "drone sign"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++

# Overview

Drone requires you to sign the Yaml file and will verify the signature before injecting secrets into your build environment. The sign subcommand is used to generate the signature.

# Signature

Generate the signature for the named repository. This will read the `.drone.yml` file in the working directory and write the signature to a `.drone.yml.sig` file in the working directory:

```
$ drone sign octocat/hello-world
```

Don't forget to commit the signature file to your repository:

```
$ git add .drone.yml.sig
$ git commit -m "signed yaml file"
$ git push
```
