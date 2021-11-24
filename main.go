package main

import (
	"Metascan/main/parser"
	"Metascan/main/scanners/Kics"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	baseDir := flag.String("d", ".", "the base directory for the recursive search")
	kicksEnable := flag.Bool("kics", true, "use kics")
	keyFinderEnable := flag.Bool("kf", true, "use keyFinder")
	gitSecretEnable := flag.Bool("gits", true, "use git Secret")

	flag.Parse()

	fmt.Println("base directory : " + *baseDir)
	start := time.Now()

	extFiles := parser.GetFiles(*baseDir)
	//log.Println(extFiles)

	outputChannel := make(chan string)

	if _, ok := extFiles["kics"]; ok && *kicksEnable {
		k := Kics.New(*baseDir, ".", outputChannel)
		go k.Scan()
	}
	if *keyFinderEnable {
		fmt.Println("USE KEY FINDER")
	}
	if *gitSecretEnable {
		fmt.Println("USE GIT SECRET")
	}

	fmt.Println(outputChannel)
	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
}
