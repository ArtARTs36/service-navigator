# Service Navigator

Service Navigator - navigator for your local docker projects in single network

![](./docs/preview.png)

# Config

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
