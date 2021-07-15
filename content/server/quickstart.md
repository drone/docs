---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: tphoney
weight: 1
toc: true
description: |
  Installation quickstart.
---

This section of the documentation will help you install and configure Drone quickly. If you want to try Drone without installing you can try our playground <here>. 

Drone is made up of 2 components, the Drone _Server_ and a _Runners_. The drone server hosts the UI and the DB along with scheduling builds. A runner is a standalone daemon that polls the server for pending pipelines to execute.

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
