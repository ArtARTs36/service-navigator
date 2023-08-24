package search

import "strings"

type Provider struct {
	Name           string `yaml:"name"`
	URL            string `yaml:"url"`
	QueryParamName string `yaml:"queryParamName"`
}

func ResolveProviders(providers []Provider) []Provider {
	newProviders := make([]Provider, 0, len(providers))

	for _, provider := range providers {

		url := resolveProviderURL(provider)
		queryParamName := resolveProviderQueryParamName(provider)

		newProviders = append(newProviders, Provider{
			Name:           provider.Name,
			URL:            url,
			QueryParamName: queryParamName,
		})
	}

	return newProviders
}

func resolveProviderURL(provider Provider) string {
	var url string

	if provider.URL == "" {
		prName := strings.ToLower(provider.Name)

		if defaultURL, exists := getDefaultProviderNameURLMap()[prName]; exists {
			url = defaultURL
		}
	} else {
		url = provider.URL
	}

	return url
}

func resolveProviderQueryParamName(provider Provider) string {
	if provider.QueryParamName == "" {
		return "q"
	}

	return provider.QueryParamName
}
