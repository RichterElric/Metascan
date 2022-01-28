package main

import (
	"Metascan/main/log_templates/Entry"
	"Metascan/main/log_templates/Log"
	"Metascan/main/parser"
	Dependency_checker "Metascan/main/scanners/Dependency-checker"
	"Metascan/main/scanners/Kics"
	"Metascan/main/scanners/cppchecker"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	baseDir := flag.String("d", "/opt/scan", "the base directory for the recursive search")
	kicksEnable := flag.Bool("kics", true, "use kics")
	keyFinderEnable := flag.Bool("kf", false, "use keyFinder") // experimental
	gitSecretEnable := flag.Bool("gits", true, "use git Secret")
	dependencyCheckerEnable := flag.Bool("dc", true, "use dependencyChecker")
	cppcheckEnable := flag.Bool("cpp", true, "use cppchecker")

	flag.Parse()

	fmt.Println("base directory : " + *baseDir)
	start := time.Now()

	extFiles := parser.GetFiles(*baseDir)
	//log.Println(extFiles)

	outputChannel := make(chan []Entry.Entry)
	nbOutput := 0

	cppChannel := make(chan bool)

	if _, ok := extFiles["kics"]; ok && *kicksEnable {
		k := Kics.New(*baseDir, "/opt/scan/", outputChannel)
		go k.Scan()
		nbOutput++
	}
	if *keyFinderEnable {
		fmt.Println("USE KEY FINDER")
		//nbOutput++
	}
	if *gitSecretEnable {
		fmt.Println("USE GIT SECRET")
		//nbOutput++
	}
	if *dependencyCheckerEnable {
		k := Dependency_checker.New(*baseDir, "/opt/scan/", outputChannel)
		go k.Scan()
		nbOutput++
	}
	if *cppcheckEnable {
		cpps := cppchecker.New(*baseDir, cppChannel)
		goCpp := false
		if _, ok := extFiles[".c"]; ok {
			go cppchecker.Scan(cpps)
			goCpp = true
		}
		if _, ok := extFiles[".cpp"]; ok && !goCpp {
			go cppchecker.Scan(cpps)
		}
	}

	var entries []Entry.Entry
	var high = 0
	var medium = 0
	var low = 0
	var info = 0

	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02 15:04:05")

	for i := 0; i < nbOutput; i++ {
		// TODO: Récupérer les outputs
		// TODO: Incrémenter high,medium,low et info en fonction de la sévérité
		entriesThread := <-outputChannel
		for _, entry := range entriesThread {
			entries = append(entries, entry)
		}
	}
	var scan_types []string
	if *kicksEnable {
		scan_types = append(scan_types, "kicks")
	}
	if *keyFinderEnable {
		scan_types = append(scan_types, "key finder")
	}
	if *gitSecretEnable {
		scan_types = append(scan_types, "git secrets")
	}
	if *dependencyCheckerEnable {
		scan_types = append(scan_types, "dependency checker")
	}

	var severity_counters [4]int
	severity_counters[0] = high
	severity_counters[1] = medium
	severity_counters[2] = low
	severity_counters[3] = info

	result := Log.New(currentDate, scan_types, severity_counters, entries)
	// we wait for the cpp scan to end
	<-cppChannel

	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
	log.Println(result)
}
