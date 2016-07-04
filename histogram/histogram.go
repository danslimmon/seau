package histogram

type Histogram struct {
	totalCount int
}

func New() *Histogram {
	return new(Histogram)
}

func (h *Histogram) TotalCount() int {
	return h.totalCount
}

func (h *Histogram) RecordValue(x float32) {
	h.totalCount++
}
