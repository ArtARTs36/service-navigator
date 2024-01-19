package domain

type Image struct {
	ID       string
	Name     string
	Size     int64
	SizeText string
	Short    NameDetails
	VCS      *VCS
	Unknown  bool
}

type NameDetails struct {
	Name        string
	Version     string
	RegistryURL string
	Vendor      string
}
