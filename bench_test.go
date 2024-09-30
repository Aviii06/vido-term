package main

import (
    "testing"
    "vido-term/pkg/frame"
)

func BenchmarkDraw(b *testing.B) {
    b.Run("Bechmark draw optimised", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            frame.DrawOptimised("./assets/akane-main.png")
        }
    })
}

