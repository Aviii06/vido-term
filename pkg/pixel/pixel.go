package pixel

import (
	"bytes"
	"math"
	"strconv"
)


var ColorToAscii []byte = []byte(" .':^*&@")

type Pixel struct {
    red byte
    green byte
    blue byte
}

// Red: 3 bits, Green: 2 bits, Blue: 3 bits
func ByteToPixel(by byte) *Pixel {
    // 001 11 110
    // r = 47, g = 223, b = 207
    var r byte = by >> 5
    var b byte = by & 0b00000111
    var g byte = (by - (r << 5) - b) >> 3

    r = (r << 1 + 1) << 4 - 1
    b = (b << 1 + 1) << 4 - 1
    g = (g << 1 + 1) << 5 - 1

    px := new(Pixel)
    px.red = r
    px.green = g
    px.blue = b

    return px
}

func PixelToByte(px Pixel) byte {
    // (25, 230, 203)
    // 000 11 110 = 
    var by byte
    
    r := px.red & 0b11100000
    g := px.green >> 6
    b := px.blue >> 5

    by = r + g << 3 + b

    return by
}

func MakePixel(red byte, green byte, blue byte) *Pixel {
    px := new(Pixel)
    px.red = red
    px.blue = blue
    px.green = green

    return px
}

func GetPixelArray(pxs []byte) []byte {
    n := len(pxs)
    var buffer bytes.Buffer
    for i := 0; i < n; i++ {
        px := ByteToPixel(pxs[i])
        buffer.Write([]byte("\x1b[38;2;"))
        buffer.Write([]byte(strconv.Itoa(int(px.red))))
        buffer.Write([]byte(";"))
        buffer.Write([]byte(strconv.Itoa(int(px.green))))
        buffer.Write([]byte(";"))
        buffer.Write([]byte(strconv.Itoa(int(px.blue))))
        buffer.Write([]byte("m"))

        var avg float64 = (float64(px.red) + float64(px.blue) + float64(px.green)) / (3 * 32.0)
        buffer.Write([]byte{ColorToAscii[int(math.Round(avg))]})
    }

    return buffer.Bytes()
}

