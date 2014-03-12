Deployment
==========

Drone can automate your deployment steps at the end of each successful build. Drone includes a
number of deployment plugins that can be configured in the `deploy` section of the `.drone.yml` file.

Deployments are **skipped** for:

* Failed builds
* Pull requests

Cloudfoundry
------------

We are looking for volunteers to add this plugin.

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
            force: false

target
  name of the git remote server to push to

force
  true | false to use the git --force flag


Heroku
------

Deploy to the `Heroku <https://www.heroku.com>`_` hosting service.

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

Deploy to the `modulus.io <https://modulus.io>`_` hosting service.

.. code-block:: console

    deploy:
        modulus:
            project: my-modulus-app
            token: 5f05189c


Nodejitsu
---------

We are looking for volunteers to add this plugin.

Open Shift
----------

We are looking for volunteers to add this plugin.



