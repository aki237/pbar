package main

import (
	"time"

	"math/rand"

	"github.com/aki237/pbar"
)

func main() {
	bar := pbar.NewBar("something.iso")
	bar.SetSpeedInfo(30, "MB/s")
	bar.Start()
	rand.Seed(time.Now().Unix())
	progress := int64(0)
	for progress < 100 {
		progress += int64(rand.Intn(10))
		bar.Update(progress)
		x := time.Duration(600 + rand.Intn(100))
		time.Sleep(x * time.Millisecond)
	}
}
