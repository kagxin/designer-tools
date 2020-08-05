package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var count = 0
	var im image.Config
	var decode func(r io.Reader) (image.Config, error)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), "png") {
			decode = png.DecodeConfig
		} else if strings.HasSuffix(f.Name(), "jpg") {
			decode = jpeg.DecodeConfig
		} else {
			continue
		}
		func() {
			if reader, err := os.Open(f.Name()); err == nil {
				defer reader.Close()
				im, err = decode(reader)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %v\n", f.Name(), err)
					return
				}
				fileName := fmt.Sprintf("%s-%d.txt", strings.Split(f.Name(), ".")[0], im.Height)
				fmt.Println(fileName)
				if f, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666); err != nil {
					defer f.Close()
				}
				count++
			} else {
				fmt.Printf("Impossible to open the file: %s, err:%#v\n", f.Name(), err)
			}
		}()
	}
	fmt.Printf("完成，共创建%d个txt\n", count)
	var ts string
	fmt.Scanf("%s", &ts)
}
