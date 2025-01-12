package ascii

import (
	_ "image"
	"image/color"
	"math"
	"reflect"

	"github.com/aybabtme/rgbterm"
)



func PixelToAsciiPixelString(pixel color.Color) string {
    asciiPixel := PixelToAsciiPixel(pixel)
    rawChar, r, g, b := asciiPixel.Char, asciiPixel.R, asciiPixel.G, asciiPixel.B
    // rawChar := asciiPixel.Char
    //for gray scale return string([]byte{rawChar})
    return colorMyAscii(r, g, b, rawChar) 
    // return string([]byte{rawChar})
}

func PixelToAsciiPixel(pixel color.Color) Pixel {
    convertOptions := DefaultOptions
    // r, g, b, a := pixel.RGBA()
    r := reflect.ValueOf(pixel).FieldByName("R").Uint()
	g := reflect.ValueOf(pixel).FieldByName("G").Uint()
	b := reflect.ValueOf(pixel).FieldByName("B").Uint()
	a := reflect.ValueOf(pixel).FieldByName("A").Uint()
    // fmt.Println("R:",r,"G:",g,"B:",b,"A:",a)
    value := Intensity(r,g,b,a)
    precision := float64(255 * 3 / (len(convertOptions.Pixels) - 1))
	rawChar := convertOptions.Pixels[roundValue(float64(value)/precision)]
    
    return Pixel{
        Char: rawChar,
        R: uint8(r),
        G: uint8(g),
        B: uint8(b),
        A: uint8(a),
    }
}

func Intensity(r, g, b, a uint64) uint64 {
	return (r + g + b) * a / 255
}

func roundValue(value float64) int {
	return int(math.Floor(value + 0.5))
}

func colorMyAscii(r, g, b uint8, rawChar byte) string {
    coloredChar := rgbterm.FgString(string([]byte{rawChar}), uint8(r), uint8(g), uint8(b))
    return coloredChar
}
