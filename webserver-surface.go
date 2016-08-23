package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 600
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8004", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, bool1 := corner(i+1, j)
			bx, by, bz, bool2 := corner(i, j)
			cx, cy, cz, bool3 := corner(i, j+1)
			dx, dy, dz, bool4 := corner(i+1, j+1)
			z = (az + bz + cz + dz) / 4
			if bool1 || bool2 || bool3 || bool4 {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#ff0000'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	finite := true

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(sx, 0) || math.IsNaN(sx) || math.IsInf(sy, 0) || math.IsNaN(sy) {
		finite = false
	}

	return sx, sy, z, finite
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return r*r/300 - 1
}
