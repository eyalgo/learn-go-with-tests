package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {

		sleeper.Sleep()
		_, _ = fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	_, _ = fmt.Fprint(out, finalWord)
}
