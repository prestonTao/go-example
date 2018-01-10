package main

import (
	"fmt"
	"github.com/pierrre/imageserver"
	"github.com/pierrre/imageserver/graphicsmagick"
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

/*
	给图片瘦身
	图片不会改变尺寸，但是会将其图片质量下降到85%（肉眼不可辩）
	转换后的图片大小大概最起码会减少三分之二左右。
	GrahpicsMagick：
	find . -iname “*.jpg” -exec gm convert -strip +profile “*” -quality 85 {} {} \;
	​
	ImageMagick：
	find . -iname “*.jpg” -exec convert -strip +profile “*” -quality 85 {} {} \;
*/
func slimming() {

}
