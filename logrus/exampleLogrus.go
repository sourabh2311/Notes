package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	f, err := os.OpenFile("/Users/sourabhaggarwal/Desktop/testlogrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f) // can put os.Stdout
	log.Info("A walrus appears")
	log.Info("Checking this again")
	log.Warn("Warning")
}
