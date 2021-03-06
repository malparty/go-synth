package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/malparty/go-synth/lib"
	"github.com/malparty/go-synth/lib/generators"
	"github.com/malparty/go-synth/lib/generators/effects"
	"github.com/malparty/go-synth/lib/generators/oscillators"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func usage() {
	fmt.Printf("usage: %s freq\n", os.Args[0])
	fmt.Println("where freq must be an integer from 1 to 24000")
	fmt.Println("24000 because samplerate of 48000 is hardcoded")
}
func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	f, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		return
	}
	speaker.Init(beep.SampleRate(lib.SampleRate), 4800)
	// s, err := generators.SinTone(beep.SampleRate(48000), f)
	// if err != nil {
	// 	panic(err)
	// }

	limiter := &effects.Limiter{
		Rate: 20.0,
	}

	reverb := &effects.Reverb{
		MixRate:  100,
		FadeRate: 80,
		DelayMs:  10,
		Freq:     float64(f),
	}

	chainFunction := &generators.ChainGenerator{
		GeneratorFuncs: []generators.GeneratorFunction{
			oscillators.SawFunc,
			limiter.GetLimiterFunc(),
			reverb.GetReverbFunc(),
		},
	}

	s2, err := generators.NewGenerator(beep.SampleRate(48000), f, chainFunction.ChainFunc)
	if err != nil {
		panic(err)
	}
	// speaker.Play(s)
	speaker.Play(s2.GetOsc())

	reader := bufio.NewReader(os.Stdin)

	for {
		userInput, _ := reader.ReadByte()

		switch string(userInput) {
		case "k":
			// s2.SetFreq(s2.GetFreq() + 10.0)
			reverb.SetDelay(reverb.DelayMs - 10)
		case "j":
			// s2.SetFreq(s2.GetFreq() - 10.0)
			reverb.SetDelay(reverb.DelayMs + 10)
		case "q":
			return
		}

	}
}
