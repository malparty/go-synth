package effects

import (
	"fmt"

	"github.com/malparty/go-synth/lib"
	"github.com/malparty/go-synth/lib/generators"
)

type Reverb struct {
	MixRate  float64
	FadeRate float64
	DelayMs  float64
	Freq     float64

	buffer        []float64
	bufferb       []float64
	bufferc       []float64
	bufferd       []float64
	currentIndex  int
	currentIndexb int
	currentIndexc int
	currentIndexd int
	delta         float64
}

func (r *Reverb) SetDelay(delay float64) {
	r.DelayMs = delay

	fmt.Printf("SET DELAYS: %.0f \n", delay)

	r.resetBufferSize()
}

func (r *Reverb) GetReverbFunc() generators.GeneratorFunction {
	r.currentIndex = 0
	r.buffer = []float64{}
	fmt.Printf("FREQ: %.0f \n", r.Freq)
	r.delta = r.Freq / float64(lib.SampleRate)

	r.resetBufferSize()

	return func(stat float64, delta float64) (reverbLevel float64) {
		if r.currentIndex >= len(r.buffer) {
			r.currentIndex = 0
		}
		if r.currentIndexb >= len(r.bufferb) {
			r.currentIndexb = 0
		}
		if r.currentIndexc >= len(r.bufferc) {
			r.currentIndexc = 0
		}
		if r.currentIndexd >= len(r.bufferd) {
			r.currentIndexd = 0
		}

		reverbLevel = r.buffer[r.currentIndex] * r.FadeRate / 100
		reverbLevelb := r.bufferb[r.currentIndexb] * r.FadeRate / 100
		reverbLevelc := r.bufferc[r.currentIndexc] * r.FadeRate / 100
		reverbLeveld := r.bufferd[r.currentIndexd] * r.FadeRate / 100

		r.buffer[r.currentIndex] = reverbLeveld + stat*r.MixRate/100
		r.bufferb[r.currentIndexb] = reverbLevel + stat*r.MixRate/100
		r.bufferc[r.currentIndexc] = reverbLevelb + stat*r.MixRate/100
		r.bufferd[r.currentIndexd] = reverbLevelc + stat*r.MixRate/100
		r.currentIndex++
		r.currentIndexb++
		r.currentIndexc++
		r.currentIndexd++

		return reverbLeveld + stat*r.MixRate/100
	}
}

func (r *Reverb) resetBufferSize() {
	r.buffer = []float64{}

	steps := (1.0 / r.delta)

	fmt.Printf("Delta: %.5f \n", r.delta)
	fmt.Printf("Steps: %.1f \n", steps)

	bufferSize := steps * r.DelayMs / 1000

	r.buffer = r.resizeBuffer(r.buffer, bufferSize)
	r.bufferb = r.resizeBuffer(r.bufferb, bufferSize+5)
	r.bufferc = r.resizeBuffer(r.bufferc, bufferSize+3)
	r.bufferd = r.resizeBuffer(r.bufferd, bufferSize-5)
}

func (r *Reverb) resizeBuffer(buffer []float64, size float64) []float64 {
	sizeInt := int(size)
	for i := 0; i < sizeInt; i++ {
		buffer = append(buffer, 0)
	}

	fmt.Printf("Buffer size: %v \n\n", len(buffer))

	return buffer
}
