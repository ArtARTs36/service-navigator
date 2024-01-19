package parser

import (
	"fmt"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
)

const (
	vendorImagePartsCount  = 2
	imageVersionPartsCount = 2
	dockerHubHost          = "https://hub.docker.com"
)

type ImageParser struct {
}

func (p *ImageParser) ParseFromURL(imageURI string) *domain.NameDetails {
	partsByVersion := strings.Split(imageURI, ":")
	imageNameParts := strings.Split(partsByVersion[0], "/")

	if len(imageNameParts) == 1 && len(partsByVersion) == 1 {
		// local build
		return nil
	}

	version := "latest"

	if len(partsByVersion) == imageVersionPartsCount {
		version = partsByVersion[1]
	}

	registryURL := "http://" + partsByVersion[0]
	imageName := ""
	vendor := ""

	if len(imageNameParts) == 1 {
		imageName = imageNameParts[0]
		registryURL = fmt.Sprintf("%s/_/%s", dockerHubHost, imageName)
	} else if len(imageNameParts) >= 2 {
		imageName = imageNameParts[len(imageNameParts)-1]
		vendor = imageNameParts[len(imageNameParts)-2]

		if len(imageNameParts) == vendorImagePartsCount {
			registryURL = fmt.Sprintf("%s/r/%s/%s", dockerHubHost, vendor, imageName)
		}
	}

	return &domain.NameDetails{
		Name:        imageName,
		Version:     version,
		RegistryURL: registryURL,
		Vendor:      vendor,
	}
}
