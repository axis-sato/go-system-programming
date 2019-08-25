package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) {
	fmt.Printf("start timer: %d seconds. \n", duration)
	<-time.After(duration * time.Second)
	fmt.Println("timer is finished")
}

func main() {
	timer(2)
}
