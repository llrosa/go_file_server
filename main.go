package main

import (
	"fmt"
	"net/http"
	"os"
	"flag"
)


func main() {
	port := flag.Int("port", 8080, "File server port")
	volume := flag.String("volume", "/tmp/", "File server port")

	flag.Parse()

	fileServerhostname := fmt.Sprintf(":%d", *port)

	fmt.Printf("Starting File Server at %s with volume %s\n", fileServerhostname, *volume)
	createFolderIfNeeded(*volume)
	panic(http.ListenAndServe(fileServerhostname, http.FileServer(http.Dir(*volume))))
}

func createFolderIfNeeded(path string) {
	//If path does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Creating Folder %s\n", path)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
