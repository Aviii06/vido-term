package pixel

import (
	"bytes"
	"math"
	"strconv"
)


var ColorToAscii []byte = []byte(" .':^*&@")

type Pixel struct {
    Red byte
    Green byte
    Blue byte
}

// Red: 3 bits, Green: 2 bits, Blue: 3 bits
func ByteToPixel(by byte) *Pixel {
    // 001 11 110
    // r = 47, g = 223, b = 207
    var r byte = by & 0b11100000
    var b byte = by & 0b00000111
    var g byte = by >> 3

    r = r + 15
    b = b << 5 + 15
    g = g << 6 + 31

    px := new(Pixel)
    px.Red = r
    px.Green = g
    px.Blue = b

    return px
}

func PixelToByte(px Pixel) byte {
    // (25, 230, 203)
    // 000 11 110 = 
    var by byte
    
    r := px.Red & 0b11100000
    g := px.Green >> 6
    b := px.Blue >> 5

    by = r + g << 3 + b

    return by
}

func MakePixel(Red byte, Green byte, Blue byte) *Pixel {
    px := new(Pixel)
    px.Red = Red
    px.Blue = Blue
    px.Green = Green

    return px
}

func GetPixelArray(pxs []byte) []byte {
    n := len(pxs)
    var buffer bytes.Buffer
    for i := 0; i < n; i++ {
        px := ByteToPixel(pxs[i])
        buffer.Write([]byte("\x1b[38;2;"))
        buffer.Write([]byte(strconv.Itoa(int(px.Red))))
        buffer.Write([]byte(";"))
        buffer.Write([]byte(strconv.Itoa(int(px.Green))))
        buffer.Write([]byte(";"))
        buffer.Write([]byte(strconv.Itoa(int(px.Blue))))
        buffer.Write([]byte("m"))

        var avg float64 = (float64(px.Red) + float64(px.Blue) + float64(px.Green)) / (3 * 32.0)
        buffer.Write([]byte{ColorToAscii[int(math.Round(avg))]})
    }

    return buffer.Bytes()
}



