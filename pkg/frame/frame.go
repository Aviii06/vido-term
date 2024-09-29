package frame

import "time"
import "encoding/binary"
import "vido-term/pkg/pixel"
import "bufio"
import "os"
import "image/png"
import "image/color"



func Draw(file_path string) {

    benchmark_time := time.Now()

    file, err := os.Open(file_path)

    if err != nil {
        println("Error opening file:", err)
        return    
    }


    defer file.Close()

    img, err := png.Decode(file)
	if err != nil {
		println("Error decoding PNG:", err)
		return 
	}

    bounds := img.Bounds()

	// Create a slice to hold the uint32 values (size: width * height)
	width := bounds.Max.X
	height := bounds.Max.Y
    println(width)
    println(height)

    // Convert each pixel to uint32 in 0x00RRGGBB format
    var by []byte = make([]byte, 20 * (width + 100) * height + 1000)

    currByte := 0
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            // Get the pixel color
            rgba := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
            // Pack the R, G, B values into a uint32: 0x00RRGGBB
            r := uint32(rgba.R)
            g := uint32(rgba.G)
            b := uint32(rgba.B)
            // Note: The first byte is left as 0x00 (for alpha or padding)
            bb := (r << 24) | (g << 16) | b << 8
            a := make([]byte, 4)
            binary.LittleEndian.PutUint32(a, bb)

            r1, r2 ,r3 := convertToBytes(a[3])
            g1, g2 ,g3 := convertToBytes(a[2])
            b1, b2 ,b3 := convertToBytes(a[1])

            var avg byte = (a[0] >> 5 + a[1] >> 5 + a[2] >> 5) / 3
            var pixByte [20]byte = [20]byte{27, 91, 52, 56, 59, 50, 59, r1, r2, r3, 59, g1, g2, g3, 59, b1, b2, b3, 109, pixel.ColorToAscii[avg+1]}

            for k := 0; k < 20; k++ {
                by[currByte + k] = pixByte[k]
            }

            currByte += 20

        }
        by[currByte] = 10
        currByte += 1
    }


    timeElapsed1 := time.Since(benchmark_time).Microseconds()

    writer := bufio.NewWriter(os.Stdout)
    writer.Write(by)
    writer.Flush()
    timeElapsed2 := time.Since(benchmark_time).Microseconds()

    println("\033[0mTime to time to Make frame in μs: ", timeElapsed1)
    println("\033[0mTime to stdout in μs: ", timeElapsed2)
}

