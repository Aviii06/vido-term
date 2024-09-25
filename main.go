package main

import (
	"time"
	filehandler "vido-term/pkg/file_handler"
	"vido-term/pkg/frame"
)

func main() {
    benchmark_time := time.Now()
    pxs,_ := filehandler.ReadVidoFile("./akane.vido")
    timeElapsed1 := time.Since(benchmark_time).Microseconds()

    benchmark_time = time.Now() 
    w := 128
    h := 72
    f := frame.MakeFrame(w,  h, pxs)

    timeElapsed2 := time.Since(benchmark_time).Microseconds()

    benchmark_time = time.Now()    
    f.Draw()
    timeElapsed3 := time.Since(benchmark_time).Microseconds()

    println("\033[0mTime to time to read: in μs ", timeElapsed1)
    println("\033[0mTime to time to Make frame in μs: ", timeElapsed2)
    println("\033[0mTime to draw in μs: ", timeElapsed3)
}



