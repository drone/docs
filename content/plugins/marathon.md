+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Marathon"
description = "Publish applications to Mesos cluster through Marathon"
user = "drone-plugins"
repo = "drone-marathon"
image = "plugins/marathon"
tags = ["mesos", "marathon", "cluster", "docker"]
categories = "publish"
draft = false
date = 2016-07-30T23:50:05Z
menu = ""
weight = 1

+++


This plugin can be used to deploy applications to [Marathon](https://mesosphere.github.io/marathon/).

The following parameters are used to configure this plugin:

| Parameter     | Description                | |
|---------------|----------------------------|-------
|`server`                     | marathon complete url | 
|`username`                   | (optional) authentication username | 
|`password`                   | (optional) authentication password | 
|`id`                         | unique identifier for the application | 
|`instances`                  | the number of instances of this application to start |
|`cpus`                       | the number of CPU shares this application needs per instance |
|`mem`                        | the amount of memory in MB that is needed for the application per instance |
|`cmd`                        | the command that is executed.  This value is wrapped by Mesos via `/bin/sh -c ${app.cmd}` |
|`args`                       | an array of strings that represents an alternative mode of specifying the command to run |
|`uris`                       | an array of URIs resolved before the application gets started (Deprecated since Marathon v0.15.0) |
|`fetch`                      | an array of URIs are passed to Mesos fetcher module and resolved in runtime (Available since Marathon v0.15.0) |
||*`uri`*                     | URI to be fetched by Mesos fetcher module |
||*`executable`*              | (optional default: false) set fetched artifact as executable |
||*`extract`*                 | (optional default: false) extract fetched artifact if supported by Mesos fetcher module |
||*`cache`*                   | (optional default: false) cache fetched artifact if supported by Mesos fetcher module |
|`health_checks`              | an array of checks to be performed on running tasks. Keys are: |
||*`protocol`*                | (optional default: HTTP) Protocol of the requests to be performed *[HTTP, HTTPS, TCP, COMMAND]*
||*`path`*                    | (optional default: /) Path to endpoint exposed by the task that will provide health status 
||*`port`*                    | the specific port to connect to. In case of dynamic ports, see port_index 
||*`port_index`*              | index in this app's ports array to be used for health requests 
||*`command`*                 | the health check command to execute 
||*`grace_period_seconds`*    | health check failures are ignored within this number of seconds of the task being started or until the task becomes healthy for the first time 
||*`interval_seconds`*        | number of seconds to wait between health checks
||*`timeout_seconds`*         | number of seconds after which a health check is considered a failure regardless of the response
||*`max_consecutive_failures`*| number of consecutive health check failures after which the unhealthy task should be killed
|`min_health_capacity`        | (optional default: 1) a number between 0 and 1 that is multiplied with the instance count. Marathon will make sure, during the upgrade process, that at any point of time this number of healthy instances are up |
|`max_over_capacity`          | (optional default: 1) a number between 0 and 1 which is multiplied with the instance count. This is the maximum number of additional instances launched at any point of time during the upgrade process |
|`constraints`                | an array of constraints for the application. Keys are: |
||*`field`*                   | the constraint subject
||*`operator`*                | possible values are *[UNIQUE, CLUSTER, GROUP_BY, LIKE, UNLIKE]*
||*`value`*                   | (optional) the constraint value
|`accepted_resource_roles`    | an array of resource roles |
|`backoff_factor`             | configures exponential backoff behavior when launching potentially sick apps. The backoff period is multiplied by the factor for each consecutive failure until it reaches max_launch_delay_seconds |
|`backoff_seconds`            | see above |
|`max_launch_delay_seconds`   | see above |
|`dependencies`               | an array of services upon which this application depends |
|`disk`                       | (optional default: 0) how much disk space is needed for this application |
|`process_environment`        | key value pairs that get added to the environment variables of the process to start |
|`docker_image`               | the name of the docker image to use |
|`docker_network`             | (optional default: BRIDGE) the networking mode, possible values are *[BRIDGE, HOST, NONE]* | 
|`docker_privileged`          | (optional default: false) run docker image in privileged mode |
|`docker_force_pull`          | (optional default: false) image will be pulled, regardless if it is already available on the local system |
|`docker_port_mappings`       | an array of host-container port mapping. Keys are: |
||*`container_port`*          | (optional default: 0) the port the application listens to inside of the container. If 0 marathon assigns the container port the same value as the assigned host_port 
||*`host_port`*               | (optional default: 0) random port from the range included in the Mesos resource offer
||*`service_port`*            | (optional default: 0) helper port intended for doing service discovery
||*`protocol`*                | (optional default: tcp) protocol of the port. Possible values are *[tcp, udp, udp,tcp]*
|`docker_volumes`             | an array of volumes mapped in the container. Keys are: |
||*`container_path`*          | the path of the volume in the container
||*`host_path`*               | the path of the volume on the host
||*`mode`*                    | possible values are RO for ReadOnly and RW for Read/Write 
|`docker_parameters`          | a map of arbitrary parameters to be passed to docker CLI |
|`debug`                      | print request and response info |


The following are sample configuration for your .drone.yml file.

*Simple deploy using a Docker image*

```yaml
  marathon:
    server: http://marathon.mycluster.io:8080
    id: myapp
    instances: 1
    cpus: 0.5
    mem: 64.0
    docker_image: busybox
    cmd: while [ true ] ; do echo 'Hello Drone' ; sleep 5 ; done
```

*More complicate deploy*

```yaml
  marathon:
    server: http://marathon.mycluster.io:8080
    id: myapp
    instances: 1
    cpus: 0.5
    mem: 64.0
    docker_image: acme/someapp
    uris:
      - http://placehold.it/350x150
      - http://placehold.it/150x150
    health_checks:
      - protocol: HTTP
        path: /health
        port: 1111
        grace_period_seconds: 3
        interval_seconds: 10
        timeout_seconds: 10
        max_Consecutive_failures: 3
    min_health_capacity: 0.4
    max_over_capacity: 1
    constraints:
      - field: hostname
        operator: UNIQUE
    docker_port_mappings:
      - container_port: 8080
        host_port: 0
        protocol: tcp
    docker_volumes:
      - container_path: /foo
        host_path: /tmp
        mode: RW
    docker_parameters:
      hostname: uswastla05.acme.eu
    process_environment:
      BAMBOO_ZK_HOST: zk.acme.eu
```

