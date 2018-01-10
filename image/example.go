package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	catImage_2()
}

/*
	裁剪图片，将图片左上角100*100像素裁剪下来
*/
func cutImg() {
	f1, err := os.Open("./testFile/f.jpg") //file to scale
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	m, _, _ := image.Decode(f1) // 图片文件解码
	rgbImg := m.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(100, 100, 200, 200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1

	bounds := subImg.Bounds()

	m2 := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m2, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	draw.Draw(m2, bounds, subImg, image.ZP, draw.Src)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("./testFile/newcut.jpg")
	jpeg.Encode(imgw, m2, &jpeg.Options{100})

	defer imgw.Close()
}

//裁剪图片
func catImage_2() {
	//原始图片是sam.jpg
	imgb, _ := os.Open("./testFile/f.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	//创建一个长300，高300大小的图片
	b := image.Rect(0, 0, 300, 300)
	fmt.Println(b.Dx(), b.Dy())
	// draw.Draw(, r, src, sp, op)
	m := image.NewNRGBA(b)

	b2 := image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy())
	offset := image.Pt(100, 100)
	draw.Draw(m, b2.Sub(offset), img, image.ZP, draw.Src)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("./testFile/newcut.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()
	fmt.Println("水印添加结束,请查看new.jpg图片...")
}
