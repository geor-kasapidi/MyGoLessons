package main

import (
	"fmt"
	"io"
	"math"
)

func DrawSurfaceTo(out io.Writer) {
	const (
		width, height = 600, 320
		cells         = 100
		xyrange       = 30.0
		xyscale       = width / 2 / xyrange
		zscale        = height * 0.4
		angle         = math.Pi / 6
	)

	var (
		sinA, cosA = math.Sincos(angle)
	)

	f := func(x, y float64) float64 {
		return math.Pow(y/xyrange, 2) - math.Pow(x/xyrange, 2)
		// return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
		// return math.Cos(r) / r
	}

	corner := func(i, j int) (float64, float64) {
		x := xyrange * (float64(i)/cells - 0.5)
		y := xyrange * (float64(j)/cells - 0.5)

		z := f(x, y)

		sx := width/2 + (x-y)*cosA*xyscale
		sy := height/2 + (x+y)*sinA*xyscale - z*zscale

		return sx, sy
	}

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>\n", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(out, "</svg>")
}
