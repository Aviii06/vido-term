package frame

import (
    "os"
    "vido-term/pkg/pixel"
    "syscall"
    "time"
)

type Frame struct {
    width int
    height int
    pixels []byte
}

func (f *Frame) Draw() {
    var by []byte = make([]byte, 20 * (f.width + 1) * f.height + 1)

    benchmark_time := time.Now()
    currByte := 0
    for j := 0; j < f.height; j++ {
        // Write rendering logic here
        for i := 0; i < f.width; i++ {
            bb := f.pixels[j*f.width + i]

            // Some bitwise magic trust me it works
            var r byte = bb & 0b11100000
            var b byte = bb & 0b00000111
            var g byte = bb >> 3

            r = r + 15
            b = b << 5 + 15
            g = g << 6 + 31

            r1, r2 ,r3 := convertToBytes(r)
            g1, g2 ,g3 := convertToBytes(g)
            b1, b2 ,b3 := convertToBytes(b)

            var avg byte = (r >> 5 + b >> 5 + g >> 5) / 3
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

    Stdout := os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stdout.Write(by)
    timeElapsed2 := time.Since(benchmark_time).Microseconds()

    println("\033[0mTime to time to Make frame in μs: ", timeElapsed1)
    println("\033[0mTime to stdout in μs: ", timeElapsed2)
}

func convertToBytes(by byte) (byte, byte, byte) {
    d1 := by / 100
    d2 := (by % 100) / 10
    d3 := (by % 10) 

    return (d1 + 48), (d2 + 48), (d3 + 48)

}

func MakeFrame(width int, height int, pixels []byte) *Frame {
    f := new(Frame)
    f.width = width
    f.height = height
    f.pixels = pixels

    return f
}

