Services
========

Drone provides a facility for connecting an application to one or more service
containers. A *service* is any Docker container that provides facilities that
your primary container needs to use. The most common examples are databases
(like ``mysql``, ``postgres``, ``mongo``), log services (``log4j``), and 
search servers (``elasticsearch``). These and many others are provided out of
the box. But just as you can use any image as a basis for your Drone build,
you can also use any image as a ``service``.

This section explains how to connect a service to your build using the
``.drone.yml`` file.

Available Services
------------------

Like build images, Drone comes with many built-in service images. And these,
too, have short alias names.

The following services are built-in:

- cassandra (relateiq/cassandra)
- couchdb (bradrydzewski/couchdb:1.5)
- elasticsearch:0.90 (bradrydzewski/elasticsearch:0.90)
- neo4j (bradrydzewski/neo4j:1.9)
- memcached (bradrydzewski/memcached)
- mongodb (bradrydzewski/mongodb:2.4)
- mysql (bradrydzewski/mysql:5.5)
- postgres (bradrydzewski/postgres:9.1)
- rabbitmq (bradrydzewski/rabbitmq:3.2)
- redis (bradrydzewski/redis:2.8)
- riak (guillermo/riak)
- zookeeper (jplock/zookeeper:3.4.5)

**NOTE:** Services are added regularly, and many services have more than one
version available. If what you are looking for is not in the list above,
you may want to check the GitHub project at https://github.com/drone/drone.

Example: Using Postgres
-----------------------

Here is an example that uses the Postgres database as a Drone service. It uses
an existing tool (Goose) to connect to the database and run a SQL migration.

.. code-block:: yaml

    image: go1.2

    services:
      - postgres

    script:

      # Make sure we have the client installed.
      - sudo apt-get -y install postgresql-client

      # Create a postgres database
      - createdb -h localhost -U postgres drone

      # Get and install the Goose DB tool
      - go get bitbucket.org/liamstask/goose/cmd/goose

      # Have goose set up our DB. This will run all of
      # the migrations in db/migrations
      - $GOPATH/bin/goose --env=drone up


The example above does the following:

- Uses the ``go1.2`` image. Goose is built in Go.
- Declares one service: ``postgres``, which will create a new Postgres container
  and map it to our build container.
- Creates a new database called ``drone`` using the Postgres command ``createdb``.
- Installs Goose, a simple database migration tool. We will use Goose as an
  example. It is not, though, required.
- Runs ``goose up`` to connect to the postgres service and execute some SQL.

Goose is a database migration tool. Its configuration file points it to a
database, and it connects according to the rules therein. Here's an excerpt from
Goose's ``db/dbconf.yml`` file:

.. code-block:: yaml

    drone:
        driver: postgres
        open: user=postgres host=localhost dbname=drone sslmode=disable

The above tells Goose to connect to a Postgres database using:

- The user ``postgres`` (Postgres' default account)
- The server is listening on ``localhost``.- The database ``drone``, which we 
  created in the ``createdb`` clause above.
- SSL disabled.

Note that there is no password or port specified. Drone does not override
the defaults.

Why point to localhost when Postgres is running in a different container? This
is a feature of Drone many of Drone's built-in images. They are configured to 
proxy connections from the local host to the correct services containers.

You can adapt the information above for your own application configuration
and connect your builds to the Drone service.

Tips for Using Services
-----------------------

For the most part, the services that Drone provides are as close to their
default configuration as possible. For example, above we saw that the ``postgres``
service uses the default user name (``postgres``) and no password, and attaches
to the default port.

If you would like to inspect a service, you may start it up as a plain Docker
container (See the Docker example here: http://docs.docker.io/examples/postgresql_service/).
Remember that you can create your own services.

Remember also that data is not persisted from one build to another. In our
example above, each time the build runs, we will have to create the new database.

More Information
----------------

The Drone.io database guide (http://docs.drone.io/databases.html) contains a
list of services and their configurations.
