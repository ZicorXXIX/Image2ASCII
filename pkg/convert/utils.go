package convert

import (
	"image"
	"os"
    "fmt"
    // Support decode jpeg image
	_ "image/jpeg"
	// Support deocde the png image
	_ "image/png"
    "golang.org/x/image/draw"
)

func OpenImageFile(imageFileName string) (image.Image, error) {
    // fmt.Println("Opening image file:", imageFileName)
	f, err := os.Open(imageFileName)
	if err != nil {
        fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer f.Close()
	
	// Decode the image
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, err
	}

    fmt.Println("Image successfully opened and decoded")
	return img, nil
}


func resizeImage(src image.Image, newWidth, newHeight int) image.Image {
	srcBounds := src.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	// Calculate the aspect ratio of the source image
	srcAspectRatio := float64(srcWidth) / float64(srcHeight)

	// Calculate the target width and height while maintaining the aspect ratio
	var targetWidth, targetHeight int
	if float64(newWidth)/float64(newHeight) > srcAspectRatio {
		targetWidth = int(float64(newHeight) * srcAspectRatio)
		targetHeight = newHeight
	} else {
		targetWidth = newWidth
		targetHeight = int(float64(newWidth) / srcAspectRatio)
	}

	dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
    draw.CatmullRom.Scale(dst, dst.Rect, src, srcBounds, draw.Over, nil)
	return dst
}
