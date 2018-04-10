+++
date = "2017-04-15T14:39:04+02:00"
title = "Substitution"
url = "substitution"

[menu.usage]
  weight = 4
  identifier = "substitution"
  parent = "usage_config"
+++

Drone provides the ability to substitute repository and build metadata to facilitate dynamic pipeline configurations.

Example commit substitution:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: "${DRONE_COMMIT_SHA}"
```

Example tag substitution:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: "${DRONE_TAG}"
```

# String Operations

Drone provides partial emulation for bash string operations. This can be used to manipulate string values prior to substitution.

Example variable substitution with substring:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: "${DRONE_COMMIT_SHA:0:8}"
```

Example variable substitution strips `v` prefix from `v1.0.0`:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: "${DRONE_TAG##v}"
```

Example variable substitution replaces `/` with `-`:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: ${DRONE_BRANCH/\//-}
```

# String Operations Reference

List of emulated string operations:

OPERATION             | DESC
----------------------|---------------------------------------------------------
`${param}`            | parameter substitution
`${param,}`           | parameter substitution with lowercase first char
`${param,,}`          | parameter substitution with lowercase
`${param^}`           | parameter substitution with uppercase first char
`${param^^}`          | parameter substitution with uppercase
`${param:pos}`        | parameter substitution with substring
`${param:pos:len}`    | parameter substitution with substring and length
`${param=default}`    | parameter substitution with default
`${param##prefix}`    | parameter substitution with prefix removal
`${param%%suffix}`    | parameter substitution with suffix removal
`${param/old/new}`    | parameter substitution with find and replace

# Environment Variables Reference

Please see the reference environment [documentation]({{< ref "usage/reference/environment.md" >}}) for the complete list of Drone environment variables.
