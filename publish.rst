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