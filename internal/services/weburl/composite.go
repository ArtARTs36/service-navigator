package weburl

import "github.com/docker/docker/api/types"

type CompositeResolver struct {
	resolvers []UrlResolver
}

func NewCompositeResolver(resolvers []UrlResolver) UrlResolver {
	return &CompositeResolver{
		resolvers: resolvers,
	}
}

func (r *CompositeResolver) Resolve(container *types.ContainerJSON) string {
	for _, resolver := range r.resolvers {
		val := resolver.Resolve(container)

		if val != "" {
			return val
		}
	}

	return ""
}
