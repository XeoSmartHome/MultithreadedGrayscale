package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

const Program = "./multithreaded-filter"

func main() {
	for i := 1; i <= 10; i++ {
		cmd := Program + " " + strconv.Itoa(i)
		start := time.Now()
		_, err := exec.Command("sh", "-c", cmd).Output()
		end := time.Now()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(i, end.Sub(start))
	}
}
