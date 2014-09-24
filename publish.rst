Publish
=======

Drone can automate publishing files at the end of each successful build. Drone includes a
number of publish plugins that can be configured in the `publish` section of the `.drone.yml` file.

Publishes are **skipped** for:

* Failed builds
* Pull requests

S3
--

Publish files to Amazon S3

.. code-block:: console

    publish:
        s3:
            acl: public-read
            region: us-east-1
            bucket: downloads.drone.io
            access_key: C24526974F365C3B
            secret_key: 2263c9751ed084a68df28fd2f658b127
            source: /tmp/drone.deb
            target: latest/

acl
  Access control to apply to file

region
  Region where the bucket resides that you are publishing to

bucket
  Bucket to publish the file to

access_key
  AWS Access key

secret_key
  AWS Secret Key

source
  Source file to publish

target
  Destiantion publish location

Swift
-----

Publish files to OpenStack Swift or Rackspace CloudFiles

.. code-block:: console

    publish:
        swift:
            username: someuser
            password: 030e39a1278a18828389b194b93211aa
            auth_url: https://identity.api.rackspacecloud.com/v2.0
            region: DFW
            container: drone
            source: /tmp/drone.deb
            target: latest/drone.deb

username
  OpenStack/Rackspace Username to authenticate with

password
  OpenStack Keystone password or Rackspace API Key

auth_url
  URL of the Keystone identitiy endpoint

region
  Region where the container resides that you are publishing to

container
  Container to publish the file to

source
  Source file to publish

target
  Destination publish location (only required when source is a single file)

PyPI
----

Publish a Python package to pypi.python.org

.. code-block:: console

    publish:
        pypi:
            username: someuser
            password: somepassword

username
  PyPI username to authenticate with

password
  PyPI password to authenticate with

NPM
---

Publish a Node.js package to npm registry

.. code-block:: console

    publish:
        npm:
            username: someuser
            email: someuser@example.com
            password: somepassword
            registry: https://somereg.example.com/
            folder: my-node-project/
            tag: latest

username
  npm registry username to authenticate with

email
  npm registry email to authenticate with

password
  npm registry password to authenticate with

registry
  npm registry URL. default to http://registry.npmjs.org/ (optional)

folder
  a folder containing a package.json file (optional)

tag
  registers the published package with the given tag (optional)

Docker
------

Publish a Docker image to a specified repo or registry. Supports the following configurations:

 * Private Docker Registry (unauthenticated)

 * Private Docker Registry (authenticated)
   e.g. login with `username` + push to `docker.example.com/image:tag`
 * Push to Docker Hub user ID 
   e.g. `username/image:tag`
 * Push to Docker Hub company or group 
   e.g. login with `username` but push to `company/image:tag`

.. code-block:: console

    publish:
        docker:
            dockerfile: MyDockerFile
            docker_host: docker.example.com
            docker_port: 1000
            docker_version: 1.0
            registry_host: docker.example.com
            registry_protocol: https
            registry_login: true
            registry_login_uri: /registry/v1/
            username: myuser
            password: mypassword
            email: myuser@example.com
            image_name: my-webapp
            push_latest: true
            keep_builds: false
            tag: 0.1

dockerfile
  The Dockerfile you want to use to build your final image.
  **Default:** `./Dockerfile` in the root of your codebase.

docker_server
  IP address or hostname for the Docker server that you want to connect to for building/pushing your image.
  **Note:** This does not need to match the final destination/end-point for your image.

docker_port
  The TCP port on which the Docker daemon is listening for remote connections (configured by adding `-H tcp://{IP_ADDRESS}:{PORT}` to `DOCKER_OPTS` in `/etc/default/docker` on the server you want to use to build images).

docker_version
  The version of Docker Engine that is running on the remote Docker server (not the registry).

registry_login_url (optional)
  The full login URI used to post authentication details (e.g. `https://docker.company.com/v1/`)

