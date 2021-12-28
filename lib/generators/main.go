package generators

import "math"

// create stream which will produce infinite osciator tone with the given frequency
// use other wrappers of this package to change amplitude or add time limit
// sampleRate must be at least two times grater then frequency, otherwise this function will return an error
type oscStream struct {
	stat    float64 // progress from 0 to 1
	delta   float64 // space between two calculation
	oscFunc func(stat float64, delta float64) float64
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
