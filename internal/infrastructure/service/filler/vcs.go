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

func (r *VCSFiller) Fill(service *domain.ServiceStatus, container *datastruct.Container) {
	vcs, err := vcs2.ParseFromLabels(container.Short.Labels)
	if err != nil {
		if !errors.Is(err, vcs2.ErrNotFound) {
			log.Warnf("[Service][VCSFiller] vcs resolving failed: %s", err)
		}

		return
	}

	service.VCS = vcs
}
