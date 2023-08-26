package filler

import (
	"fmt"
	"strings"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
)

const vendorImagePartsCount = 2

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

		if len(partsByVersion) == 1 {
			service.Image = domain.Image{
				Name: strings.Join([]string{
					imageNameParts[0],
					partsByVersion[0],
				}, "/"),
				Version:     "latest",
				RegistryURL: fmt.Sprintf("https://hub.docker.com/r/%s/%s", imageNameParts[0], partsByVersion[0]),
			}

			return
		}

		service.Image = domain.Image{
			Name: strings.Join([]string{
				imageNameParts[0],
				partsByVersion[0],
			}, "/"),
			Version:     partsByVersion[1],
			RegistryURL: fmt.Sprintf("https://hub.docker.com/r/%s/%s", imageNameParts[0], partsByVersion[0]),
		}

		return
	}
}

func (f *ImageFiller) createOfficialDockerImage(name string, version string) domain.Image {
	return domain.Image{
		Name:        name,
		Version:     version,
		RegistryURL: fmt.Sprintf("https://hub.docker.com/_/%s", name),
	}
}
