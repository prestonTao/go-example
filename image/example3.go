package main

import (
	"fmt"
	"graphics"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", doImageHandler)
	http.ListenAndServe("127.0.0.1:6789", nil)
}

func doImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%q\n", strings.Split(r.URL.Path, "/"))
	url := strings.Split(r.URL.Path, "/")
	if len(url) != 3 {
		return
	}
	newdx, uerr := strconv.Atoi(url[1])
	if uerr != nil {
		log.Fatal(uerr)
	}
	src, err := LoadImage(url[2])
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	if err != nil {
		log.Fatal(err)
	}
	// 缩略图的大小
	dst := image.NewRGBA(image.Rect(0, 0, newdx, newdx*dy/dx))
	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	header := w.Header()
	header.Add("Content-Type", "image/jpeg")

	png.Encode(w, dst)
}

// Load Image decodes an image from a file of image.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}


package mas
 
import (
    "code.google.com/p/graphics-go/graphics"
    "fmt"
    z "github.com/nutzam/zgo"
    "image"
)
 
func (ma *Master) ActiveImage(pobj string) error {
 
    // 文件绝对路径
    var path string = pobj
 
    // 保留源图Image结构
    var img image.Image
 
    // 图片类型
    typef := z.FileType(path)
 
    // 按照图片格式加载图片
    switch typef {
    // JPEG
    case "jpeg":
        // ImageJPEG
        img = z.ImageJPEG(path)
    // JPG
    case "jpg":
        // ImageJPEG
        img = z.ImageJPEG(path)
    // PNG
    case "png":
        // ImagePNG
        img = z.ImagePNG(path)
    }
 
    // 判断加载原图片是否成功
    if img == nil {
        // 返回错误
        return fmt.Errorf("active image decode exception ...")
    }
 
    // -------------------------------------------------------- //
 
    // 获取屏幕数量
    moniSize := ma.NodeConf.MonitorSize(ma)
 
    // 获取屏幕分辨率
    width := ma.NodeConf.Resolution.Width
    height := ma.NodeConf.Resolution.Height
 
    // 获取素材平均值
    widthMoni := img.Bounds().Dx() / moniSize.Col
    heightMoni := img.Bounds().Dy() / moniSize.Row
 
    // -------------------------------------------------------- //
 
    // 遍历屏幕,切割图片
    for _, monis := range ma.NodeConf.Layout {
 
        // 遍历节点屏幕
        for _, moni := range monis {
 
            // 获取图片
            row := moni.Display.Row
            col := moni.Display.Col
 
            // 生成目标背景图
            backgroundSrc := z.ImageRGBA(widthMoni, heightMoni)
 
            // 生成目标图
            z.ImageDrawRGBA(backgroundSrc, img, (col-1)*widthMoni, (row-1)*heightMoni)
 
            // 生成最终背景图
            background := z.ImageRGBA(width, height)
 
            // 产生最终图
            graphics.Scale(background, backgroundSrc)
 
            // 按照图片格式保存图片
            switch typef {
            // JPEG
            case "jpeg":
                // ImageEncodeJPEG
                z.ImageEncodeJPEG(fmt.Sprintf("%s.pic_result/%d_%d.%s", path, col, row, typef), background)
            // JPG
            case "jpg":
                // ImageEncodeJPEG
                z.ImageEncodeJPEG(fmt.Sprintf("%s.pic_result/%d_%d.%s", path, col, row, typef), background)
            // PNG
            case "png":
                // ImageEncodePNG
                z.ImageEncodePNG(fmt.Sprintf("%s.pic_result/%d_%d.%s", path, col, row, typef), background)
            }
 
        }
 
    }
 
    // 返回
    return nil
 
}

package z
 
import (
    "image"
    "image/draw"
    "image/jpeg"
    "image/png"
    "os"
)
 
// 读取JPEG图片返回image.Image对象
func ImageJPEG(ph string) image.Image {
    // 打开图片文件
    f, fileErr := os.Open(ph)
    if fileErr != nil {
        return nil
    }
    // 退出时关闭文件
    defer f.Close()
    // 解码
    j, jErr := jpeg.Decode(f)
    if jErr != nil {
        return nil
    }
    // 返回解码后的图片
    return j
}
 
// 读取PNG图片返回image.Image对象
func ImagePNG(ph string) image.Image {
    // 打开图片文件
    f, fileErr := os.Open(ph)
    if fileErr != nil {
        return nil
    }
    // 退出时关闭文件
    defer f.Close()
    // 解码
    p, pErr := png.Decode(f)
    if pErr != nil {
        return nil
    }
    // 返回解码后的图片
    return p
}
 
// 按照分辨率创建一张空白图片对象
func ImageRGBA(width, height int) *image.RGBA {
    // 建立图像,image.Rect(最小X,最小Y,最大X,最小Y)
    return image.NewRGBA(image.Rect(0, 0, width, height))
}
 
// 将图片绘制到图片
func ImageDrawRGBA(img *image.RGBA, imgcode image.Image, x, y int) {
    // 绘制图像
    // image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
    // image.ZP ZP is the zero Point
    // image.Pt Pt is shorthand for Point{X, Y}
    draw.Draw(img, img.Bounds(), imgcode, image.Pt(x, y), draw.Over)
}
 
// JPEG将编码生成图片
// 选择编码参数,质量范围从1到100,更高的是更好 &jpeg.Options{90}
func ImageEncodeJPEG(ph string, img image.Image) error {
    // 确保文件父目录存在
    FcheckParents(ph)
    // 打开文件等待写入
    f := FileW(ph)
    // 保证文件正常关闭
    defer f.Close()
    // 写入文件
    return jpeg.Encode(f, img, &jpeg.Options{100})
}
 
// PNG将编码生成图片
func ImageEncodePNG(ph string, img image.Image) error {
    // 确保文件父目录存在
    FcheckParents(ph)
    // 打开文件等待写入
    f := FileW(ph)
    // 保证文件正常关闭
    defer f.Close()
    // 写入文件
    return png.Encode(f, img)
}