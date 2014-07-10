package main

import (
	// "code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	// "image/color"
	"image/draw"
	"image/jpeg"
	// "image/png"
	"os"
)

func main() {
	catImage()
}

//裁剪图片
//
func catImage() {
	//原始图片是sam.jpg
	imgb, err := os.Open("./testFile/5.jpg")
	if err != nil {
		fmt.Println("图片路径不正确")
	}
	img, err := jpeg.Decode(imgb)
	if err != nil {
		fmt.Println("图片格式不正确", err.Error())
	}
	defer imgb.Close()

	//创建一个长100，高200大小的图片
	b := image.Rect(0, 0, 442, 442)
	fmt.Println(b.Dx(), b.Dy())
	// draw.Draw(, r, src, sp, op)
	m := image.NewNRGBA(b)

	offset := image.Pt(10, 10)
	draw.Draw(m, b.Add(offset), img, image.ZP, draw.Src)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("./testFile/new4.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()
	fmt.Println("水印添加结束,请查看new.jpg图片...")
}
