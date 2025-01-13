package convert

import (
	"image"

	"github.com/ZicorXXIX/Image2ASCII/pkg/ascii"
)

type ImageConverter struct {
    Ascii   ascii.PixelConverter
}

type Converter interface {
    ImageToAsciiMatrix(image.Image) [][]ascii.Pixel 
    ImageToAsciiSlice(image.Image) []string 
    ImageToAsciiString(image.Image) string 
    ImageFileToAsciiString(string) string 
}

type Options struct {
    Width    int
    Height   int
	Pixels   []byte
	Reversed bool
	Colored  bool
}

var DefaultOptions = Options{
    Width:    80,
    Height:   80,
    Pixels:   []byte(" .,:;i1tfLCG08@"),
    Reversed: false,
    Colored:  true,
}

