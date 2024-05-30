package shared_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/artarts36/service-navigator/internal/shared"
)

func TestChunkSlice(t *testing.T) {
	cases := []struct {
		items     []string
		chunkSize int
		expected  [][]string
	}{
		{
			items:     []string{},
			chunkSize: 0,
			expected:  [][]string{},
		},
		{
			items:     []string{},
			chunkSize: 1,
			expected:  [][]string{},
		},
		{
			items: []string{
				"a",
			},
			chunkSize: 1,
			expected: [][]string{
				{"a"},
			},
		},
		{
			items: []string{
				"a",
			},
			chunkSize: 2,
			expected: [][]string{
				{"a"},
			},
		},
		{
			items: []string{
				"a", "b", "c",
				"d", "e", "f",
			},
			chunkSize: 2,
			expected: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e", "f"},
			},
		},
		{
			items: []string{
				"a", "b", "c",
				"d", "e", "f",
				"g",
			},
			chunkSize: 2,
			expected: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e", "f"},
				{"g"},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			chunks := shared.ChunkSlice(c.items, c.chunkSize)

			assert.Equal(t, c.expected, chunks)
		})
	}
}
