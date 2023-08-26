package config

import "strings"

type SearchProvider struct {
	Name           string `yaml:"name"`
	URL            string `yaml:"url"`
	QueryParamName string `yaml:"queryParamName"`
}

func getDefaultProviderNameURLMap() map[string]string {
	return map[string]string{
		"google":        "https://www.google.com/search",
		"stackoverflow": "https://stackoverflow.com/search",
	}
}

func ResolveProviders(providers []SearchProvider) []SearchProvider {
	newProviders := make([]SearchProvider, 0, len(providers))

	for _, provider := range providers {

		url := resolveProviderURL(provider)
		queryParamName := resolveProviderQueryParamName(provider)

		newProviders = append(newProviders, SearchProvider{
			Name:           provider.Name,
			URL:            url,
			QueryParamName: queryParamName,
		})
	}

	return newProviders
}

func resolveProviderURL(provider SearchProvider) string {
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

func resolveProviderQueryParamName(provider SearchProvider) string {
	if provider.QueryParamName == "" {
		return "q"
	}

	return provider.QueryParamName
}
