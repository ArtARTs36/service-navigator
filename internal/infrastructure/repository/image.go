package repository

import "github.com/artarts36/service-navigator/internal/domain"

type ImageRepository struct {
	images []*domain.Image
}

func (r *ImageRepository) Set(images []*domain.Image) {
	r.images = images
}

func (r *ImageRepository) All() []*domain.Image {
	return r.images
}

func (r *ImageRepository) FindByID(id string) (*domain.Image, error) {
	for _, image := range r.images {
		if image.ID == id {
			return image, nil
		}
	}
	return nil, domain.ErrImageNotFound
}
