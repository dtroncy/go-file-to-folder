package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
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
			fmt.Println("Traitement du fichier : " + element.Name())
			fmt.Println("Fichier sans extension : " + strings.TrimSuffix(element.Name(), path.Ext(element.Name())))
			os.Mkdir(strings.TrimSuffix(element.Name(), path.Ext(element.Name())), 666)
		}

	}
}
