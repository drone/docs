+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Bintray"
description = "Publish files and artifacts to Bintray"
user = "drone-plugins"
repo = "drone-bintray"
image = "plugins/drone-bintray"
tags = ["publish", "bintray", "artifacts"]
categories = "publish"
draft = false
date = 2016-02-13T09:01:27Z
menu = ""
weight = 1

+++

Use the Bintray plugin to publish one or more packages to a Bintray repository. You must supply the following parameters in a **bintray** subsection of the .drone.yml **publish** section.

Parameter | Description
--------- | -----------
username  | A Bintray username
api_key | The API key associated with the username (obtain from Bintray settings)
host | Optional host of the bintray api server with scheme. Example: https://bin.company.com
branch | The GitHub branch of the commit.  A non-empty value other than "master" causes the plugin to upload artifacts to test/branchname/target instead of just target
debug | If true, causes the plugin to write additional output to stdout, including its invocation arguments and output curl command.
insecure | Enable insecure connections
artifacts | One or more descriptions of artifacts to upload

The following parameters describe each artifact

Artifact Parameter | Description
--------- | -----------
file | Path to the file to upload within the build directory (the original build directory, not the current working directory)
owner | Bintray repository owner (called subject in Bintray API documentation)
type | Artifact type.  Specify "Debian" for a Debian artifact or "Maven" for a Maven artifact.  Any other value indicates a generic artifact.
repository | Name of the repository where the file should be uploaded
package | Name of the package within the repository
version | Version to upload.  Required for non-Maven artifacts and ignored for Maven.
target | Name to be given to the artifact in the repository.  The name is prefixed by "test/branchname" if the branch parameter is specified and has a non-empty value other than "master"
distr | Debian distribution.  Required for Debian artifacts and ignored for others.
component | Debian component.  Required for Debian artifacts and ignored for others.
arch | Debian architecture(s).  Required for Debian artifacts and ignored for others.
override | true if the artifact can override (replace) a previously-published artifact; false otherwise
publish | true if the artifact should be published after upload; false otherwise

## Example

The follow .drone.yml file uploads two artifacts to Bintray

```yaml
publish:
  bintray:
    username: $$BINTRAY_USERNAME
    password: $$BINTRAY_PASSWORD
    branch: $$BRANCH
    artifacts:
      - file: dist/myfile
        owner: mycompany
        type: executable
        repository: reponame
        package: pkgname
        version: 1.0
        target: myfile
        publish: true
        override: true
      - file: dist/myfile.deb
        owner: mycompany
        type: Debian
        repository: debian-repo
        package: pkgname
        version: 1.0
        target: myfile.deb
        distr: ubuntu
        component: main
        arch:
          - amd64
        publish: true
        override: true
```

