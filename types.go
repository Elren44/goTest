package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	x := make(map[string][]string)

	folders, err := ioutil.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}
	for _, folder := range folders {
		files, err := ioutil.ReadDir("../" + folder.Name())
		if err != nil {
			log.Fatal(err)
		}
		var values []string
		for _, file := range files {
			values = append(values, file.Name())
		}
		x[folder.Name()] = values
	}
	for v, a := range x {
		fmt.Printf("Key: %s - value: %s\n", v, a)
	}

}
