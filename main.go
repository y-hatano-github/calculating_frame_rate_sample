package main

import (
	"fmt"
	"math/rand"
	"time"
)

type FPS struct {
	startTime int64
	endTime   int64
	fps       int64
	gap       int64
	frame     int64
}

func NewFPS(fps int64) *FPS {
	return &FPS{
		fps:       fps,
		startTime: 0,
		endTime:   0,
		frame:     0,
	}
}

func (f *FPS) Update() {
	if f.frame == 0 {
		f.startTime = time.Now().UnixNano() / int64(time.Millisecond)
	}
	if f.frame == f.fps {
		f.frame = 0
		f.startTime = time.Now().UnixNano() / int64(time.Millisecond)
	}
	f.frame++
}

func (f *FPS) Wait() {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	wait := f.frame*1000/f.fps - (now - f.startTime)

	for time.Now().UnixNano()/int64(time.Millisecond)-now < wait {
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fps := NewFPS(60) // FPS 60

	var diffTotal int64
	for i := 0; i < 30; i++ {
		fmt.Println("-----------")
		start := time.Now().UnixNano() / int64(time.Millisecond)
		fps.Update()
		// r = processing time
		r := rand.Intn(20)
		time.Sleep(time.Duration(r) * time.Millisecond)
		fps.Wait()
		end := time.Now().UnixNano() / int64(time.Millisecond)

		fmt.Println("start-time:", start)
		fmt.Println("end-time  :", end)
		fmt.Println("proc-time :", r)
		fmt.Println("time-diff :", end-start)

		diffTotal += end - start
	}
	fmt.Println("-----------")
	fmt.Println("diff-avg. :", diffTotal/30)
}
