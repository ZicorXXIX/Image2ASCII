package main

import (
	"fmt"
	"os"
    "flag"

	"github.com/ZicorXXIX/Image2ASCII/pkg/convert"
)

func main() {
    fmt.Println("Hello World")
    var path string
	flag.StringVar(&path, "path", "pikachu.png", "Path to image")
    flag.Parse()
    fmt.Println("path:",path)
    converter:= convert.NewImageConverter(path)
    matrix := converter.ImageToAsciiString()
    saveToFile("dum.txt", matrix)
    fmt.Println(matrix)
    // f, err := os.Open("test.png")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(f)
    // wd, err := os.Getwd()
    // if err != nil {
    //     fmt.Println("Error getting working directory:", err)
    //     return
    // }
    // fmt.Println("Current Working Directory:", wd)
}

func saveToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed when we're done

	_, err = file.WriteString(content) // Write the ASCII art to the file
	return err
}


