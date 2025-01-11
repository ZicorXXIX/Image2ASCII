package convert

import (
	"bytes"
	"errors"
	"image"
	"image/color"

	"github.com/ZicorXXIX/Image2ASCII/pkg/ascii"
)

type ImageConverter struct {
    img     image.Image
}

func NewImageConverter(imageFileName string) *ImageConverter {
    i, err := OpenImageFile(imageFileName)
    if err != nil {
        errors.New("Error Opening file")
    }
    return &ImageConverter{
        img: i,
    }
}



func (ic *ImageConverter) ImageToAsciiMatrix() [][]ascii.Pixel {
    sz := ic.img.Bounds()
    width := sz.Max.X
    height := sz.Max.Y

    //The initial length is set to 0 and capacity is set to the total number of pixels in the image plus an additional space for line breaks (rows of ASCII characters)
    asciiMatrix := make([][]ascii.Pixel, 0, int(height))

    for i:=0; i < int(height); i++ {
        line := make([]ascii.Pixel, 0, width)
        for j:=0; j < int(width); j++ {
            pixel := color.NRGBAModel.Convert(ic.img.At(j,i))
            asciiPixel := ascii.PixelToAsciiPixel(pixel)
            line = append(line, asciiPixel)
        }
        asciiMatrix = append(asciiMatrix, line)
    }

    return asciiMatrix

}

func (ic *ImageConverter) ImageToAsciiSlice() []string {
	// Resize the convert first
	sz := ic.img.Bounds()
	newWidth := sz.Max.X
	newHeight := sz.Max.Y
	rawCharValues := make([]string, 0, int(newWidth*newHeight+newWidth))
	for i := 0; i < int(newHeight); i++ {
		for j := 0; j < int(newWidth); j++ {
			pixel := color.NRGBAModel.Convert(ic.img.At(j, i))
			// Convert the pixel to ascii char
			rawChar := ascii.PixelToAsciiPixelString(pixel)
			rawCharValues = append(rawCharValues, rawChar)
		}
		rawCharValues = append(rawCharValues, "\n")
	}
	return rawCharValues
}

func (ic *ImageConverter) ImageToAsciiString() string {
    buf := ic.ImageToAsciiSlice()
    var buffer bytes.Buffer

    for i := 0; i< len(buf); i++ {
        buffer.WriteString(buf[i])
    }
    return buffer.String()
}

