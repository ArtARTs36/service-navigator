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
