


```diff
pipeline:
  build:
    image: ruby
    commands:
      - bundler install
    environment:
+     - GEM_HOME=
+     - GEM_PATH=
+     - BUNDLE_PATH=
+     - BUNDLE_BIN=
+     - BUNDLE_CACHE_PATH=
```
