package main

import (
	"Metascan/main/log_templates/Entry"
	"Metascan/main/log_templates/Log"
	"Metascan/main/parser"
	Dependency_checker "Metascan/main/scanners/Dependency-checker"
	"Metascan/main/scanners/Kics"
	"Metascan/main/writers/htmlWriter"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	baseDir := flag.String("d", "/opt/scan", "the base directory for the recursive search")
	kicksEnable := flag.Bool("kics", true, "use kics")
	keyFinderEnable := flag.Bool("kf", false, "use keyFinder") // experimental
	gitSecretEnable := flag.Bool("gits", true, "use git Secret")
	dependencyCheckerEnable := flag.Bool("dc", true, "use dependencyChecker")

	flag.Parse()

	fmt.Println("base directory : " + *baseDir)
	start := time.Now()

	extFiles := parser.GetFiles(*baseDir)
	//log.Println(extFiles)

	outputChannel := make(chan []Entry.Entry)
	nbOutput := 0

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

	var entries []Entry.Entry
	var high = 0
	var medium = 0
	var low = 0
	var info = 0

	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02 15:04:05")

	for i := 0; i < nbOutput; i++ {
		entriesThread := <-outputChannel
		for _, entry := range entriesThread {
			entry.Filename = strings.TrimPrefix(entry.Filename, "\"")
			entry.Filename = strings.TrimSuffix(entry.Filename, "\"")
			entry.Severity = strings.TrimPrefix(entry.Severity, "\"")
			entry.Severity = strings.TrimSuffix(entry.Severity, "\"")
			entry.CVE = strings.TrimPrefix(entry.CVE, "\"")
			entry.CVE = strings.TrimSuffix(entry.CVE, "\"")
			entry.CWE = strings.TrimPrefix(entry.CWE, "\"")
			entry.CWE = strings.TrimSuffix(entry.CWE, "\"")


			print(strings.ToLower(entry.Severity))
			if strings.Contains(strings.ToLower(entry.Severity),"high") {
				high += 1
			} else if strings.Contains(strings.ToLower(entry.Severity),"medium") {
				medium += 1
			} else if strings.Contains(strings.ToLower(entry.Severity),"low") {
				low += 1
			} else if strings.Contains(strings.ToLower(entry.Severity),"info") {
				info += 1
			}

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

	//TODO: Ajouter un paramètre pour choisir le type de fichier ? Si pas besoin, retirer les if et la variable output_type
	var output_type = "html"
	if (output_type == "html"){
		htmlWriter.WriteHtml(*result)
	}
	if (output_type == "json"){
		//writeJson()
	}

	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
	log.Println(result)
}
