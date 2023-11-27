package domain

type Image struct {
	ID       string
	Name     string
	Size     int64
	SizeText string
	Short    ImageShort
	VCS      *VCS
}

type ImageShort struct {
	Name        string
	Version     string
	RegistryURL string
}
