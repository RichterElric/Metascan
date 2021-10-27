package main

import (
	"Metascan/main/parser"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	baseDir := flag.String("d", ".", "the base directory for the recursive search")

	flag.Parse()

	fmt.Println("base directory : " + *baseDir)
	start := time.Now()
	parser.GetFiles(*baseDir)
	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
}
