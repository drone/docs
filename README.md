[![Build Status](https://cloud.drone.io/api/badges/drone/docs/status.svg)](https://cloud.drone.io/drone/docs)

This repository contains the source for [docs.drone.io](http://docs.drone.io).
To generate the documentation you will need to download and install the [hugo](https://gohugo.io/overview/installing/) static website engine.

Generate the documentation:

```
hugo
```

Generate the documentation and serve at `localhost:1313`:

```
hugo server -b localhost:1313 -w
```
