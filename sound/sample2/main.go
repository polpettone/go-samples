package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type EngineOne struct {
	streamer beep.Streamer
}

func (e EngineOne) PlaySound() {

	fmt.Println("PlaySound")

	done := make(chan bool)
	speaker.Play(beep.Seq(e.streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func setup() beep.Streamer {
	f, err := os.Open("sample.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	fast := beep.ResampleRatio(4, 5, streamer)

	return fast
}

func main() {

	e := EngineOne{}
	e.streamer = setup()

	e.PlaySound()

	e.PlaySound()

}
