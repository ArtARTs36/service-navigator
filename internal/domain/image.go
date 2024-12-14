package domain

import (
	"fmt"
	"time"

	"github.com/artarts36/depexplorer"
)

var ErrImageNotFound = fmt.Errorf("image not found")

type Image struct {
	ID          string
	Name        string
	Size        int64
	SizeText    string
	NameDetails NameDetails
	VCS         *VCS
	Unknown     bool
	Languages   []*Language
	Frameworks  []*depexplorer.Framework
	DepFiles    map[depexplorer.DependencyManager]*depexplorer.File
	CreatedAt   time.Time
}

type Language struct {
	Name    string
	Version string
}

type NameDetails struct {
	Name                string
	Version             string
	RegistryURL         string
	RegistryIsDockerHub bool
	Vendor              string
}

func (d *NameDetails) IsLatest() bool {
	return d.Version == "latest"
}
