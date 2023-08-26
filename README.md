# Service Navigator

Service Navigator - navigator for your local docker projects in single network

![](./docs/preview.png)

# Config

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
      - url: http://github.com/artarts36/service-navigator
        title: Github
    profile:
      links:
        - url: http://github.com/artarts36/service-navigator
          title: Github
    search:
      providers:
        - name: google
        - name: stackoverflow
        - name: Jira
          url: https://jira.host.name/secure/QuickSearch.jspa
          queryParamName: searchString

# This section contains settings for backend
backend:
  # Docker network name
  #
  # Required
  network_name: infra

  # Poll for finding information about services
  poll:
    # Interval for services polling
    #
    # Default: "2s"
    interval: "2s"

    metrics:
      # Count of stored records per service
      #
      # Optional, default: 50
      depth: 10

      # A flag that determines whether to store only unique metrics per service
      #
      # Optional, default: false
      only_unique: true
```

## Add Jira Search Provider

You can add search for Jira.

```yaml
frontend:
  navbar:
    search:
      providers:
        - name: Jira
          url: https://jira.host.name/secure/QuickSearch.jspa
          queryParamName: searchString
```

# How Service Navigator finding information about service

## Resolving service url

Service Navigator considers the `NGINX_PROXY` environment variable to be the web url.

## Resolving repository url

Service Navigator looks at labels:
* `org.service_navigator.gitlab_repository`
* `org.service_navigator.github_repository`
* `org.service_navigator.bitbucket_repository`
* `org.opencontainers.image.source`
