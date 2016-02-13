+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Artifactory"
description = "Publish files and artifacts to Artifactory"
user = "drone-plugins"
repo = "drone-artifactory"
image = "plugins/drone-artifactory"
tags = ["artifactory"]
categories = "publish"
draft = false
date = 2016-02-13T09:01:55Z
menu = ""
weight = 1

+++

Use this plugin to publish artifacts from the build to Artifactory.
You can override the default configuration with the following parameters:

* `url` - Artifactory URL (Includes scheme)
* `username` - Artifactory username, default to blank
* `password` - Artifactory password, default to blank
* `pom` - An optional pom.xml path were to read project details
* `group_id` - Project group id, default to value from Pom file
* `artifact_id` - Project artifact id, default to value from Pom file
* `version` - Artifact version, default to value from Pom file
* `repo_key` - Target repository key, default to 'libs-snapshot-local' if version contains 'snapshot', 'libs-release-local' otherwise
* `files` - List of files to deploy
* `force_upload` - Force upload if a file already exists

All file paths must be relative to current project sources

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  artifactory:
    url: http://arti.company.com
    username: admin
    password: password 
    pom: pom.xml 
    repo_key: libs-snapshot-local
    files: 
      - target/*.jar
      - target/*.war
```

## pom.xml deployment

If a pom parameter is specified it will be automatically deployed. It is not necessary to specify the pom under the files parameter. 

In the example above, pom.xml will be deployed as ```<groupId>-<artifactId>-<version>.pom```


