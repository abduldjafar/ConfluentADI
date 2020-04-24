package main

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"math"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	windowSize  = kingpin.Flag("size", "Window size as a percentage.").Short('s').Default("5").Float64()
	percentile  = kingpin.Flag("percentile", "Window percentile.").Short('p').Default("90").Float64()
	targetValue = kingpin.Flag("target", "Target value when scaling output.").Short('t').Default("240").Int()
	files       = kingpin.Arg("files", "Images to process.").Required().ExistingFiles()
)

func timed(name string) func() {
	if len(name) > 0 {
		fmt.Printf("%s... ", name)
	}
	start := time.Now()
	return func() {
		fmt.Println(time.Since(start))
	}
}

func ensureGray(im image.Image) (*image.Gray, bool) {
	switch im := im.(type) {
	case *image.Gray:
		return im, true
	default:
		dst := image.NewGray(im.Bounds())
		draw.Draw(dst, im.Bounds(), im, image.ZP, draw.Src)
		return dst, false
	}
}

func histogramPercentile(hist []int, n int, p float64) int {
	if p <= 0.5 {
		m := int(float64(n) * p)
		for v, c := range hist {
			m -= c
			if m <= 0 {
				return v
			}
		}
	} else {
		m := int(float64(n) * (1 - p))
		for v := 255; v >= 0; v-- {
			m -= hist[v]
			if m <= 0 {
				return v
			}
		}
	}
	panic("oops")
}

func columnPercentiles(im *image.Gray, p float64, x, r int) []int {
	b := im.Bounds()
	x0 := x - r
	x1 := x + r + 1
	if x0 < b.Min.X {
		x0 = b.Min.X
	}
	if x1 > b.Max.X {
		x1 = b.Max.X
	}
	y0 := b.Min.Y
	y1 := b.Max.Y
	result := make([]int, b.Dy())
	hist := make([]int, 256)
	n := 0
	for y := y0; y < y0+r; y++ {
		i := im.PixOffset(x0, y)
		for x := x0; x < x1; x++ {
			hist[im.Pix[i]]++
			i++
			n++
		}
	}
	for y := y0 + r; y < y1; y++ {
		yy := y - r - r
		if yy >= 0 {
			i := im.PixOffset(x0, yy)
			for x := x0; x < x1; x++ {
				hist[im.Pix[i]]--
				i++
				n--
			}
		}
		i := im.PixOffset(x0, y)
		for x := x0; x < x1; x++ {
			hist[im.Pix[i]]++
			i++
			n++
		}
		result[y-r] = histogramPercentile(hist, n, p)
	}
	for y := y1; y < y1+r; y++ {
		yy := y - r - r
		i := im.PixOffset(x0, yy)
		for x := x0; x < x1; x++ {
			hist[im.Pix[i]]--
			i++
			n--
		}
		result[y-r] = histogramPercentile(hist, n, p)
	}
	return result
}

func imagePercentile(im *image.Gray, p float64) int {
	hist := make([]int, 256)
	b := im.Bounds()
	n := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		i := im.PixOffset(b.Min.X, y)
		for x := b.Min.X; x < b.Max.X; x++ {
			hist[im.Pix[i]]++
			i++
			n++
		}
	}
	return histogramPercentile(hist, n, p)
}

func processFile(filename string) {
	var done func()

	fmt.Println(filename)

	s := *windowSize / 100
	p := *percentile / 100
	t := float64(*targetValue)

	done = timed("loading input")
	src, err := gg.LoadImage(filename)
	if err != nil {
		log.Fatal(err)
	}
	done()

	done = timed("converting to grayscale")
	im, _ := ensureGray(src)
	dst := image.NewGray(im.Bounds())
	gradient := image.NewGray(im.Bounds())
	level := image.NewGray(im.Bounds())
	done()

	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	size := int(math.Sqrt(float64(w*h))*s + 0.5)

	done = timed("processing image")
	for x := 0; x < w; x++ {
		column := columnPercentiles(im, p, x, size/2)
		for y, a := range column {
			i := im.PixOffset(x, y)
			v := float64(im.Pix[i])
			v = (v / float64(a)) * t
			if v < 0 {
				v = 0
			}
			if v > 255 {
				v = 255
			}
			dst.Pix[i] = uint8(v)
			gradient.Pix[i] = uint8(a)
		}
	}
	done()

	done = timed("leveling image")
	lo := float64(imagePercentile(dst, 0.0001))
	hi := float64(imagePercentile(dst, 0.97))
	// fmt.Println(lo, hi)
	m := 255 / (hi - lo)
	for i, v := range dst.Pix {
		nv := int((float64(v)-lo)*m + 0.5)
		if nv < 0 {
			nv = 0
		}
		if nv > 255 {
			nv = 255
		}
		level.Pix[i] = uint8(nv)
	}
	done()

	done = timed("writing outputs")
	ext := filepath.Ext(filename)
	basename := filename[:len(filename)-len(ext)]
	err = gg.SavePNG(basename+".gray.png", im)
	if err != nil {
		log.Fatal(err)
	}
	err = gg.SavePNG(basename+".grad.png", gradient)
	if err != nil {
		log.Fatal(err)
	}
	err = gg.SavePNG(basename+".rbgg.png", dst)
	if err != nil {
		log.Fatal(err)
	}
	err = gg.SavePNG(basename+".lvld.png", level)
	if err != nil {
		log.Fatal(err)
	}
	done()
}

func main() {
	kingpin.Parse()
	for _, filename := range *files {
		processFile(filename)
	}
}

// copy paste from https://raw.githubusercontent.com/fogleman/rbgg/master/main.go
// and edied for custom function
