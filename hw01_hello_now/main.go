package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	start := time.Now()
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("current time: " + start.Round(0).String())
	fmt.Println("exact time: " + exactTime.Round(0).String())
}
