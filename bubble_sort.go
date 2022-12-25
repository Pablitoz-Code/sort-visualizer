//Create GIF visualization of bubble sort
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func makeImage(bars []int, counter int) *image.Paletted {
	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	img := image.NewPaletted(image.Rect(0, 0, 500, 500), palette)
	for y := 0; y < 500; y++ {
		for x := 0; x < 500; x++ {
			if bars[int(x/10)]*10 >= 500-y {
				img.Set(x, y, color.NRGBA{
					R: uint8(255),
					G: uint8(255),
					B: uint8(255),
					A: 255,
				})
			} else {
				img.Set(x, y, color.NRGBA{
					R: uint8(0),
					G: uint8(0),
					B: uint8(0),
					A: 255,
				})
			}

		}
	}
	return img
}
func main() {
	const width, height = 50, 50
	images := []*image.Paletted{}
	delays := []int{}
	bars := []int{8, 3, 10, 41, 37, 27, 30, 23, 40, 34, 45, 13, 42, 14, 21, 22, 24, 16, 33, 36, 7, 35, 1, 50, 2, 43, 49, 20, 31, 39, 6, 5, 15, 17, 11, 18, 25, 12, 32, 44, 46, 48, 47, 38, 19, 28, 9, 29, 26, 4}
	n := len(bars)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if bars[j] > bars[j+1] {
				bars[j], bars[j+1] = bars[j+1], bars[j]
				images = append(images, makeImage(bars, i*1000+j))
				delays = append(delays, 5)
			}
		}
	}
	delays[len(delays)-1] = 250
	f, err := os.Create("bubblesort.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
