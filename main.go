package main

import (
	"flag"
	"log"

	"github.com/andreykaipov/goobs"
)

var (
	password = flag.String("obs_password", "", "Web-socket password")
)

func main() {
	c, err := goobs.New("localhost:4455", goobs.WithPassword(*password))
	if err != nil {
		log.Fatalf("failed to connect to OBS: %v", err)
	}
	defer c.Disconnect()
	if _, err := c.Outputs.SaveReplayBuffer(); err != nil {
		log.Fatalf("failed to save replay buffer: %v", err)
	}
}

func init() {
	flag.Parse()
}

