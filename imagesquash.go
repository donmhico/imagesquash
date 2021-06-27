package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Adjust accordingly.
var quality = 60

// Full path of the directory to watch for new files.
// E.g.
// var watchDir = "C:\\Users\\donmhico\\Pictures"
var watchDir = "FULL DIRECTORY PATH TO WATCH"

// Full path where the compressed files are generated.
// You may have to create the folder first.
// E.g.
// var watchDir = "C:\\Users\\donmhico\\Documents\\compressed"
var outputDir = "FULL DIRECTORY PATH FOR COMPRESSED FILES"

var timer *time.Timer

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Listen to new files in directory.
				if event.Op&fsnotify.Write == fsnotify.Write {
					addQueue(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(watchDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// `WRITE` event fires off multiple times when copying files or
// when saving files in the `watchDir`.
// For the meantime, I workaround this by assuming that the operation
// always complete in a second.
//
// TODO - Make a more intelligent workaround.
func addQueue(newFilePath string) {
	if timer != nil {
		return
	}

	// Change the duration if working on bigger files.
	timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C

		// Release timer
		timer = nil

		compressFile(newFilePath)
	}()
}

func compressFile(newFilePath string) {
	log.Println("file to be compressed", newFilePath)

	// Only get the filename from `newFilePath`.
	split := strings.Split(newFilePath, `\`)
	targetFile := split[len(split)-1]
	targetPath := outputDir + "/" + targetFile

	_, err := os.Stat(targetPath)

	// `err` returns `nil` if `path` exists.
	if err == nil {
		log.Printf("Target path: \"%v\" already exists!", targetPath)
		return
	}

	// Open the image file.
	imageFilePath := watchDir + "/" + targetFile
	imageFile, err := os.Open(imageFilePath)

	// Check if image does not exists.
	if os.IsNotExist(err) {
		imageFile.Close()
		log.Printf("File %v not found.", imageFilePath)
		return
	}

	// Create the target image.
	targetImage, err := os.Create(targetPath)

	if err != nil {
		targetImage.Close()
		log.Println("Error when creating target image")
		log.Print(err)
		return
	}

	image, _, err := image.Decode(imageFile)

	if err != nil {
		targetImage.Close()
		imageFile.Close()
		log.Println("Error when decoding image")
		log.Print(err)
		return
	}

	errEncode := jpeg.Encode(targetImage, image, &jpeg.Options{Quality: quality})

	if errEncode != nil {
		targetImage.Close()
		imageFile.Close()
		log.Println("Error when compressing image")
		log.Print(err)
		return
	}

	// Release resources.
	targetImage.Close()
	imageFile.Close()

	log.Printf("Compressed successfully: %v", targetPath)
}