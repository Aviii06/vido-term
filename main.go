package main

import (
	"os"
	"vido-term/pkg/frame"
)

func main() {
    // benchmark_time := time.Now()
    // pxs,_ := filehandler.ReadVidoFile("./akane.vido")

    // pxs, w, h, _ := filehandler.ReadImageFile(os.Args[1])

    // timeElapsed1 := time.Since(benchmark_time).Microseconds()

    f := frame.MakeFrame(os.Args[1])
    f.Draw()

    //println("\033[0mTime to time to read: in μs ", timeElapsed1)
    // println(w)
    // println(h)
}



