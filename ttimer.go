package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/drgrib/ttimer/agent"

	"github.com/drgrib/ttimer/parse"
)

var args struct {
	DURATION string `arg:"positional,required" help:"duration of the timer in the format: 1h2m3s"`
	Quit     bool   `arg:"-q,--quit" help:"make the timer automatically exit after finishing"`
	Help     bool   `arg:"-h,--help" help:"this help screen"`
}

func main() {

	arg.MustParse(&args)

	d, title, err := parse.Args(args.DURATION)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("\nPlease refer to https://github.com/drgrib/ttimer for usage instructions.")
		return
	}

	// start timer
	t := agent.Timer{Title: title}
	t.AutoQuit = args.Quit
	t.Start(d)

	// run UI
	t.CountDown()
}
