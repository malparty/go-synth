package effects

import (
	"github.com/malparty/go-synth/lib/generators"
)

type Reverb struct {
	MixRate  float64
	FadeRate float64
	Time     int

	buffer       []float64
	currentIndex int
}

func (r *Reverb) SetTime(time int) {
	r.Time = time
	if len(r.buffer) < r.Time {
		length := len(r.buffer)
		for i := length; i < r.Time; i++ {
			r.buffer = append(r.buffer, 0)
		}
	}
}

func (r *Reverb) GetReverbFunc() generators.GeneratorFunction {
	r.currentIndex = 0
	r.buffer = []float64{}

	for i := 0; i <= r.Time; i++ {
		r.buffer = append(r.buffer, 0)
	}

	return func(stat float64, _ float64) (reverbLevel float64) {
		reverbLevel = r.buffer[r.currentIndex] * r.FadeRate / 100

		r.buffer[r.currentIndex] = reverbLevel + stat*r.MixRate/100
		r.currentIndex += 3
		if r.currentIndex > r.Time-3 {
			r.currentIndex = 0
		}

		return reverbLevel + stat
	}
}
