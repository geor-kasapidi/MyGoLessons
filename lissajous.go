package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

func DrawLissajousTo(out io.Writer) error {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 2
	)

	var palette = []color.Color{
		color.Black,
		color.RGBA{R: 0xFF, G: 0x0, B: 0x00, A: 1},
		color.RGBA{R: 0x0, G: 0xFF, B: 0x00, A: 1},
		color.RGBA{R: 0x0, G: 0x00, B: 0xFF, A: 1},
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	return gif.EncodeAll(out, &anim)
}
