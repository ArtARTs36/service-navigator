package filler

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const vendorImagePartsCount = 2
const imageVersionPartsCount = 2

type ImageFiller struct {
}

func (f *ImageFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	imageNameParts := strings.Split(container.Short.Image, "/")

	if len(imageNameParts) == 1 {
		partsByVersion := strings.Split(imageNameParts[0], ":")

		if len(partsByVersion) == 1 {
			// local build
			return
		}

		service.Image = f.createOfficialDockerImage(partsByVersion[0], partsByVersion[1])

		return
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

		service.Image = domain.ImageShort{
			Name:        imageName,
			Version:     version,
			RegistryURL: fmt.Sprintf("https://hub.docker.com/r/%s/%s", imageNameParts[0], partsByVersion[0]),
		}

		return
	}

	rURL, err := url.Parse(container.Short.Image)

	if err != nil {
		return
	}

	partsByVersion := strings.Split(rURL.Path, ":")
	version := "latest"
	imageName := partsByVersion[0]

	if len(partsByVersion) == imageVersionPartsCount {
		version = partsByVersion[1]
	}

	service.Image = domain.ImageShort{
		Name:        imageName,
		Version:     version,
		RegistryURL: "http://" + imageName,
	}
}

func (f *ImageFiller) createOfficialDockerImage(name string, version string) domain.ImageShort {
	return domain.ImageShort{
		Name:        name,
		Version:     version,
		RegistryURL: fmt.Sprintf("https://hub.docker.com/_/%s", name),
	}
}
