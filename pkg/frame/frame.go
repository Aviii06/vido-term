package frame

import (
	"bufio"
	"encoding/binary"
    "image"
	// "image/color"
	"image/png"
	"os"
	"time"
	"vido-term/pkg/pixel"
)

type Frame struct {
    width int
    height int
    pixels []uint32
}

func (f *Frame) Draw() {
    var by []byte = make([]byte, 20 * (f.width + 1) * f.height + 1)

    benchmark_time := time.Now()
    currByte := 0
    for j := 0; j < f.height; j++ {
        // Write rendering logic here
        for i := 0; i < f.width; i++ {
            bb := f.pixels[j*f.width + i]
            a := make([]byte, 4)
            binary.LittleEndian.PutUint32(a, bb)

            // Some bitwise magic trust me it works
            r1, r2 ,r3 := convertToBytes(a[3])
            g1, g2 ,g3 := convertToBytes(a[2])
            b1, b2 ,b3 := convertToBytes(a[1])

            var avg byte = (a[0] >> 5 + a[1] >> 5 + a[2] >> 5) / 3
            var pixByte [20]byte = [20]byte{27, 91, 51, 56, 59, 50, 59, r1, r2, r3, 59, g1, g2, g3, 59, b1, b2, b3, 109, pixel.ColorToAscii[avg+1]}

            for k := 0; k < 20; k++ {
                by[currByte + k] = pixByte[k]
            }

            currByte += 20

        }

        // New line
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

func DrawOptimised(file_path string) {
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

    timeElapsed1 := time.Since(benchmark_time).Microseconds()

    benchmark_time = time.Now()
    var by []byte = make([]byte, 20 * (width + 1) * height + 1)

    currByte := 0

    // Takes about 34 ms on m1 pro 8 gigs.
    go func() {
        for y :=0; y < height; y++ {
            for x := 0; x < width; x++ {
                // Get the pixel color
                imgNRGBA, ok := img.(*image.NRGBA)
                if !ok {
                    return
                }
                offset := imgNRGBA.PixOffset(x, y)
                r := imgNRGBA.Pix[offset+0]
                g := imgNRGBA.Pix[offset+1]
                b := imgNRGBA.Pix[offset+2]

                r1, r2 ,r3 := convertToBytes(r)
                g1, g2 ,g3 := convertToBytes(g)
                b1, b2 ,b3 := convertToBytes(b)
                var pixByte [20]byte = [20]byte{27, 91, 52, 56, 59, 50, 59, r1, r2, r3, 59, g1, g2, g3, 59, b1, b2, b3, 109, 32}

                for k := 0; k < 20; k++ {
                    by[currByte + k] = pixByte[k]
                }
                currByte += 20
            }

            by[currByte] = 10
            currByte += 1
        }
    }()


    timeElapsed2 := time.Since(benchmark_time).Microseconds()

    benchmark_time = time.Now()

    // Slowest part
    writer := bufio.NewWriter(os.Stdout)
    writer.Write(by)
    writer.Flush()
    timeElapsed3 := time.Since(benchmark_time).Microseconds()

    println("\033[0mTime to read in μs: ", timeElapsed1)
    println("\033[0mTime to time to Make frame in μs: ", timeElapsed2)
    println("\033[0mTime to stdout in μs: ", timeElapsed3)

    // println("Width: ", width)
    // println("Height: ", height)
}

func MakeFrame(width int, height int, pixels []uint32) *Frame {
    f := new(Frame)
    f.width = width
    f.height = height
    f.pixels = pixels

    return f
}

