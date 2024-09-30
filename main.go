package main

import (
	"os"
	"time"
	"vido-term/pkg/frame"
)

func main() {
    benchmark_time := time.Now()
    frame.DrawOptimised(os.Args[1])
    timeElapsed := time.Since(benchmark_time).Microseconds()

    println("\033[0mTotal time ", timeElapsed)
    return
}



