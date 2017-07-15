package main

import (
	"log"

	"github.com/gabrtv/logbomb/logbomb"
)

func main() {
	log.Println("initializing logbomb...")
	logBomb, err := logbomb.NewLogBomb()
	if err != nil {
		log.Fatalf("error building a log bomb: %s", err)
	}
	log.Println("starting detonation...")
	if err := logBomb.Detonate(); err != nil {
		log.Fatalf("error detonating the log bomb: %s", err)
	}
	log.Println("detonation complete!")
}
