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
	hash := md5.Sum([]byte(username))
	seed := int64(binary.LittleEndian.Uint64(hash[:8]))
	r := rand.New(rand.NewSource(seed))

	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	baseColor := color.NRGBA{
		R: uint8(r.Intn(200) + 30),
		G: uint8(r.Intn(200) + 30),
		B: uint8(r.Intn(200) + 30),
		A: 255,
	}

	for y := range make([]int, 64) {
		for x := range make([]int, 64) {
			if (x/10+y/10)%2 == 0 {
				img.Set(x, y, baseColor)
			} else {
				img.Set(x, y, color.NRGBA{
					R: 255 - baseColor.R,
					G: 255 - baseColor.G,
					B: 255 - baseColor.B,
					A: 255,
				})
			}
		}
	}

	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	return fyne.NewStaticResource(username+".png", buf.Bytes())
}
