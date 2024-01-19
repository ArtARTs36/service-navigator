package domain

type Image struct {
	ID          string
	Name        string
	Size        int64
	SizeText    string
	NameDetails NameDetails
	VCS         *VCS
	Unknown     bool
}

type NameDetails struct {
	Name                string
	Version             string
	RegistryURL         string
	RegistryIsDockerHub bool
	Vendor              string
}
