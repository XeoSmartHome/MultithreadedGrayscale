package main

import "sync"

func worker(images <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for img := range images {
		processImage(img)
	}
}
