package effects

import (
	"github.com/malparty/go-synth/lib/generators"
)

type Limiter struct {
	Rate float64
}

func (l *Limiter) GetLimiterFunc() generators.GeneratorFunction {

	return func(stat float64, _ float64) float64 {
		return stat * l.Rate / 100
	}
}
