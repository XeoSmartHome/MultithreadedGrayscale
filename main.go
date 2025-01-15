package main

import (
	"log"
	"os"
	"sync"
	"time"
)

const ThreadsNumber = 4
const InputDirectory = "input"
const OutputDirectory = "output"

func main() {
	log.Println("Starting image processing using", ThreadsNumber, "threads")

	startTime := time.Now()

	var imageQueue = make(chan string, 100)
	var waitGroup = sync.WaitGroup{}

	for range ThreadsNumber {
		waitGroup.Add(1)
		go worker(imageQueue, &waitGroup)
	}

	inputFolder, err := os.ReadDir(InputDirectory)

	if err != nil {
		log.Fatal("Error reading folder", err)
	}

	for _, file := range inputFolder {
		if file.IsDir() {
			continue
		}

		imageQueue <- file.Name()
	}

	close(imageQueue)

	waitGroup.Wait()

	endTime := time.Now()
	log.Println("Image converted in", endTime.Sub(startTime))
}
