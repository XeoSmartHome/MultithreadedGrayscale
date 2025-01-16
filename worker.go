package main

import (
	"log"
	"sync"
)

func worker(id int, images <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for img := range images {
		log.Printf("Worker[%d] processing image %s\n", id, img)
		processImage(img)
		log.Printf("Worker[%d] processed image %s\n", id, img)
	}
}
