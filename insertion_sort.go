// Create GIF visualization of insertion sort
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
)

func makeImage(bars []int) *image.Paletted {
	palette := []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	img := image.NewPaletted(image.Rect(0, 0, len(bars)*10, len(bars)*10), palette)
	for y := 0; y < len(bars)*10; y++ {
		for x := 0; x < len(bars)*10; x++ {
			if bars[int(x/10)]*10 >= len(bars)*10-y {
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
	images := []*image.Paletted{}
	delays := []int{}
	bars := []int{8, 3, 10, 41, 37, 27, 30, 23, 40, 34, 45, 13, 42, 14, 21, 22, 24, 16, 33, 36, 7, 35, 1, 50, 2, 43, 49, 20, 31, 39, 6, 5, 15, 17, 11, 18, 25, 12, 32, 44, 46, 48, 47, 38, 19, 28, 9, 29, 26, 4}
	for i := 1; i < len(bars); i++ {
		key := bars[i]
		j := i - 1
		for j >= 0 && key < bars[j] {
			bars[j+1] = bars[j]
			j -= 1
		}
		bars[j+1] = key
		images = append(images, makeImage(bars))
		delays = append(delays, 5)
	}
	delays[len(delays)-1] = 250
	f, err := os.Create("insertionsort.gif")
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
