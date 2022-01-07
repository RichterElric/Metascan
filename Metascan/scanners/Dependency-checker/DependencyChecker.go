package Dependency_checker

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

type DependencyChecker struct {
	path          string
	output        string
	outputChannel chan string
}

func New(_path string, _output string, _outputChannel chan string) DependencyChecker {
	dc := DependencyChecker{_path, _output, _outputChannel}
	return dc
}

func (dc DependencyChecker) Scan() {
	cmd := exec.Command("bin\\dependency-check\\bin\\dependency-check.bat",
		"--enableExperimental", "--go", "C:\\Users\\vducros\\sdk\\go1.17.2\\bin\\go.exe", "-f", "JSON", "-o", dc.output+"/dependency-check.json", "-s", dc.path)
	//TODO : mettre les paths en absolu avec le docker

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
	for _, _ = range jsonParsed.S("dependencies").Children() {
		//filePath := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".filePath").String()
		//fileName := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".fileName").String()
		vulnerabilities := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities")

		if vulnerabilities != nil {
			j := 0
			for _, _ = range jsonParsed.S("dependencies", strconv.Itoa(i), "vulnerabilities").Children() {
				cve := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities." + strconv.Itoa(j) + ".name")
				severity := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities." + strconv.Itoa(j) + ".severity")
				description := jsonParsed.Path("dependencies." + strconv.Itoa(i) + ".vulnerabilities." + strconv.Itoa(j) + ".description")

				fmt.Println(cve, severity, description)

				j++
			}
		}
		i++
	}

	dc.outputChannel <- string(body)
}
