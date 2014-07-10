//golang图片式水印处理
//未实现Web上传等代码...
//sam
//2013.6.19
//http://zituo.net
package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//原始图片是sam.jpg
	imgb, _ := os.Open("sam.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("text.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("水印添加结束,请查看new.jpg图片...")
}
