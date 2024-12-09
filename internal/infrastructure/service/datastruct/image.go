package datastruct

type ImageMeta struct {
	Labels   map[string]string
	RepoTags []string
	Created  int64
}
