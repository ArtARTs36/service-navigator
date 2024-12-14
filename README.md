# Service Navigator

![Testing](https://github.com/ArtARTs36/service-navigator/workflows/Testing/badge.svg?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker Pulls](https://img.shields.io/docker/pulls/artarts36/service-navigator)](https://hub.docker.com/r/artarts36/service-navigator)

Service Navigator - navigator for your local docker projects in single network

![](./docs/preview.png)

# Setup

1. Download config file: `curl https://raw.githubusercontent.com/ArtARTs36/service-navigator/master/service_navigator.yaml > service_navigator.yaml`
2. Define docker network name in `service_navigator.yaml` in section `backend.network_name`
3. Add next lines into your **docker-compose.yaml**:
```yaml
services:
  infra:
    image: artarts36/service-navigator:0.1.5
    ports:
      - "9101:8080"
    volumes:
      - type: bind
        source: "/var/run/docker.sock"
        target: "/var/run/docker.sock"
        read_only: true
      - ./:/app
    environment:
      USER: "${USER}"
    networks:
      - {YOUR_NETWORK_NAME}
```

## Config

Config is described in YAML file with name [service_navigator.yaml](./service_navigator.yaml)

```yaml
# This section contains settings for frontend
frontend:
  # Application Name
  #
  # Optional, default: "ServiceNavigator"
  app_name: ServiceNavigator

  # Navbar in header
  navbar:
    links:
      - url: /
        title: Services
      - url: /images
        title: Images
      - url: /volumes
        title: Volumes
      - url: http://github.com/artarts36/service-navigator
        title: Github
    profile:
      links:
        - url: http://github.com/artarts36/service-navigator
          title: Github
        - url: http://my-iam.service/login
          title: Generate IAM Token
          form:
            method: "POST"
            inputs:
              - name: "login"
                value: "developer"
              - name: "password"
                value: "developer"
    search:
      providers:
        - name: google
        - name: stackoverflow
        - name: Jira
          url: https://jira.host.name/secure/QuickSearch.jspa
          queryParamName: searchString # search <form> input name

  # Pages config
  #
  # Optional
  pages:
    # Images page config
    #
    # Optional
    images:
      # Settings for selecting the display of counters for an image
      #
      # Optional, default: no show counters
      counters:
        # Show image pulls count
        #
        # Optional, default: false
        pulls: false
        # Show image stars
        #
        # Optional, default: false
        stars: false

# This section contains settings for backend
backend:
  # Docker network name
  #
  # Required
  network_name: ${NETWORK_NAME}

  # Services configuration
  services:
    # Poll for finding information about services
    poll:
      # Interval for services polling
      #
      # Default: "2s"
      interval: "2s"

      # Count of goroutines for polling
      #
      # Default: 0 = count of services
      concurrent: 2

      metrics:
        # Count of stored records per service
        #
        # Optional, default: 50
        depth: 10

        # A flag that determines whether to store only unique metrics per service
        #
        # Optional, default: false
        only_unique: true

  # Images configuration
  images:
    poll:
      # Interval for images polling
      #
      # Default: "1m"
      interval: "10m"

      # Scan image repository configuration
      scan_repo:
        # Determine main repository languages
        #
        # Default: false
        lang: true

        # Fetch repository dependencies
        #
        # Default: false
        deps: true

  # Volumes configuration
  volumes:
    poll:
      # Interval for volumes polling
      #
      # Default: "1m"
      interval: "1m"

# Application parameters
parameters:
  # Log Level
  #
  # Available values: trace, debug, info, warn, error, fatal, panic
  log_level: debug

# Credentials for services
credentials:
  github_token: ${GITHUB_TOKEN}
```

# How Service Navigator finding information about service

## Resolving service url

Service Navigator checks:
* `NGINX_PROXY` environment variable
* Public port as `http://localhost:{PORT}`

## Resolving repository url

Service Navigator looks at labels:
* `org.service_navigator.gitlab_repository`
* `org.service_navigator.github_repository`
* `org.service_navigator.bitbucket_repository`
* `org.opencontainers.image.source`
