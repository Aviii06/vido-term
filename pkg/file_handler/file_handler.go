package filehandler

import (
	"bufio"
	"os"
    "image/color"
    "image/png"
)

func ReadVidoFile(file_path string) ([]byte, error) {
    f, err := os.Open(file_path)

    if err != nil {
        panic(err)
    }

    defer f.Close()

    stats, statsErr := f.Stat()
    if statsErr != nil {
        return nil, statsErr
    }

    var size int64 = stats.Size()
    bytes := make([]byte, size)

    bufr := bufio.NewReader(f)
    _,err = bufr.Read(bytes)

    return bytes, err
}

func ReadImageFile(file_path string) ([]uint32, int, int, error) {
    f, err := os.Open(file_path)

    if err != nil {
        println("Error opening file:", err)
        return nil, 0, 0, err
    }
    defer f.Close()

    img, err := png.Decode(f)
	if err != nil {
		println("Error decoding PNG:", err)
		return nil, 0, 0, err
	}

    bounds := img.Bounds()

	// Create a slice to hold the uint32 values (size: width * height)
	width := bounds.Max.X
	height := bounds.Max.Y

    // println(bounds.Max.X)    

	pixelArray := make([]uint32, width*height)

	// Convert each pixel to uint32 in 0x00RRGGBB format
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the pixel color
			rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			// Pack the R, G, B values into a uint32: 0x00RRGGBB
			r := uint32(rgba.R)
			g := uint32(rgba.G)
			b := uint32(rgba.B)
			// Note: The first byte is left as 0x00 (for alpha or padding)
			pixelArray[y*width+x] = (r << 24) | (g << 16) | b << 8
		}
	}

    
    return pixelArray, width, height, err
}


