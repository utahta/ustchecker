package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/utahta/ustchecker"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("ERROR: %v", err))
		os.Exit(1)
	}
}

func run() error {
	var name string
	flag.StringVar(&name, "name", "", "Specifies the ustream channel name")
	flag.Parse()

	c, err := ustchecker.New()
	if err != nil {
		return err
	}

	live, err := c.IsLive(name)
	if err != nil {
		return err
	}

	status := "offline"
	if live {
		status = "live"
	}
	fmt.Fprintln(os.Stdout, status)
	return nil
}
