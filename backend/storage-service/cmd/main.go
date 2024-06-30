package main

import (
	storage "backend/storage-service/internal"
	"log"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	figure.NewFigure("CloudShareX", "slant", true).Print()
	log.Println("Storage service is starting...")
	storage.StartStorageCheck()
}
