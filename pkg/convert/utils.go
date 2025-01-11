package convert

import (
	"image"
	"os"
    "fmt"
    // Support decode jpeg image
	_ "image/jpeg"
	// Support deocde the png image
	_ "image/png"
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
	d := resizeImage(img, 80, 80)

    fmt.Println("Image successfully opened and decoded")
	return d, nil
}
