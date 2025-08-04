package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"image"
	"image/color"
	"image/png"
	"math/rand"

	"fyne.io/fyne/v2"
)

func GenerateAvatar(username string) fyne.Resource {
	const (
		size      = 128 // Image size in pixels
		gridSize  = 5   // 5x5 grid
		blockSize = size / gridSize
	)

	hash := md5.Sum([]byte(username))
	seed := int64(binary.LittleEndian.Uint64(hash[:8]))
	r := rand.New(rand.NewSource(seed))

	// Primary color
	baseColor := color.NRGBA{
		R: uint8(r.Intn(200) + 30),
		G: uint8(r.Intn(200) + 30),
		B: uint8(r.Intn(200) + 30),
		A: 255,
	}

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	bgColor := color.NRGBA{R: 240, G: 240, B: 240, A: 255} // Light background

	for y := range make([]struct{}, size) {
		for x := range make([]struct{}, size) {
			img.Set(x, y, bgColor)
		}
	}

	for y := range make([]struct{}, gridSize) {
		for x := range make([]struct{}, (gridSize+1)/2) {
			if r.Intn(2) == 0 {
				continue
			}

			drawBlock(img, x, y, blockSize, baseColor)
			drawBlock(img, gridSize-1-x, y, blockSize, baseColor) // Mirror
		}
	}

	// Encode to PNG
	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	return fyne.NewStaticResource(username+".png", buf.Bytes())
}

func drawBlock(img *image.RGBA, x, y, size int, col color.Color) {
	startX := x * size
	startY := y * size
	for dy := range size {
		for dx := range size {
			img.Set(startX+dx, startY+dy, col)
		}
	}
}
