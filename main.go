package main

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"log"
	"math"
	"time"
)

func main() {
	sr := beep.SampleRate(44100)
	if err := speaker.Init(sr, sr.N(time.Second/10)); err != nil {
		log.Fatal(err)
	}

	streamer := &Sinus{sampleRate: sr}

	done := make(chan bool)
	speaker.Play(
		beep.Seq(
			beep.Take(sr.N(5*time.Second), streamer),
			beep.Callback(func() { done <- true }),
		),
	)
	<-done
}

type Sinus struct {
	sampleRate beep.SampleRate
	position   uint64
}

func (sinus *Sinus) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		l := math.Sin(float64(sinus.position+uint64(i)) / 20)
		r := math.Cos(float64(sinus.position+uint64(i)) / 30)
		samples[i][0] = l
		samples[i][1] = r
	}
	sinus.position += uint64(len(samples))
	log.Println("position: ", sinus.position)
	return len(samples), true
}
func (sinus *Sinus) Err() error { return nil }
