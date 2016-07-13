+++
date = "2016-07-12T16:00:21-08:00"
title = "Drone Architecture"
weight = 1

[menu.main]
	parent="reference"
+++

Drone is comprised of two primary components: Drone Server and Drone Agent. These can be ran on the same machine, or distributed out over a number of different hosts.

# Drone Server

Each Drone setup has one Drone Server. Duties handled include:

* Listening for build webhooks
* Listening for Drone Agent websocket connections
* Hosting the Drone HTTP API
* Serving the Drone web UI
* Tracking and assigning builds to Agents
* Storing and managing build state and build logs

Drone Server is our central source of truth, and is the only component that communicates directly with whatever database you have configured.

Drone Server's pool of available workers is determined by which Drone Agents have active, healthy connections established. Agents can come and go at will, with builds being re-ran after a period of inactivity. Another interesting characteristic of this setup is that the pool Drone Agents can be auto-scaled without any changes to your Drone Server's configuration.

# Drone Agent

Drone Agents are very simple processes that connect to Drone Server via websockets to request builds to run. The Agents speak a standard protocol that will be documented and stabilized in the future, allowing end users to provide their own implementations.
 
When a build is pulled, the Agent communicates with the neighboring local Docker daemon to run the build container, services, and plugins. The vast majority of the work happens within these containers. The Agent simple manages the source volume, the containers, and pipes the stdout/stderr to the master.
