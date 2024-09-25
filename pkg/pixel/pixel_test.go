package pixel

import "testing"


func TestByteToPixel(t *testing.T) {
    var input byte = 0b00111110

    px := new(Pixel)
    px.red = byte(47)
    px.green = byte(223)
    px.blue = byte(207)

    px2 := ByteToPixel(input)

    if (px.red != px2.red || 
    px.blue != px2.blue || 
    px.green != px2.green) {
        t.Fatalf("TestByteToPixel failed")
    }
}


func TestPixelToByte(t *testing.T) {
    px := new(Pixel)
    px.red = byte(25)
    px.green = byte(230)
    px.blue = byte(203)

    var expect byte = byte(0b00011110)

    if (expect != PixelToByte(*px)) {
        t.Fatalf("TestByteToPixel failed")
    }
}


