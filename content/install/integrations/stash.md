+++
date = "2017-04-15T14:39:04+02:00"
title = "Bitbucket Server"
url = "install-for-bitbucket-server"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-bitucket-server"
  weight = 6
+++

Drone comes with experimental support for Bitbucket Server, formerly known as Atlassian Stash. To enable Bitbucket Server you should configure the Drone container using the following environment variables:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_STASH=true
+     - DRONE_STASH_GIT_USERNAME=foo
+     - DRONE_STASH_GIT_PASSWORD=bar
+     - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
+     - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
    volumes:
+     - /path/to/key.pem:/path/to/key.pem

  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

# Private Key File

The OAuth process in Bitbucket server requires a private and a public RSA certificate. This is how you create the private RSA certificate.

```nohighlight
openssl genrsa -out /etc/bitbucket/key.pem 1024
```

This stores the private RSA certificate in `key.pem`. The next command generates the public RSA certificate and stores it in `key.pub`.

```nohighlight
openssl rsa -in /etc/bitbucket/key.pem -pubout >> /etc/bitbucket/key.pub
```

Please note that the private key file can be mounted into your Drone conatiner at runtime or as an environment variable

Private key file mounted into your Drone container at runtime as a volume.

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
    - DRONE_OPEN=true
    - DRONE_HOST=${DRONE_HOST}
      - DRONE_STASH=true
      - DRONE_STASH_GIT_USERNAME=foo
      - DRONE_STASH_GIT_PASSWORD=bar
      - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
      - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
+  volumes:
+     - /etc/bitbucket/key.pem:/etc/bitbucket/key.pem
```

Private key as environment variable

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
    - DRONE_OPEN=true
    - DRONE_HOST=${DRONE_HOST}
      - DRONE_STASH=true
      - DRONE_STASH_GIT_USERNAME=foo
      - DRONE_STASH_GIT_PASSWORD=bar
      - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA_STRING=contentOfPemKeyAsString
      - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
```

# Service Account

Drone users `git+https` to clone repositories, however, Bitbucket Server does not currently support cloning repositories with oauth token. To work around this limitation, you must create a service account and provide the username and password to Drone. This service account will be used to authenticate and clone private repositories.

# Registration

You must register your application with Bitbucket Server in order to generate a consumer key.

Navigate to your account settings (e.g. http://stash.mycompany.com/admin) and choose "Application Links" from the menu on the left.

Enter the URL of your Drone server (e.g. http://drone.mycompany.com) and click "Create new link". Bitbucket server will warn you that no response was received from the URL you entered. Ignore this and click "Continue".

Give the application a name (e.g. "Drone CI") and click "Continue".

Edit the application you just created. Click on "Incoming Authentication". In the "Consumer Key" field enter the value of your `DRONE_STASH_CONSUMER_KEY` environment variable. In the "Consumer Name" field enter a suitable name, e.g. "Drone CI" (the exact value doesn't really matter). In the "Public Key" field enter the public key that you generated above (either the value of `DRONE_STASH_CONSUMER_RSA_STRING` or the contents of the file `DRONE_STASH_CONSUMER_RSA`). In the "Consumer Callback URL" enter the authorize URL, e.g. `http://drone.mycompany.com/authorize`. Click "Close".


# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.


DRONE_STASH=true
: Set to true to enable the Bitbucket Server (Stash) driver.

DRONE_STASH_URL
: Bitbucket Server address.

DRONE_STASH_CONSUMER_KEY
: Bitbucket Server oauth1 consumer key

DRONE_STASH_CONSUMER_RSA
: Bitbucket Server oauth1 private key file

DRONE_STASH_CONSUMER_RSA_STRING
: Bibucket Server oauth1 private key as a string

DRONE_STASH_GIT_USERNAME
: Machine account username used to clone repositories.

DRONE_STASH_GIT_PASSWORD
: Machine account password used to clone repositories.
