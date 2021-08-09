package hashes

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func toGrayscale(img image.Image) *image.Gray16 {
	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dest.Set(x, y, gray)
		}
	}
	return dest
}

func resizeGray(img *image.Gray16) *image.Gray16 {
	dst := image.NewGray16(image.Rect(0, 0, 9, 8))
	draw.BiLinear.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dst
}

func dHash(img *image.Gray16) uint64 {
	hash := uint64(0)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			y1, _, _, _ := img.At(x, y).RGBA()
			y2, _, _, _ := img.At(x+1, y).RGBA()
			if y1 < y2 {
				hash += (1 << (63 - y*8 - x))
			}
		}
	}
	return hash
}

func CalcDHash(path string) uint64 {
	j, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer j.Close()

	img, err := png.Decode(j)

	if err != nil {
		panic(err)
	}

	gray_img := toGrayscale(img)

	resized_gray_img := resizeGray(gray_img)

	hash := dHash((resized_gray_img))
	return hash
}
