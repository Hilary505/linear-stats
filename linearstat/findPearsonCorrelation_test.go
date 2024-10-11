package linearstat

import (
	"testing"
)

func TestPearson(t *testing.T) {
	input := CalculatePearsonCorrelation([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	want := 1.0

	if input != want {
		t.Errorf("input %v want %f", input, want)
	}

	input2 := CalculatePearsonCorrelation([]float64{-1.0, -2.0, -4.0, 6.0, 5.0})
	want2 := 0.7124704998790964

	if input2 != want2 {
		t.Errorf("input2 %v want2 %f", input2, want2)
	}
}