package linearstat

import (
	"testing"
)

func TestRegression(t *testing.T) {
	input := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	want1 := 1.000000
	want2 := 1.000000

	slope, intercept := CalculateLinearRegression(input)

	if slope != want1 {
		t.Errorf("input %v, slope %f, want1 %f", input, slope, want1)
	}

	if intercept != want2 {
		t.Errorf("input %v,intercept %f, want2 %f", input, intercept, want2)
	}
}
