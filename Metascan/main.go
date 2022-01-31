package main

import (
	"Metascan/main/log_templates/Entry"
	"Metascan/main/log_templates/Log"
	"Metascan/main/parser"
	Dependency_checker "Metascan/main/scanners/Dependency-checker"
	Dotenv_linter "Metascan/main/scanners/Dotenv-linter"
	"Metascan/main/scanners/GitSecrets"
	"Metascan/main/scanners/Kics"
	"Metascan/main/scanners/PMD"
	"Metascan/main/scanners/PyLint"
	"Metascan/main/scanners/cppchecker"
	"Metascan/main/writers/htmlWriter"
	"Metascan/main/writers/jsonWriter"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// Creating base foler
	if _, err := os.Stat("/opt/scan/metascan_results"); os.IsNotExist(err) {
		err := os.Mkdir("/opt/scan/metascan_results", 0755)

		if err != nil {
			log.Fatal(err)
		}
	}

	baseDir := flag.String("d", "/opt/scan", "the base directory for the recursive search")
	kicksEnable := flag.Bool("kics", true, "use kics")
	pmdEnable := flag.Bool("PMD", true, "use PMD")
	pyLintEnable := flag.Bool("pylint", true, "use PyLint")
	gitSecretEnable := flag.Bool("gits", true, "use git Secret")
	dependencyCheckerEnable := flag.Bool("dc", true, "use dependencyChecker")
	cppcheckEnable := flag.Bool("cpp", true, "use cppchecker")
	dotenvLinterEnable := flag.Bool("dl", true, "use dotenv-linter")
	formatOut := flag.String("f", "all", "format_out (json, html, ...)")

	flag.Parse()

	fmt.Println("base directory : " + *baseDir)
	start := time.Now()

	extFiles := parser.GetFiles(*baseDir)
	//log.Println(extFiles)

	outputChannel := make(chan []Entry.Entry)
	nbOutput := 0

	cppChannel := make(chan bool)
	gitSecretsChannel := make(chan bool)
	pyLintChannel := make(chan bool)
	pmdChannel := make(chan bool)

	if _, ok := extFiles["kics"]; ok && *kicksEnable {
		k := Kics.New(*baseDir, "/opt/scan/metascan_results", outputChannel)
		go k.Scan()
		nbOutput++
	}
	goPMD_java := false
	goPMD_xml := false
	if *pmdEnable {
		pmd := PMD.New("/opt/scan/", pmdChannel)
		if _, ok := extFiles[".jar"]; ok {
			go pmd.Scan("java")
			goPMD_java = true
		}
		if _, ok := extFiles[".xml"]; ok {
			go pmd.Scan("xml")
			goPMD_xml = true
		}
	}
	goPyLint := false
	if *pyLintEnable {
		p := PyLint.New("/opt/scan/", pyLintChannel)
		if _, ok := extFiles[".py"]; ok {
			go p.Scan()
			goPyLint = true
		}
	}
	goGitSecret := false
	if *gitSecretEnable {
		if _, ok := extFiles[".git"]; ok {
			g, err := GitSecrets.New("/opt/scan/", gitSecretsChannel)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			} else {
				go g.Scan()
				goGitSecret = true
			}
		}
	}
	if *dependencyCheckerEnable {
		k := Dependency_checker.New(*baseDir, "/opt/scan/metascan_results", outputChannel)
		go k.Scan()
		nbOutput++
	}
	if _, ok := extFiles[".env"]; ok && *dotenvLinterEnable {
		dl := Dotenv_linter.New(extFiles[".env"], "/opt/scan/metascan_results", outputChannel)
		go dl.Scan()
		nbOutput++
	}
	goCpp := false
	if *cppcheckEnable {
		cpps := cppchecker.New(*baseDir, cppChannel)
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

			if strings.Contains(strings.ToLower(entry.Severity), "high") {
				high += 1
			} else if strings.Contains(strings.ToLower(entry.Severity), "medium") {
				medium += 1
			} else if strings.Contains(strings.ToLower(entry.Severity), "low") {
				low += 1
			} else if strings.Contains(strings.ToLower(entry.Severity), "info") {
				info += 1
			}
			entries = append(entries, entry)
		}
	}
	var scan_types []string
	if *kicksEnable {
		scan_types = append(scan_types, "kicks")
	}
	if *gitSecretEnable {
		scan_types = append(scan_types, "git secrets")
	}
	if *cppcheckEnable {
		scan_types = append(scan_types, "cpp check")
	}
	if *pmdEnable {
		scan_types = append(scan_types, "pmd")
	}
	if *pyLintEnable {
		scan_types = append(scan_types, "pylint")
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
	// we wait for the cpp scan to end

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
	if goCpp {
		<-cppChannel
	}
	if goGitSecret {
		<-gitSecretsChannel
	}
	if goPyLint {
		<-pyLintChannel
	}
	if goPMD_java {
		<-pmdChannel
	}
	if goPMD_xml {
		<-pmdChannel
	}

	elapsed := time.Since(start)
	log.Printf("FileParser took %s", elapsed)
}
