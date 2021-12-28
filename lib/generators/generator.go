package generators

import (
	"errors"
	"math"

	"github.com/faiface/beep"
)

type oscFunction func(stat float64, delta float64) float64

// create stream which will produce infinite osciator tone with the given frequency
// use other wrappers of this package to change amplitude or add time limit
// sampleRate must be at least two times grater then frequency, otherwise this function will return an error
type oscStream struct {
	stat    float64 // progress from 0 to 1
	delta   float64 // space between two calculation
	oscFunc oscFunction
}

type generator struct {
	freq       float64
	sampleRate float64
	osc        *oscStream
}

func NewGenerator(sr beep.SampleRate, freq int, oscFunc oscFunction) (*generator, error) {
	if int(sr)/freq < 2 {
		return nil, errors.New("faiface beep tone generator: samplerate must be at least 2 times grater then frequency")
	}
	osc := &oscStream{
		oscFunc: oscFunc,
		stat:    0.0,
	}

	g := &generator{
		sampleRate: float64(sr),
		osc:        osc,
	}

	g.SetFreqInt(freq)

	return g, nil
}

func (g *generator) GetOsc() *oscStream {
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
	g.osc.delta = 1.0 / steps

	g.freq = freq
}

func (c *oscStream) nextSample() float64 {
	r := c.oscFunc(c.stat, c.delta)
	_, c.stat = math.Modf(c.stat + c.delta)
	return r
}

func (c *oscStream) Stream(buf [][2]float64) (int, bool) {
	for i := 0; i < len(buf); i++ {
		s := c.nextSample()
		buf[i] = [2]float64{s, s}
	}
	return len(buf), true
}

func (c *oscStream) Err() error {
	println("error with osc at stat: ", c.stat)
	return nil
}
