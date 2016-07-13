+++
date = "2015-12-05T16:00:21-08:00"
title = "Drone 0.5"
weight = 500
toc = true

[menu.main]
	parent="releases"
+++

# What's new in Drone 0.5

## Completely re-architected worker system

In earlier versions of Drone, a fixed list of worker addresses was maintained within the database. When builds were triggered, Drone would choose a healthy worker and fire off the build by sending API requests against the worker's Docker daemon address. This required that the Docker Server be able to open a socket to the worker's Docker daemon, and was generally fragile. 

As of Drone 0.5, we have introduced the notion of a Drone Agent. This is a simple process that connects to the Drone Server over websockets, _pulling_ jobs off of the queue. This simplifies firewall rules and makes auto-scaling the Agent fleet as easy as firing up additional processes.

## Entirely new UI

The UI seen in 0.4 was a placeholder that served us well in getting the project moving. Rather than bandaid some of the issues, the whole thing was re-designed and re-written from the ground up.
 
It is now much easier to navigate repositories, see which repos have pending builds, and get a feel for the overall health of your projects. Within the detail views, we now show stdout and stderr for the various service containers you may be using.

Additionally, we've reworked log streaming to be more durable and efficient.

# Backwards incompatible changes in 0.5

## Revamped plugin inputs

Prior to Drone 0.5, arguments to plugin containers were passed into stdin via a JSON blob. This worked, but led to quite a bit of boilerplate (depending on the language).

In Drone 0.5, we have shifted towards passing in environment variables instead of JSON blobs + stdin. This should greatly simplify the parsing and handling of input for many languages.

Unfortunately, this does mean that all existing 0.4 plugins will need adapting.

## Reworked secret management and injection

In Drone 0.4, secrets were injected by string interpolation in `.drone.yml`. We protected against malicious pull requests by signing `.drone.yml` and suggested restrictions of plugins during PR builds.

In Drone 0.5, we have removed the string interpolated secrets in `.drone.yml` in favor of defining secrets via the Drone CLI. These are passed into plugins as environment variables, like other forms of input. You are able to restrict secret usage by repo, plugin image, or event types.

See [Secrets](../../../usage/secrets) for more details.

## Per-project SSH keys are completely removed

Even in Drone 0.4, `.netrc` injection was the recommended approach for pulling and cloning repositories. In Drone 0.5, the last vestigal remnants of the per-project SSH keys have been removed.

This should be transparent to most. If you are seeing authentication errors, ensure that you haven't overridden your `clone` step to assume SSH keys.
