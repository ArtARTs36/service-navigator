package entity

type Service struct {
	Name        string
	WebURL      string
	Status      string
	VCS         *VCS
	ContainerID string
}
