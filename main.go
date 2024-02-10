package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"tawesoft.co.uk/go/dialog"
)

type config struct {
	msg      string
	duration time.Duration
	// disattach bool // not implemented yet
}

func main() {

	var cfg config
	flag.StringVar(&cfg.msg, "m", "Knock knock", "Message to print")
	flag.DurationVar(&cfg.duration, "d", 1*time.Second, "Duration to wait")
	// flag.BoolVar(&cfg.disattach, "disattach", false, "Disattach the message from the terminal")
	flag.Parse()

	d, err := time.ParseDuration(cfg.duration.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	w := sync.WaitGroup{}
	c := cron.New(cron.WithSeconds())
	w.Add(1)
	c.AddJob("@every "+d.String(), cron.FuncJob(func() {
		defer w.Done()
		// remove the the job from the scheduler
		c.Remove(cron.EntryID(1))
		dialog.Alert(cfg.msg)
	}))

	c.Start()
	w.Wait()
	c.Stop()

}
