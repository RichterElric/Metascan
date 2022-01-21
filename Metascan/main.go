package main

import (
	"Metascan/main/log_templates/Entry"
	"Metascan/main/log_templates/Log"
	"Metascan/main/parser"
	Dependency_checker "Metascan/main/scanners/Dependency-checker"
	Dotenv_linter "Metascan/main/scanners/Dotenv-linter"
	"Metascan/main/scanners/Kics"
	"Metascan/main/writers/htmlWriter"
	"Metascan/main/writers/jsonWriter"
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
	dotenvLinterEnable := flag.Bool("dl", true, "use dotenv-linter")
	formatOut := flag.String("f", "all", "format_out (json, html, ...)")

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
	if _, ok := extFiles[".env"]; ok && *dotenvLinterEnable {
		dl := Dotenv_linter.New(extFiles[".env"], "/opt/scan", outputChannel)
		go dl.Scan()
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
	if *dotenvLinterEnable {
		scan_types = append(scan_types, "dotenv_linter")
	}

	var severity_counters [4]int
	severity_counters[0] = high
	severity_counters[1] = medium
	severity_counters[2] = low
	severity_counters[3] = info

	result := Log.New(currentDate, scan_types, severity_counters, entries)

	switch *formatOut {
	case "html":
		htmlWriter.WriteHtml(*result)
		break
	case "json":
		jsonWriter.WriteJSON(*result)
		break
	default:
		htmlWriter.WriteHtml(*result)
		jsonWriter.WriteJSON(*result)

	}

	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
}
