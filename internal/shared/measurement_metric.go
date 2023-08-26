package shared

import "time"

type MeasurementMetric struct {
	Used      int64
	UsedText  string
	Total     int64
	TotalText string
	CreatedAt time.Time
}

type MeasurementMetricBuffer struct {
	items     []*MeasurementMetric
	maxLength int
	unique    bool
}

func NewMetricBuffer(maxLength int, unique bool) *MeasurementMetricBuffer {
	return &MeasurementMetricBuffer{
		items:     make([]*MeasurementMetric, 0, maxLength),
		maxLength: maxLength,
		unique:    unique,
	}
}

func (b *MeasurementMetricBuffer) Push(item *MeasurementMetric) {
	length := len(b.items)

	if b.unique && length > 0 && b.items[0].Used == item.Used {
		return
	}

	newItems := make([]*MeasurementMetric, 0, b.maxLength)
	newItems = append(newItems, item)

	if length == b.maxLength {
		newItems = append(newItems, b.items[:length-1]...)
	} else {
		newItems = append(newItems, b.items...)
	}

	b.items = newItems
}

func (b *MeasurementMetricBuffer) All() []*MeasurementMetric {
	return b.items
}
