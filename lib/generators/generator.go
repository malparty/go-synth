package generators

import (
	"errors"
	"math"

	"github.com/faiface/beep"
)

type GeneratorFunction func(stat float64, delta float64) float64

// create stream which will produce infinite osciator tone with the given frequency
// use other wrappers of this package to change amplitude or add time limit
// sampleRate must be at least two times grater then frequency, otherwise this function will return an error
type OscStream struct {
	Stat    float64 // progress from 0 to 1
	Delta   float64 // space between two calculation
	OscFunc GeneratorFunction
}

type generator struct {
	freq       float64
	sampleRate float64
	osc        *OscStream
}

func NewGenerator(sr beep.SampleRate, freq int, oscFunc GeneratorFunction) (*generator, error) {
	if int(sr)/freq < 2 {
		return nil, errors.New("faiface beep tone generator: samplerate must be at least 2 times grater then frequency")
	}
	osc := &OscStream{
		OscFunc: oscFunc,
		Stat:    0.0,
	}

	g := &generator{
		sampleRate: float64(sr),
		osc:        osc,
	}

	g.SetFreqInt(freq)

	return g, nil
}

func (g *generator) GetOsc() *OscStream {
	return g.osc
}

func (g *generator) GetFreq() float64 {
	return g.freq
}
func (g *generator) SetFreqInt(freq int) {
	g.SetFreq(float64(freq))
}

func (g *generator) SetFreq(freq float64) {
	if freq < 0 {
		return
	}

	steps := g.sampleRate / freq
	g.osc.Delta = 1.0 / steps

	g.freq = freq
}

func (c *OscStream) nextSample() float64 {
	r := c.OscFunc(c.Stat, c.Delta)
	_, c.Stat = math.Modf(c.Stat + c.Delta)
	return r
}

func (c *OscStream) Stream(buf [][2]float64) (int, bool) {
	for i := 0; i < len(buf); i++ {
		s := c.nextSample()
		buf[i] = [2]float64{s, s}
	}
	return len(buf), true
}

func (c *OscStream) Err() error {
	println("error with osc at stat: ", c.Stat)
	return nil
}
