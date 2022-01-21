package Kics

import (
	"Metascan/main/log_templates/Entry"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

type Kics struct {
	path          string
	output        string
	outputChannel chan []Entry.Entry
}

func New(_path string, _output string, _outputChannel chan []Entry.Entry) Kics {
	k := Kics{_path, _output, _outputChannel}
	return k
}

func (k Kics) Scan() {
	cmd := exec.Command("./bin/kics/kics",
		"-s", "scan", "-p", k.path, "-o", k.output, "--output-name", "kics.json")

	err := cmd.Run()

	if err != nil {
		if err.Error() != "exit status 50" && err.Error() != "exit status 40" && err.Error() != "exit status 30" && err.Error() != "exit status 20" && err.Error() != "exit status 0" {
			log.Fatal(err)
		}
	}

	body, err := ioutil.ReadFile(k.output + "/kics.json")
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

	for range jsonParsed.S("queries").Children() {
		files := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files")

		if files != nil {
			j := 0
			for range jsonParsed.S("queries", strconv.Itoa(i), "files").Children() {
				issueName := jsonParsed.Path("queries." + strconv.Itoa(i) + ".description").String()

				severity := jsonParsed.Path("queries." + strconv.Itoa(i) + ".severity").String()
				fileName := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files." + strconv.Itoa(j) + ".file_name").String()
				fileNameCut := regexp.MustCompile("../../../opt/scan/").Split(fileName, -1)
				line := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files." + strconv.Itoa(j) + ".line").String()
				description := jsonParsed.Path("queries."+strconv.Itoa(i)+".files."+strconv.Itoa(j)+".actual_value").String() + " at line : " + line
				fix := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files." + strconv.Itoa(j) + ".expected_value").String()

				entry := Entry.New(fileNameCut[1], issueName, severity, "", "", description, fix)
				entries = append(entries, *entry)

				j++
			}
		}
		i++
	}

	k.outputChannel <- entries
}
