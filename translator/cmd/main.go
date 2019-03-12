package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	argFilename := os.Args[1]
	text := newScrollFromFile(argFilename)
	text.printScrollSummary()
	fmt.Printf("\n Glooglon report took %v \n", time.Since(start))
}
