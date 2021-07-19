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
- quick to get started

Cons
- not configurable


# Local Developer Setup

Pros
- quick to get started
- small footprint

Cons
- configured to work with github
- uses docker only
- only one runner

pre requisites

steps
