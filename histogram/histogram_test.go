package histogram

import (
	"testing"
)

// Tests basic histogram functionality with floating-point numbers.
func TestHistogramFloat(t *testing.T) {
	t.Parallel()
	inputs := []float32{
		98.01424,
		98.01424,
		98.014241,
		0.631,
		0.632,
		0.636,
		0.0,
		0.0,
		-1.881,
		-1.88109,
		314159265358979.0,
	}

	h := New()
	for _, x := range inputs {
		h.RecordValue(x)
	}

	if h.TotalCount() != len(inputs) {
		t.Logf("Incorrect value count. Got %d; expected %d.", h.TotalCount(), len(inputs))
		t.Fail()
	}
}
