package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	blueIndex  = 3
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			fmt.Sprint(err)
			return
		}
		lissajous(w, u.Query())
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8003", nil))
}

func lissajous(out io.Writer, u url.Values) {
	cycles = 5
	res = 0.001
	size = 200
	nframes = 64
	delay = 8

	if cycles, ok := strconv.Atoi(u.Get("cycles")); ok != nil {
	} else {
		cycles := 5
	}
	if res, ok := strconv.Atoi(u.Get("res")); ok != nil {
	} else {
		res := 0.001
	}
	if size, ok := strconv.Atoi(u.Get("size")); ok != nil {
	} else {
		size := 5
	}
	if nframes, ok := strconv.Atoi(u.Get("nframes")); ok != nil {
	} else {
		nframes := 200
	}
	if delay, ok := strconv.Atoi(u.Get("delay")); ok != nil {
	} else {
		delay := 8
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		colorIndex := uint8(blackIndex)
		switch {
		case i%9 <= 2:
			colorIndex = greenIndex
		case i%9 >= 3 && i%9 <= 5:
			colorIndex = redIndex
		case i%9 >= 6 && i%9 <= 8:
			colorIndex = blueIndex
		}
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
