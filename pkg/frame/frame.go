package frame

import (
	"bytes"
	"vido-term/pkg/pixel"
)

// TODO: Instead of having Pixels, frame will have bytes and use ByteToPixel.
type Frame struct {
    width int
    height int
    pixels []byte
}

func (f *Frame) Draw() {
    var buf bytes.Buffer
    buf.Grow(f.width * f.height)
    for i := 0; i < f.height; i++ {
        // Write rendering logic here
        buf.Write(pixel.GetPixelArray(f.pixels[i * f.width : (i + 1)*f.width]))

        // New line
        buf.Write([]byte("\n"))
    }

    println(buf.String())
}

func MakeFrame(width int, height int, pixels []byte) *Frame {
    f := new(Frame)
    f.width = width
    f.height = height
    f.pixels = pixels

    return f
}




