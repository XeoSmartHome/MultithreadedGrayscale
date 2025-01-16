package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const DefaultNumberOfThreads = 1
const InputDirectory = "input"
const OutputDirectory = "output"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: multithreaded-filter <number of threads>")
	}

	numberOfThreads, err := strconv.Atoi(os.Args[1])

	if err != nil {
		numberOfThreads = DefaultNumberOfThreads
		log.Printf("Error parsing threads number, using default %d", DefaultNumberOfThreads)
	} else {
		log.Printf("Using %d threads", numberOfThreads)
	}

	startTime := time.Now()

	var imageQueue = make(chan string, 100)
	var waitGroup = sync.WaitGroup{}

	for i := range numberOfThreads {
		waitGroup.Add(1)
		go worker(i, imageQueue, &waitGroup)
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
