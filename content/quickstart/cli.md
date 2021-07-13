---
date: 2000-01-01T00:00:00+00:00
title: Command Line Runner
title_in_header: Command Line Runner
author: bradrydzewski
weight: 500
toc: true

hidden: true
icon: /icons/shell.svg
description: |
  Execute and test pipelines locally using our command line tools
---

You can use the [command line tools]({{< relref "/cli/install" >}}) to run Docker pipelines locally, on your laptop or personal computer, for development and debugging purposes. This is our answer to the common complaint that _it works fine on my computer_.

# Installation

* Install the command line tools on OSX using [Homebrew](https://brew.sh/):
  ```
  $ brew install drone-cli
  ```

* Install the command line tools on Windows using [Scoop](https://scoop.sh/):
  ```
  $ scoop install drone
  ```

# Usage

The command line runner requires a working Docker installation. Drone executes your pipeline locally on your host machine using your local Docker daemon. _Local execution does not have any communication with the Drone server._

1. Navigate to the root directory of your git repository where your `.drone.yml` file is located. _Here is a basic configuration you can use for testing purposes:_
   ```
   kind: pipeline
   type: docker
   name: default
   
   steps:
   - name: test
     image: alpine
     commands:
     - echo hello
     - echo world
   ```

2. Execute your pipeline from the command line:
   ```
   $ drone exec
   ```

3. The command streams the pipeline logs to your terminal for analysis. The command returns a non-zero exit code if the pipeline fails.
   ```
   $ drone exec
   [test:1] + echo hello
   [test:2] hello
   [test:3] + echo world
   [test:4] world
   ```



<!-- * Example Go configuration:
   ```
   kind: pipeline
   type: docker
   name: default
   
   steps:
   - name: test
     image: golang
     commands:
     - go test -v
   ```

* Example Node configuration:

   ```
   kind: pipeline
   type: docker
   name: default
   
   steps:
   - name: test
     image: node
     commands:
     - npm install
     - npm test
   ``` -->

# Workspace

The command line runner mounts your current working directory, using docker volumes, as the working directory of your pipeline steps. The exposes your code to your pipeline. _The clone step is therefore skipped when running locally._

# Run Specific Pipelines

The command line runner runs the _default_ pipeline. If you use a different name for you pipeline, or you define multiple pipelines in your yaml, you can execute a named pipeline using the `--pipeline` flag.

* Example configuration with a pipeline named _test_
  ```
  kind: pipeline
  type: docker
  name: test

  steps:
  - name: dist
    image: node
    commands:
    - npm install
    - npm run dist
  ```

* Example command executes the _test_ pipeline

  ```
  drone exec --pipeline=test
  ```

# Run Specific Steps

When running pipelines locally you can limit which pipeline steps are executed and which pipeline steps are skipped.

* Execute only pipeline steps _build_ and _test_
   ```
   drone exec --include=build --include=test
   ```

* Execute all pipeline steps except _deploy_
   ```
   drone exec --exclude=deploy
   ```

# Emulate Secrets

The command line runner does not have read access to secrets stored in your server. You can provide secrets to your local pipeline by passing secrets to the command line runner via text file.

1. Example pipeline that requires username and password environment variables, sourced from secrets.
   ```
   kind: pipeline
   type: docker
   name: default

   steps:
   - name: test
     environment:
      USERNAME:
        from_secret: USERNAME
      PASSWORD:
        from_secret: PASSWORD
      SSH_KEY:
        from_secret: ssh_key
      CLOUD_AUTH:
        from_secret: cloud_auth_json
   ```
1. Create a simple text file with secrets defined one per line in key value format. _For the purposes of this demo we name the file `.env-drone`._
   ```
   # The key casing (upper/lower) must match the drone secret case
   #
   # Normal key value pairs are unquoted
   USERNAME=root
   PASSWORD=password
   # Quoting the value will cause the reader to parse escape codes like \n and therefore give you a multiline value.
   ssh_key="-----BEGIN RSA PRIVATE KEY-----\nREDACTED\n-----END RSA PRIVATE KEY-----"
   # Unquoted strings are not parsed, so we make this a single line of JSON, to preserve it exactly as it is, with all the escape codes in the 'private_key' field.
   cloud_auth_json={"type": "READCATED","private_key": "-----BEGIN PRIVATE KEY-----\nREDACTED\n-----END PRIVATE KEY-----\n","client_email": "REDACTED","client_id": "REDACTED"}
   ```

2. Provide your secrets file via command line flags when executing your pipeline.
   ```
   $ drone exec --secret-file=.env-drone
   ```

_The command line runner uses the [dotenv](https://github.com/joho/godotenv) package to read and parse the secrets file. If you are having problems with the secrets file please consult the official package [documentation](https://github.com/joho/godotenv)._

# Emulate Metadata

The command line runner does not communicate with the Drone server and therefore do not have access to repository, commit and build [metadata]({{< relref "/pipeline/environment/reference" >}}) that are otherwise available to your pipeline.

You can emulate repository metadata by passing repository and build information to the command line runner using command line flags and environment variables.

* Example command sets the branch:
   ```
   $ drone exec --branch=master
   ```

* Example command sets the build event:
   ```
   $ drone exec --event=pull_request
   ```

* Example command set metadata using environment variables:
  ```
  $ DRONE_SYSTEM_HOST=drone.company.com drone exec
  ```

# Trusted Mode

If your pipeline uses configurations that require trusted mode, you can enable trusted mode locally using the `--trusted` flag. _Trusted mode grants a pipeline elevated privileges to your host machine. Please use with caution._

* Example command enables trusted mode:

  ```
  $ drone exec --trusted
  ```

# Local vs Remote

The command line runner makes reasonable effort to emulate the server environment, however, there are notable differences that can impact execution:

* The server clones your repository, while the command line runner mounts your current working directory.  Any files or dependencies that already exist in your working directory are mounted into your pipeline as well.

* The server has access to more data (commit details, repository details, etc). This data needs to be provided to the command line runner to more closely emulate the server environment.

* The server has access to your oauth credentials and uses these credentials to generate a netrc file. These credentials need to be provided to the command line runner to more closely emulate the server environment. 

* The server has access to repository, organization and encrypted secrets. The command line runner does not have access to secrets or decryption keys stored in the server environment.
