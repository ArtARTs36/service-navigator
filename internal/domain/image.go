package domain

type Image struct {
	ID       string
	Name     string
	Size     int64
	SizeText string
	Short    ImageShort
	VCS      *VCS
	Unknown  bool
}

type ImageShort struct {
	Name        string
	Version     string
	RegistryURL string
	Vendor      string
}
