package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
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
		cycles, _ := strconv.Atoi(r.FormValue("cycles"))
		res, _ := strconv.Atoi(r.FormValue("res"))
		size, _ := strconv.Atoi(r.FormValue("size"))
		nframes, _ := strconv.Atoi(r.FormValue("nframes"))
		delay, _ := strconv.Atoi(r.FormValue("delay"))
		arr := map[string]int{
			"cycles":  cycles,
			"res":     res,
			"size":    size,
			"nframes": nframes,
			"delay":   delay,
		}
		lissajous(w, arr)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8003", nil))
}

func lissajous(out io.Writer, arr map[string]int) {
	var (
		cycles  = 5
		res     = 0.001
		size    = 200
		nframes = 64
		delay   = 8
	)
	if arr["cycles"] > 0 {
		cycles = arr["cycles"]
	}
	if arr["res"] > 0 {
		res = float64(arr["res"])
	}
	if arr["size"] > 0 {
		size = arr["size"]
	}
	if arr["nframes"] > 0 {
		nframes = arr["nframes"]
	}
	if arr["delay"] > 0 {
		delay = arr["delay"]
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
