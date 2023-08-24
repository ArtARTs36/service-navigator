package search

import "strings"

type Provider struct {
	Name           string `yaml:"name"`
	Url            string `yaml:"url"`
	QueryParamName string `yaml:"queryParamName"`
}

func ResolveProviders(providers []Provider) []Provider {
	newProviders := make([]Provider, 0, len(providers))

	for _, provider := range providers {

		url := resolveProviderUrl(provider)
		queryParamName := resolveProviderQueryParamName(provider)

		newProviders = append(newProviders, Provider{
			Name:           provider.Name,
			Url:            url,
			QueryParamName: queryParamName,
		})
	}

	return newProviders
}

func resolveProviderUrl(provider Provider) string {
	var url string

	if provider.Url == "" {
		switch strings.ToLower(provider.Name) {
		case "google":
			url = "https://www.google.com/search"
			break
		case "stackoverflow":
			url = "https://stackoverflow.com/search"
			break
		}
	} else {
		url = provider.Url
	}

	return url
}

func resolveProviderQueryParamName(provider Provider) string {
	if provider.QueryParamName == "" {
		return "q"
	}

	return provider.QueryParamName
}
