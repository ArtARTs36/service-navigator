package shared

type Metric struct {
	Used      int64
	UsedText  string
	Total     int64
	TotalText string
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
	newItems := make([]*Metric, 0, b.maxLength)
	newItems = append(newItems, item)

	if len(b.items) == b.maxLength {
		newItems = append(newItems, b.items[:len(b.items)-1]...)
	} else {
		newItems = append(newItems, b.items...)
	}

	b.items = newItems
}

func (b *MetricBuffer) All() []*Metric {
	return b.items
}
