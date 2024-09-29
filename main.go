package main

import (
	"os"
	"vido-term/pkg/frame"
    // "vido-term/pkg/file_handler"
    "time"
)

func main() {
    // benchmark_time := time.Now()
    // pxs,_ := filehandler.ReadVidoFile("./akane.vido")

    // pxs, w, h, _ := filehandler.ReadImageFile(os.Args[1])

    // timeElapsed1 := time.Since(benchmark_time).Microseconds()

    // fr := frame.MakeFrame(w, h, pxs)
    // fr.Draw()

    benchmark_time := time.Now()
    frame.DrawOptimised(os.Args[1])
    timeElapsed := time.Since(benchmark_time).Microseconds()

    println("\033[0mTotal time ", timeElapsed)
    return
    // println(w)
    // println(h)
}



