The docker plugin will automatically read tags from a .tags file, you can use the .tags file with compatible plugins such as ECR and GCR.
Example of tags with GCR
build:
  image: golang:1.10.0-alpine
  commands:
    - [...]
-   - export VERSION=$(git-buildnumber)
+   - echo -n $(git-buildnumber) > .tags
push:
  image: plugins/gcr
  registry: gcr.io
  repo: ${DRONE_REPO_NAME}
- tags: "${VERSION}"
  secrets: [google_credentials]
