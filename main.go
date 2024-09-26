package main

import (
	filehandler "vido-term/pkg/file_handler"
	"vido-term/pkg/frame"
    "time"
)

func main() {
    benchmark_time := time.Now()
    pxs,_ := filehandler.ReadVidoFile("./akane.vido")
    timeElapsed1 := time.Since(benchmark_time).Microseconds()

    w := 1280
    h := 720
    f := frame.MakeFrame(w,  h, pxs)

    f.Draw()
    println("\033[0mTime to time to read: in μs ", timeElapsed1)
}



