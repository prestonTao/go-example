package main

import (
	"fmt"
	"github.com/pierrre/imageserver"
	"github.com/pierrre/imageserver/processor/graphicsmagick"
	"io/ioutil"
	"os"
)

func main() {
	simple1()
}

func simple1() {
	file, _ := os.Open("./testFile/2.jpg")
	defer file.Close()
	fileData, _ := ioutil.ReadAll(file)

	image := &imageserver.Image{
		Format: "jpg",
		Data:   fileData,
	}

	parameters := imageserver.Parameters{
		"graphicsmagick": imageserver.Parameters{
			"width":  100,
			"height": 100,
		},
	}

	processor := &graphicsmagick.GraphicsMagickProcessor{
		Executable: "D:\\test",
	}

	_, err := processor.Process(image, parameters)
	if err != nil {
		fmt.Println("处理错误", err.Error())
	}

	fmt.Println("ok")
}
