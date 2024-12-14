package vcs

import (
	"testing"

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
				labelOpenContainerImageSource: "https://github.com/prometheus/prometheus",
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
			got, err := ParseFromLabels(tc.Labels)
			require.NoError(t, err)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