registry_login (optional)
  Does your custom registry endpoint require login? Defaults to `false`
  **Note:** This is not applicable when pushing to Docker Hub, it will always require authentication.

image_name
  The name you would like to give your image (excluding the image tag)

tag (optional)
  The tag you would like to set for this image build. Default is the short git commit ID `git rev-parse --short HEAD`

username (optional for private repositories)
  The username used to authenticate to the private registry or to Docker Hub

password (optional for private repositories)
  Your authentication password

email (optional for private repositories)
  Your email address

keep_build (optional)
  Set to `true` if you would like to leave the final image on the `docker_host` used to build it. Default is `false`, which cleans up the build after successfully pushing to the registry.

push_latest (optional)
  In addition to tagging with either `custom_tag` or the git-ref of your code, should we tag an image as `:latest` before pushing it? Default behaviour is set to `true`.

Example Configs
+++++++++++++++

**Private Registry, no authentication**

.. code-block:: console

    publish:
        docker:
            docker_host: docker.example.com
            docker_port: 1000
            docker_version: 1.0
            registry_login: false
            image_name: docker.example.com/my-webapp
            push_latest: false
            keep_builds: false
            custom_tag: 0.1

Result: Image pushed to `docker-registry.example.com/my-webapp:0.1` without login.

**Private Registry, Authenticated**

.. code-block:: console

    publish:
        docker:
            docker_host: docker.example.com
            docker_port: 1000
            docker_version: 1.0
            registry_login_url: https://docker-registry.example.com/v1/
            registry_login: true
            username: myuser
            password: mypassword
            email: myuser@example.com
            image_name: docker-registry.example.com/my-webapp
            push_latest: true
            keep_builds: false

Result: Image pushed to `docker-registry.example.com/my-webapp:$(git rev-parse --short HEAD)` using `myuser` account. `docker.example.com/my-app:latest` is also tagged.

**Docker Hub, Push to Personal Account**
.. code-block:: console

    publish:
        docker:
            docker_host: docker.example.com
            docker_port: 1000
            docker_version: 1.0
            username: myuser
            password: mypassword
            email: myuser@example.com
            image_name: my-webapp
            push_latest: true
            keep_builds: false

Result: Image pushed to Docker Hub as `myuser/my-webapp:$(git rev-parse --short HEAD)` using `myuser` account.

**Docker Hub, Push to Shared Repository**

.. code-block:: console

    publish:
        docker:
            docker_host: docker.example.com
            docker_port: 1000
            docker_version: 1.0
            username: myuser
            password: mypassword
            email: myuser@example.com
            image_name: mycompany/my-webapp
            push_latest: false
            keep_builds: false
            tag: 0.1

Result: Image pushed to Docker Hub as `mycompany/image:0.1` using `myuser` account.

GitHub
------

Create a `GitHub release <https://github.com/blog/1547-release-your-software>`_.

.. code-block:: console

    publish:
        github:
            branch: master
            script:
              - make embed
              - make release
            artifacts:
              - release/
            tag: v$(cat VERSION)
            token: {{github_token}}
            user: drone
            repo: drone

script
  Optional list of commands to run to prepare a release.

artifacts
  List of files or directories to release.

tag
  Required name of the tag to create for this release.
  If the tag already exists, the plugin will do nothing.

name
  The name of the release. Defaults to tag.

description
  Description of the release. Defaults to empty.

draft:
  true/false identifier allowed on GitHub releases. Defaults to false.

prerelease:
  true/false identifier allowed on GitHub releases. Defaults to false.

token:
  Required OAuth token to use when interacting with the GitHub API.
  The token must have "repo" access, and the user associated with the token must have read and write access to the repo.

user:
  The user or organization for the repository you'd like to publish to.

repo:
  The name of the respository to publish to.

branch:
  Restrict this plugin to only run on commits to this branch. Defaults to "master".
