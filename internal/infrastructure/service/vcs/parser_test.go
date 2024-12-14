package vcs_test

import (
	"testing"

	"github.com/artarts36/service-navigator/internal/infrastructure/service/vcs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/artarts36/service-navigator/internal/domain"
)

func TestParseFromLabels(t *testing.T) {
	cases := []struct {
		Title    string
		Labels   map[string]string
		Expected *domain.VCS
	}{
		{
			Title: "parse opencontainers label",
			Labels: map[string]string{
				"org.opencontainers.image.source": "https://github.com/prometheus/prometheus",
			},
			Expected: &domain.VCS{
				Type: "github",
				URL:  "https://github.com/prometheus/prometheus",
				Host: "github.com",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Title, func(t *testing.T) {
			got, err := vcs.ParseFromLabels(tc.Labels)
			require.NoError(t, err)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
