package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	log.Println("Reading root folder content")
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	elements, err := f.Readdir(-1)

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range elements {

		if !element.IsDir() {

			log.Println("This element is not a directory")

			// Get file path
			filePath, err := filepath.Abs(filepath.Dir(element.Name()))
			if err != nil {
				log.Fatal(err)
			}

			// Get filename without extension
			fileNameWithoutExtension := strings.TrimSuffix(element.Name(), path.Ext(element.Name()))

			log.Println("    File name : " + element.Name())
			log.Println("    File name without extension : " + fileNameWithoutExtension)
			log.Println("    File path : " + filePath)

			log.Println("    Folder creation")
			// Folder creation with filename (without extension)
			os.Mkdir(fileNameWithoutExtension, 0770)

			log.Println("    Moving file from " + filePath + "/" + element.Name() + " to " + filePath + "/" + fileNameWithoutExtension + "/" + element.Name())
			// Move file from root folder to newly created folder
			os.Rename(filePath+"/"+element.Name(), filePath+"/"+fileNameWithoutExtension+"/"+element.Name())

		}

	}
}
