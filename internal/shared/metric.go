package shared

import "time"

type Metric struct {
	Used      int64
	UsedText  string
	Total     int64
	TotalText string
	CreatedAt time.Time
}

type MetricBuffer struct {
	items     []*Metric
	maxLength int
}

func NewMetricBuffer(maxLength int) *MetricBuffer {
	return &MetricBuffer{
		items:     make([]*Metric, 0, maxLength),
		maxLength: maxLength,
	}
}

func (b *MetricBuffer) Push(item *Metric) {
	length := len(b.items)

	if length > 0 && b.items[0].Used == item.Used {
		return
	}

	newItems := make([]*Metric, 0, b.maxLength)
	newItems = append(newItems, item)

	if length == b.maxLength {
		newItems = append(newItems, b.items[:length-1]...)
	} else {
		newItems = append(newItems, b.items...)
	}

	b.items = newItems
}

func (b *MetricBuffer) All() []*Metric {
	return b.items
}
