package main

import (
	"fmt"
    "flag"

	"github.com/ZicorXXIX/Image2ASCII/pkg/convert"
)



func main() {
    var path string
	var width int
	var height int
	var pixels string
	var reversed bool
	var colored bool

    DefaultOptions := convert.DefaultOptions

    flag.StringVar(&path, "path", "pikachu.png", "Path to the image")
	flag.IntVar(&width, "width", DefaultOptions.Width, "Width of the output ASCII art")
	flag.IntVar(&height, "height", DefaultOptions.Height, "Height of the output ASCII art")
	flag.StringVar(&pixels, "pixels", string(DefaultOptions.Pixels),"Characters used for ASCII art")
	flag.BoolVar(&reversed, "reversed", DefaultOptions.Reversed, "Reverse the brightness levels")
	flag.BoolVar(&colored, "colored", DefaultOptions.Colored, "Enable colored ASCII art")
	flag.Parse()

    options := convert.Options{
		Width:    width,
		Height:   height,
		Pixels:   []byte(pixels),
		Reversed: reversed,
		Colored:  colored,
	}

    converter:= convert.NewImageConverter()
    matrix := converter.ImageFileToAsciiString(path, options)
    fmt.Println(matrix)
}

