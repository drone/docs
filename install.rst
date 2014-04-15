Installation
============

Requirements
------------

Drone is tested on the following versions of Ubuntu:

* Ubuntu Precise 12.04 (64-bit)
* Ubuntu Raring 13.04 (64 bit)

Drone requires the latest version of Docker (0.8)

Installation
------------

Drone is distributed as a debian package for easy installation:

.. code-block:: console

    $ wget http://downloads.drone.io/latest/drone.deb
    $ sudo dpkg -i drone.deb

This will install the following files:

* Drone server `/usr/local/bin/droned`
* Drone client `/usr/local/bin/drone`
* Drone startup script `/etc/init/drone.conf`
* Drone sqlite database `/var/lib/drone/drone.sqlite`

We recommend running Drone on a 2GB `Digital Ocean Docker
image <https://www.digitalocean.com/community/articles/how-to-use-the-digitalocean-docker-application>`_.
This is the fastest, easiest way to get up and running with Drone. This is also how we test Drone internally.

Running
-------

Drone is started automatiaclly. You can start and stop Drone manually with
the following commands:

.. code-block:: console

    $ sudo start drone
    $ sudo stop drone

From Source
-----------

The project is hosted at https://github.com/drone/drone and can be installed
manually. You will need the latest version of Go (1.2) and the following
software packages:

.. code-block:: console

    $ sudo apt-get install make mercurial git bzr libsqlite3-dev sqlite3

.. code-block:: console

    $ git clone git://github.com/drone/drone.git
    $ cd drone
    $ make deps
    $ make

Proxy Server
------------

**NOTE:** using a proxy server is not really recommended. Drone serves most static content from a CDN
and uses the Go standard library's high-performance ``net/http`` package to serve dynamic content.

If using Nginx to proxy traffic to Drone, please ensure you have version 1.3.13
or greater. You also need to configure nginx to proxy websocket connections:

.. code-block:: bash

    # Proxy for websockets
    location = /feed {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host $http_host;

        proxy_pass http://127.0.0.1:8080;
        proxy_redirect off;

        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }

You will also want to change the default port. You can pass the port as a command line argument ``--port=:8080``
or you can change the default port in the ``/etc/default/drone`` file:

.. code-block:: bash

    # Upstart configuration file for droned.
    
    # Command line options:
    #
    #   -datasource="drone.sqlite":
    #   -driver="sqlite3":
    #   -path="":
    #   -port=":8080":
    #         
    DRONED_OPTS="--port=:8080"
