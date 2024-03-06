package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	loc, err := time.LoadLocation("EST")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Init(): current time ", time.Now().In(loc))
}
