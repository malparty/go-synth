// tones generator

package generators

import (
	"errors"
	"math"

	. "github.com/faiface/beep"
)

func SawTone(sr SampleRate, freq int) (Streamer, error) {
	if int(sr)/freq < 2 {
		return nil, errors.New("faiface beep tone generator: samplerate must be at least 2 times grater then frequency")
	}
	r := &oscStream{
		oscFunc: sawFunc,
	}
	r.stat = 0.0
	srf := float64(sr)
	ff := float64(freq)
	steps := srf / ff
	r.delta = 1.0 / steps
	return r, nil
}

func sawFunc(stat float64, delta float64) float64 {
	_, r := math.Modf(stat + delta)
	return r
}
