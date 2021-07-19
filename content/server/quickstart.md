---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: tphoney
weight: 1
toc: true
description: |
  Installation quickstart.
---

This section of the documentation will help you install and configure Drone quickly. There are 3 different paths you can take

- Try out `Drone in the cloud`, using our servers to try out building and running plugins.
- Try out `Local Developer Setup`, use docker compose to setup a build server on your system using github as your source control.
- Try out `Server install`,
  
# Overview

Super quick before you start: Drone is made up of 2 components, the Drone _Server_ and one or more _Runners_. The drone server hosts the UI and the DB along with scheduling builds. A runner is a standalone daemon that polls the server for pending pipelines to execute.

# Drone in the cloud

Pros
- quick to get started and building

Cons
- not configurable
- free for OSS

## Pre-requisites
Log into cloud.drone.io. you will need a github account to proceed, and that account will need to be at least 30 days old.

## Setup

1. Go to https://github.com/drone/drone_trainer
2. fork the repo into your namespace.
3. log into cloud.drone.io
4. click the sync button
5. click your forked repo and add it. accept the settings
6. run a build. 

## Further steps 

- edit the build file 
- use a different build file
- create a pr

# Local Developer Setup

Pros
- quick to get started
- small footprint

Cons
- configured to work with github
- uses docker only
- only one runner

## Pre-requisites

You have a github account, and you have docker desktop installed.


steps
