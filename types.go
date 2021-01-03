package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	x := make(map[string][]string)
	key := []string{}
	folders, err := ioutil.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}
	for _, folder := range folders {
		files, err := ioutil.ReadDir("../" + folder.Name())
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			//x["golang"] = file.Name()
			key = append(key, file.Name())
		}
		x[folder.Name()] = key
	}
	fmt.Println(x)
}
