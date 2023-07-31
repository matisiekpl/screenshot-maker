package main

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"io/fs"
	"os"
)

func getFont(size uint) font.Face {
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		logrus.Panic("cannot parse font")
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: float64(size),
	})
	return face
}

func drawText(screenshot Screenshot, img image.Image, content string, c color.Color) image.Image {
	dc := gg.NewContextForImage(img)
	dc.SetColor(c)

	var size uint = 96
	var w float64
	dc.SetFontFace(getFont(size))
	w, _ = dc.MeasureString(content)
	i := 0
	for uint(w) > (screenshot.Width - 100) {
		dc.SetFontFace(getFont(size))
		w, _ = dc.MeasureString(content)
		size--
		i++
		if i > 100 {
			logrus.Panic("cannot find optimal text size")
		}
	}

	dc.DrawStringAnchored(content, float64(screenshot.Width/2), float64(screenshot.MockupOffset/2), 0.5, 0.5)
	return dc.Image()
}

func loadImage(filename string, filesystem fs.FS) image.Image {
	file, err := filesystem.Open(filename)
	if err != nil {
		logrus.Panicf("cannot open image: %s", filename)
	}
	decoded, _, err := image.Decode(file)
	if err != nil {
		logrus.Panicf("cannot decode image: %s", filename)
	}
	defer file.Close()
	return decoded
}

func newEmptyImage(width, height uint, c color.Color) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: int(width), Y: int(height)}})
	for x := uint(0); x < width; x++ {
		for y := uint(0); y < height; y++ {
			img.Set(int(x), int(y), c)
		}
	}
	return img
}

func resizeImage(original image.Image, width, height uint) image.Image {
	return resize.Resize(width, height, original, resize.Lanczos3)
}

func merge(first, second image.Image, x int, y uint) *image.RGBA {
	offset := image.Pt(x, int(y))
	b := first.Bounds()
	output := image.NewRGBA(b)
	draw.Draw(output, b, first, image.Point{}, draw.Src)
	draw.Draw(output, second.Bounds().Add(offset), second, image.Point{}, draw.Over)
	return output
}

func (screenshot Screenshot) render(photoFilename, content string, backgroundColor, textColor color.Color) image.Image {
	board := newEmptyImage(screenshot.Width, screenshot.Height, backgroundColor)
	mockupOriginal := loadImage("devices/"+screenshot.Filename, devices)
	mockup := resizeImage(mockupOriginal, screenshot.Width-50, uint(mockupOriginal.Bounds().Dx()/mockupOriginal.Bounds().Dy())*screenshot.Width)

	cwd, _ := os.Getwd()
	photoOriginal := loadImage(photoFilename, os.DirFS(cwd))
	photo := resizeImage(photoOriginal, screenshot.Width-screenshot.PhotoWidth, uint(photoOriginal.Bounds().Dx()/photoOriginal.Bounds().Dy())*screenshot.Width)

	x := merge(board, photo, (board.Bounds().Dx()-photo.Bounds().Dx())/2, screenshot.PhotoOffset)
	x = merge(x, mockup, (board.Bounds().Dx()-mockup.Bounds().Dx())/2, screenshot.MockupOffset)
	output := drawText(screenshot, x, content, textColor)
	return output
}
