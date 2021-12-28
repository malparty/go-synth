package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/malparty/go-synth/lib/generators"

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
	speaker.Init(beep.SampleRate(48000), 4800)
	// s, err := generators.SinTone(beep.SampleRate(48000), f)
	// if err != nil {
	// 	panic(err)
	// }

	s2, err := generators.NewGenerator(beep.SampleRate(48000), f, generators.SawFunc)
	if err != nil {
		panic(err)
	}
	// speaker.Play(s)
	speaker.Play(s2.GetOsc())

	reader := bufio.NewReader(os.Stdin)

	for {
		userInput, _ := reader.ReadString('\n')

		switch userInput {
		case "k\n":
			s2.SetFreq(s2.GetFreq() + 10.0)
		case "j\n":
			s2.SetFreq(s2.GetFreq() - 10.0)
		case "q\n":
			return
		}

	}
}
