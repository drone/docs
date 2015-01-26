+++
weight = 0
date = "2015-01-26T12:54:12-08:00"
title = "Ruby On Rails"

[menu.tutorials]
parent = "Frameworks"
+++

## Cache gems
You need modify `.drone.yml` file:

```coffeescript
...
script:
  - mkdir -p /tmp/bundle
  - sudo chown -R ubuntu:ubuntu /tmp/bundle
  - bundle install --path /tmp/bundle
  - ...
cache:
  - /tmp/bundle
  - ...
...
```

## Testing with PostgreSQL

You need add to your project `config/database.yml.example` with content:

```coffeescript
test: &test
  host: localhost
  template: template0
  adapter: postgresql
  encoding: unicode
  database: test
  username: postgres
  password:
  port: 5432
```

Then you need modify `.drone.yml` file:

```coffeescript
...
services:
  - postgresql
  - ...
script:
  - ...
  - cp config/database.yml.example config/database.yml
  - bundle exec rake db:create db:migrate
  - ...
...
```

## Examples

Testing with RSpec ( using PostgreSQL, Redis and Gem caching )

```coffeescript
image: ruby
services:
  - postgresql
  - redis
script:
  - mkdir -p /tmp/bundle
  - sudo chown -R ubuntu:ubuntu /tmp/bundle
  - bundle install -j8 --path /tmp/bundle
  - cp config/database.yml.example config/database.yml
  - bundle exec rake db:create db:migrate spec
cache:
  - /tmp/bundle
```