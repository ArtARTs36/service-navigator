package filler_test

import (
	"testing"

	"github.com/artarts36/service-navigator/internal/domain"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/datastruct"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/filler"
	"github.com/artarts36/service-navigator/internal/infrastructure/service/monitor"
	"github.com/stretchr/testify/assert"
)

type testFillerSetValue struct {
	value string
}

func (r *testFillerSetValue) Fill(service *domain.ServiceStatus, _ *datastruct.Container) {
	service.Name = r.value
}

type testFillerNoSetValue struct {
}

func (r *testFillerNoSetValue) Fill(_ *domain.ServiceStatus, _ *datastruct.Container) {
}

func TestOrFiller_Fill(t *testing.T) {
	cases := []struct {
		fillers       []monitor.Filler
		expectedValue string
	}{
		{
			fillers: []monitor.Filler{
				&testFillerSetValue{value: "1"},
				&testFillerSetValue{value: "2"},
			},
			expectedValue: "1",
		},
		{
			fillers: []monitor.Filler{
				&testFillerNoSetValue{},
				&testFillerSetValue{value: "2"},
			},
			expectedValue: "2",
		},
		{
			fillers: []monitor.Filler{
				&testFillerNoSetValue{},
				&testFillerSetValue{value: "2"},
				&testFillerSetValue{value: "3"},
			},
			expectedValue: "2",
		},
		{
			fillers: []monitor.Filler{
				&testFillerNoSetValue{},
			},
			expectedValue: "",
		},
	}

	for _, tCase := range cases {
		f := filler.NewOrFiller(tCase.fillers)

		service := domain.ServiceStatus{}

		f.Fill(&service, &datastruct.Container{})

		assert.Equal(t, tCase.expectedValue, service.Name)
	}
}
