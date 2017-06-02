
```diff
pipeline:
  build:
    image: maven
    commands:
      - mvn install
      - mvn package
    environment:
+     - MAVEN_HOME=/drone/.m2
+     - M2_HOME=/drone/.m2
+     - GRADLE_USER_HOME=/drone/.gradle
```
