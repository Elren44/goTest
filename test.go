package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getHtml() string {
	var url string
	fmt.Print("Введите адрес сайта: ")
	fmt.Scan(&url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	htmlText := string(body)
	return htmlText
}

func filePath(path string) {
	path = strings.Replace(path, "\\", "/", -1)
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	file, err := os.Create(path + "hello.txt")
	if err != nil {
		fmt.Println("Unabe lo create: ", err)
	}
	defer file.Close()
	file.WriteString(getHtml())
	fmt.Println(file.Name())
}

func main() {
	var path string

	fmt.Print("Введите путь до папки в которой создать файл: ")
	fmt.Scan(&path)
	filePath(path)
}
