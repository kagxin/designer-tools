package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var count = 0
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() == "main.go" || f.Name() == "mktxt.exe" {
			continue
		}
		fileName := fmt.Sprintf("%s.txt", strings.Split(f.Name(), ".")[0])
		fmt.Println(fileName)
		os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
		count++
	}
	fmt.Printf("完成，共创建%d个txt\n", count)
	var ts string
	fmt.Scanf("%s", &ts)
}
