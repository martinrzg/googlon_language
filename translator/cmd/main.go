package main

import "os"

func main() {
	argFilename := os.Args[1]
	text := newScrollFromFile(argFilename)
	text.printScrollSummary()
}
