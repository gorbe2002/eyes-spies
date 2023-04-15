package main

// import (
// 	"fmt"
// 	"image/png"
// 	"os"
// )

// func main() {
// 	// Open the image file
// 	file, err := os.Open("logo.png")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	// Decode the PNG image
// 	img, err := png.Decode(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Print the image dimensions
// 	bounds := img.Bounds()
// 	fmt.Println("Image dimensions:", bounds.Dx(), "x", bounds.Dy())
// }
