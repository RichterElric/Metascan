package Dependency_checker

import (
	"Metascan/main/log_templates/Entry"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

type DependencyChecker struct {
	path          string
	output        string
	outputChannel chan []Entry.Entry
}

func New(_path string, _output string, _outputChannel chan []Entry.Entry) DependencyChecker {
	dc := DependencyChecker{_path, _output, _outputChannel}
	return dc
}

func (dc DependencyChecker) Scan() {
	cmd := exec.Command("./bin/dependency-check/bin/dependency-check.sh",
		"--enableExperimental", "--go", "/usr/local/go/bin/go", "-f", "JSON", "-o", dc.output+"/dependency-check.json", "-s", dc.path)

	err := cmd.Run()

	if err != nil {
		if err.Error() != "exit status 0" {
			log.Fatal(err)
		}
	}

	body, err := ioutil.ReadFile(dc.output + "/dependency-check.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}
	// S is shorthand for Search
	i := 0
	var entries []Entry.Entry

	for _, _ = range jsonParsed.S("dependencies").Children() {
		fileName := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".filePath").String()
		vulnerabilities := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities")

		if vulnerabilities != nil {
			j := 0
			for _, _ = range jsonParsed.S("dependencies", strconv.Itoa(i), "vulnerabilities").Children() {
				cve := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities." + strconv.Itoa(j) + ".name").String()
				severity := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities." + strconv.Itoa(j) + ".severity").String()
				description := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities." + strconv.Itoa(j) + ".description").String()

				entry := Entry.New(fileName, "", severity, cve, "", description, "")
				entries = append(entries, *entry)

				j++
			}
		}
		i++
	}

	dc.outputChannel <- entries
}
