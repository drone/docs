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

Here is an example that uses the Postgres database as a Drone service.

.. code-block:: yaml

    image: go1.2
    services:
      - postgres
    script:
      - createdb -h localhost -U postgres drone
      - psql -U postgres -h localhost -c "SELECT 1" drone



The example above does the following:

- Uses the ``go1.2`` image. Goose is built in Go.
- Declares one service: ``postgres``, which will create a new Postgres container
  and map it to our build container.
- Creates a new database called ``drone`` using the Postgres command ``createdb``.
- Runs a very simple ``SELECT`` statement.


Importantly, from the examples above, we can see the primary
information that we can use to connect application code to Postgres:

- Default user: ``postgres``
- Default password: none
- Host: ``localhost``
- Port: ``5432``
- SSL: disabled (Some clients default to enabling SSL).

Why point to localhost when Postgres is running in a different container? This
is a feature of many of Drone's built-in images. They are configured to 
proxy connections from the local host to the correct services containers.

Once a project like this is pushed, the output should look something like this"

.. code-block:: console

    $ git clone --depth=50 --recursive --branch=master https://bitbucket.org/technosophos/drone-postgres-example.git /var/cache/drone/src/bitbucket.org/technosophos/drone-postgres-example
    Cloning into '/var/cache/drone/src/bitbucket.org/technosophos/drone-postgres-example'...
    $ git checkout -qf c7e9917648ebcd3f5f5e622bc3502f41caf7cd64
    $ createdb -h localhost -U postgres drone
    $ psql -U postgres -h localhost -c "SELECT 1" drone
     ?column? 
    ----------
            1
    (1 row)

    $ exit 0


Note that in the last few lines we can see the output of the query from Postgres.

Using the same connection information presented above, you can configure your
application to connect to the service instance and use it during the Drone run.

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
