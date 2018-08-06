package main

import (
	"time"
	"log"
)

func main() {
	bigSloOperation()
}

func bigSloOperation()  {
	defer trace("bigSloOperation()")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func()  {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}