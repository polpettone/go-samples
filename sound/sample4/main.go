package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/hajimehoshi/oto"
)

func main() {

	var context *oto.Context
	context, _ = oto.NewContext(decoder.SampleRate, decoder.Channels, 2, 1024)
	player := context.NewPlayer()

	buffer := initSound()
	initSound()

	run(buffer)
}

func run(buffer beep.Buffer) {
	for {
		fmt.Print("Press [ENTER] to fire a gunshot! ")
		fmt.Scanln()
		shot := buffer.Streamer(0, buffer.Len())
		speaker.Play(shot)
	}
}

func initSound() beep.Buffer {

	f, err := os.Open("gunshot.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)

	buffer.Append(streamer)
	streamer.Close()

	return *buffer
}
