---
date: 2017-04-15T14:39:04+02:00
title: "Why does the Docker plugin fail"
url: why-does-the-docker-plugin-fail
---


The docker plugin is highly stable and heavily used. If you are experiencing issues with the plugin it is typically an issue with your yaml or host machine configuration. This document will help identify and resolve known configuration problems.

# Debugging

The first step is to enable verbose logging for your plugin. This will write the Docker daemon logs to your build output to help discover the root cause.

```diff
pipeline:
  docker:
    image: plugins/docker
+   debug: true
```

# Driver Errors

The docker plugin uses docker in docker which supports a subset of available drivers. The errors in this section indicate use of an unsupported or missing driver.

Error when unsupported devicemapper driver is being used:

```nohighlight
level=error msg="There are no more loopback devices available."
level=fatal msg="Error starting daemon: error initializing graphdriver: loopback mounting failed"
Cannot connect to the Docker daemon. Is 'docker -d' running on this host?
```

Error when the overlay driver is used by not installed:

```nohighlight
level=error msg="'overlay' not found as a supported filesystem on this host.
Please ensure kernel is new enough and has overlay support loaded."
level=fatal msg="Error starting daemon: error initializing graphdriver: driver not supported"
```



# Driver Support

There are known issues when attempting to run this plugin on Linux distributions do not have a supported storage driver installed. To verify your storage driver please run the following command:

```nohighlight
$ docker info | grep 'Storage Driver:'
Storage Driver: aufs
```

The aufs and overlay storage drivers are support. Please note that the devicemapper storage driver is not supported.

# Driver Fallback

If your host machine does not have a supported driver installed you may be able to use the vfs storage driver. This can be configured in the yaml.

```diff
pipeline:
  publish:
    image: plugins/docker
    repo: octocat/hello-world
+   storage_driver: vfs
```

If you are using the vfs driver please be aware of the potential performance implications and disk requirements.

> The vfs backend is a very simple fallback that has no copy-on-write support. Each layer is just a separate directory. Creating a new layer based on another layer is done by making a deep copy of the base layer into a new directory.

> Since this backend doesn’t share diskspace use between layers, and since creating a new layer is a slow operation this is not a very practical backend. However, it still has its uses, for instance to verify other backends against, or if you need a super robust (if slow) backend that works everywhere.

# Kernel Issues

The below error occurs when cgroups memory features are not configured at the kernel level. This is commonly encountered with Debian distros.

```nohighlight
time="2015-12-17T08:06:57Z" level=debug msg="Mounting none /sys/fs/cgroup/blkio cgroup blkio"
time="2015-12-17T08:06:57Z" level=debug msg="Mounting none /sys/fs/cgroup/perf_event cgroup perf_event"
time="2015-12-17T08:06:57Z" level=debug msg="Mounting none /sys/fs/cgroup/cpuset cgroup cpuset"
time="2015-12-17T08:06:57Z" level=debug msg="Mounting none /sys/fs/cgroup/cpu,cpuacct cgroup cpu,cpuacct"
time="2015-12-17T08:06:57Z" level=debug msg="Creating /sys/fs/cgroup/memory"
time="2015-12-17T08:06:57Z" level=debug msg="Mounting none /sys/fs/cgroup/memory cgroup memory"
time="2015-12-17T08:06:57Z" level=fatal msg="no such file or directory"
```

Edit your `grub.cfg` and add the following to your kernel image:

```nohighlight
cgroup_enable=memory swapaccount=1
```
