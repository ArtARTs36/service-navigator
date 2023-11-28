package filler

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	vcs2 "github.com/artarts36/service-navigator/internal/infrastructure/service/vcs"
)

type VCSFiller struct {
}

func (f *VCSFiller) Fill(image *domain.Image, meta *datastruct.ImageMeta) {
	vcs, err := vcs2.ParseFromLabels(meta.Labels)
	if err != nil {
		if !errors.Is(err, vcs2.ErrNotFound) {
			log.
				WithField("image", image.Name).
				Warnf("[Image][VCSFiller] vcs resolving failed: %s", err)
		}

		return
	}

	image.VCS = vcs
}
