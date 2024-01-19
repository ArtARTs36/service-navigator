package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
)

const vendorImagePartsCount = 2
const imageVersionPartsCount = 2

type ImageParser struct {
}

func (p *ImageParser) ParseFromURL(imageURI string) *domain.NameDetails {
	imageNameParts := strings.Split(imageURI, "/")

	if len(imageNameParts) == 1 {
		partsByVersion := strings.Split(imageNameParts[0], ":")

		if len(partsByVersion) == 1 {
			// local build
			return nil
		}

		return p.createOfficialDockerImage(partsByVersion[0], partsByVersion[1])
	}

	// vendor/image:version or vendor/image
	if len(imageNameParts) == vendorImagePartsCount {
		// [image, version] or [image]
		partsByVersion := strings.Split(imageNameParts[1], ":")

		imageName := strings.Join([]string{
			imageNameParts[0],
			partsByVersion[0],
		}, "/")

		version := "latest"

		if len(partsByVersion) == imageVersionPartsCount {
			version = partsByVersion[1]
		}

		return &domain.NameDetails{
			Name:        imageName,
			Version:     version,
			RegistryURL: fmt.Sprintf("https://hub.docker.com/r/%s/%s", imageNameParts[0], partsByVersion[0]),
			Vendor:      imageNameParts[0],
		}
	}

	rURL, err := url.Parse(imageURI)

	if err != nil {
		return nil
	}

	partsByVersion := strings.Split(rURL.Path, ":")

	version := "latest"

	if len(partsByVersion) == imageVersionPartsCount {
		version = partsByVersion[1]
	}

	imageNameParts = strings.Split(partsByVersion[0], "/")
	imageName := ""
	vendor := ""

	if len(imageNameParts) == 1 {
		imageName = imageNameParts[0]
	} else if len(imageNameParts) >= 2 {
		imageName = imageNameParts[len(imageNameParts)-1]
		vendor = imageNameParts[len(imageNameParts)-2]
	}

	return &domain.NameDetails{
		Name:        imageName,
		Version:     version,
		RegistryURL: "http://" + partsByVersion[0],
		Vendor:      vendor,
	}
}

func (p *ImageParser) createOfficialDockerImage(name string, version string) *domain.NameDetails {
	return &domain.NameDetails{
		Name:        name,
		Version:     version,
		RegistryURL: fmt.Sprintf("https://hub.docker.com/_/%s", name),
	}
}
