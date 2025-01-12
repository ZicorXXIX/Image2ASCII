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
