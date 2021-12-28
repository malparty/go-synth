package generators_test

import (
	"testing"

	"github.com/malparty/go-synth/lib/generators"
)

func TestSawFunc(t *testing.T) {
	scenes := []struct {
		stat   float64
		delta  float64
		result float64
	}{
		// From initial position
		{0, 0, -1},
		{0, 0.25, -0.5},
		{0, 0.5, 0},
		{0, 0.75, 0.5},

		// From middle position
		{0.5, 0, 0},
		{0.5, 0.25, 0.5},
		{0.5, 0.5, -1},

		// Overpassing level
		{1, 1, -1},
		{1, 0.5, 0},
		{0.75, 0.75, 0},
	}

	for _, scene := range scenes {
		if r := generators.SawFunc(scene.stat, scene.delta); r != scene.result {
			t.Errorf("Expect SawFunc(%f, %f) to return %f but got %f instead", scene.stat, scene.delta, scene.result, r)
		}
	}
}
