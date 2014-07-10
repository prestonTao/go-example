package main

import (
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// JpegImage()
	// readWidthAndHeight()
	// catImage()
	DeflationImage()
}

func JpegImage() {
	//原始图片是sam.jpg
	imgb, err := os.Open("./testFile/2.jpg")
	if err != nil {
		fmt.Println("文件路径错误")
	}
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()
	fmt.Println("dx", img.Bounds().Dx(), "dy", img.Bounds().Dy())
	bounds := img.Bounds()
	m := image.NewNRGBA(bounds)
	// m := image.NewRGBA(bounds)

	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	draw.Draw(m, bounds, img, image.ZP, draw.Src)
	// draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("./testFile/new2.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("水印添加结束,请查看new.jpg图片...")

}

//裁剪图片
func catImage() {
	//原始图片是sam.jpg
	imgb, _ := os.Open("./testFile/2.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	//创建一个长100，高200大小的图片
	b := image.Rect(0, 0, 442, 442)
	fmt.Println(b.Dx(), b.Dy())
	// draw.Draw(, r, src, sp, op)
	m := image.NewNRGBA(b)

	offset := image.Pt(10, 10)
	draw.Draw(m, b.Add(offset), img, image.ZP, draw.Src)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("./testFile/new2.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()
	fmt.Println("水印添加结束,请查看new.jpg图片...")
}

//缩小图片
func DeflationImage() {
	//原始图片是sam.jpg
	imgb, _ := os.Open("./testFile/1.png")
	img, _ := png.Decode(imgb)
	defer imgb.Close()

	//创建一个长100，高200大小的图片
	b := image.Rect(0, 0, 442, 442)
	fmt.Println(b.Dx(), b.Dy())
	// draw.Draw(, r, src, sp, op)
	// m := image.NewNRGBA(b)

	// offset := image.Pt(10, 10)
	// draw.Draw(m, b.Add(offset), img, image.ZP, draw.Src)
	// draw.Draw(m, b, img, image.ZP, draw.Src)

	// 缩略图的大小
	dst := image.NewRGBA(image.Rect(0, 0, 400, 400))
	err := graphics.Thumbnail(dst, img)
	if err != nil {
		fmt.Println("压缩失败")
	}

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("./testFile/new2.jpg")
	jpeg.Encode(imgw, dst, &jpeg.Options{100})
	defer imgw.Close()
	fmt.Println("水印添加结束,请查看new.jpg图片...")

}

//查看图片宽度和高度
func readWidthAndHeight() {
	pngb, _ := os.Open("./testFile/1.png")
	pngImg, _ := png.Decode(pngb)
	defer pngb.Close()
	fmt.Println("dx", pngImg.Bounds().Dx(), "dy", pngImg.Bounds().Dy())
}

// func test() {
// 	f1, err := os.Open("./testFile/1.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f1.Close()

// 	f2, err := os.Open("./testFile/2.jpg")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f2.Close()

// 	f3, err := os.Create("./testFile/3.gif")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f3.Close()

// 	m1, err := jpeg.Decode(f1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bounds := m1.Bounds()

// 	m2, err := jpeg.Decode(f2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	m := image.NewRGBA(bounds)
// 	white := color.RGBA{255, 255, 255, 255}
// 	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
// 	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
// 	// draw.Draw(m, image.Rect(100, 200, 300, 600), m2, image.Pt(250, 60), draw.Src)

// 	err = jpeg.Encode(f3, m, &jpeg.Options{90})
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("ok\n")
// }
