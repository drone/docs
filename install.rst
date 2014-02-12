Installation
============

Requirements
------------

Drone is tested on the following versions of Ubuntu:

* Ubuntu Precise 12.04 (LTS) (64-bit)
* Ubuntu Raring 13.04 (64 bit)

Drone requires the latest version of Docker (0.8)

Installation
------------

Drone is distributed a as debian package for easy installation:

.. code-block:: console

    $ wget http://downloads.drone.io/latest/drone.deb
    $ sudo dpkg -i drone.deb
    $ sudo start drone

We recommend running Drone on a 2GB `Digital Ocean Docker
image <https://www.digitalocean.com/community/articles/how-to-use-the-digitalocean-docker-application>`_.
This is the fastest, easiest way to get up and running with Drone. This is also how we test Drone internally.

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
