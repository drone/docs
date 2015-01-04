Deployment
==========

Drone can automate your deployment steps at the end of each successful build. Drone includes a
number of deployment plugins that can be configured in the `deploy` section of the `.drone.yml` file.

Deployments are **skipped** for:

* Failed builds
* Pull requests

Cloud Foundry
-------------

Deploy to `Cloud Foundry <http://www.cloudfoundry.org/index.html>`_ hosted service over `cf tool <https://github.com/cloudfoundry/cli/releases>`_ verion 6. it assumes your provide the application manifest file `manifest.yml` under project root directory if no application name provided.

.. code-block:: console

    deploy:
      cloudfoundry:
        target: https://api.run.pivotal.io
        username: my-cloudfoundry-username
        password: my-password
        organization: my-cloudfoundry-org
        space: development
        app: my-cloudfoundry-app

target
  URL of the Cloud Controller in Cloud Foundry

username
  your username

password
  your password

organization
  organization name where you want to deploy your application (optional)

space
  space name in the organization where you want to deploy your application (optional)

app
  name of the application to push (optional)

Engine Yard
------------

We are looking for volunteers to add this plugin.

Git
---

Deploy code via ``git push`` to a remote server.

.. code-block:: console

    deploy:
        git:
            target: git@foo.com/bar.git
            branch: master
            force: false

target
  name of the git remote server to push to

branch
  destination branch, optional, defaults to master

force
  true | false to use the git --force flag


Bash
----

Deploy code via any bash command, for example with [Capistrano](https://github.com/capistrano/capistrano) or [Fabric](https://github.com/fabric/fabric)

.. code-block:: console

    deploy:
        bash:
            command: bundle exec cap production deploy
            
or

.. code-block:: console

    deploy:
        bash:
            script: 
              - ./bin/prepare_for_deploy.sh
              - ./bin/make_deploy.sh
              - ./bin/finish_deploy.sh

Stage Environment Deploy

.. code-block:: console

    deploy:
        bash:
            script:
                - [ $DRONE_BRANCH == 'master'  ] && bundle install --path vendor/bundle && bundle exec cap production deploy
                - [ $DRONE_BRANCH == 'develop' ] && bundle install --path vendor/bundle && bundle exec cap staging deploy

command
  bash command that runs deploy

script
  array of bash commands that run deploy


Heroku
------

Deploy to the `Heroku <https://www.heroku.com>`_ hosting service.

.. code-block:: console

    deploy:
        heroku:
            app: my-heroku-app
            force: false


app
  name of your heroku application

force
  true | false to use the git --force flag

Modulus
-------

Deploy to the `modulus.io <https://modulus.io>`_ hosting service.

.. code-block:: console

    deploy:
        modulus:
            project: my-modulus-app
            token: 5f05189c

project
  name of your modulus project

token
  your modulus api token


Nodejitsu
---------

Deploy to the `nodejitsu <https://www.nodejitsu.com>`_ hosting service.

.. code-block:: console

    deploy:
        nodejitsu:
            user: my-nodejitsure-username
            token: 5f05189c

user
  your nodejitsu username

token
  your nodejitsu api token


Open Shift
----------

We are looking for volunteers to add this plugin.

Tsuru
------

Deploy to the `Tsuru <http://www.tsuru.io>`_ hosting service.

.. code-block:: console

    deploy:
       tsuru: 
            force: false
            remote: git@git.tsuru.io:my-tsuruapp.git

force
  true | false to use the git --force flag (default: `false`).

remote
  git remote of your tsuru application
