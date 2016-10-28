package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/utahta/uststat"
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
	if name == "" {
		flag.Usage()
		return nil
	}

	c, err := uststat.New()
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
