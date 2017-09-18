/*
Package pbar is used to show progressbar in the terminal

Usage

Sample usage of this package is specified in the exmaple directory

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

Display a progressbar with just 3 steps.

 + Initialize the Bar with pbar.NewBar
 + (Optionally) Set the total amount and unit to be shown
 + Update the pbar progress using Update method on the Bar object
   progress in percentage
*/
package pbar

import (
	"fmt"
	"strings"
	"time"
)

// Bar is a struct which contains all the information to draw a bar
// in the terminal stdout
type Bar struct {
	progress  int64
	control   chan int64
	message   string
	total     float64
	units     string
	inception time.Time
}

// NewBar returns a pointer to a new Bar struct
func NewBar(message string) *Bar {
	return &Bar{message: message, control: make(chan int64), units: "u/s", total: 100.0}
}

// SetSpeedInfo method is used to set the speed info like the total and units in string
func (b *Bar) SetSpeedInfo(total float64, units string) {
	b.total = total
	b.units = units
}

// Start method invokes a go routine which displays the bar with info like speed, percentage etc.,
func (b *Bar) Start() {
	b.inception = time.Now()
	go func(bar *Bar) {
		for bar.progress < 100 {
			bar.progress = <-bar.control
			if bar.progress > 100 {
				bar.progress = 100
			}
			speed := ((float64(b.progress) / 100.0) * (b.total / time.Now().Sub(b.inception).Seconds()))
			fmt.Printf("%c[2K\r%s\t\t [%s%s] %.2f %s %d%%",
				27,
				bar.message,
				strings.Repeat("#", int(bar.progress/5)),
				strings.Repeat(" ", 20-int(bar.progress/5)),
				speed,
				bar.units,
				bar.progress,
			)
		}
		fmt.Print("\n")
	}(b)
}

// Update method is used to update the progress of the bar
func (b *Bar) Update(progress int64) {
	b.control <- progress
}
