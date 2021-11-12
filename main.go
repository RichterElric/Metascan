package main

import (
	"Metascan/main/parser"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	baseDir := flag.String("d", "D:\\GIT_Perso\\sample-programs-main", "the base directory for the recursive search")

	flag.Parse()

	fmt.Println("base directory : " + *baseDir)
	start := time.Now()

	extFiles := parser.GetFiles(*baseDir)
	log.Println(extFiles)

	if _, ok := extFiles["kicks"]; ok {
		fmt.Println("USE KICKS")
		fmt.Println(extFiles["kicks"])
	}

	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
}
