package Dotenv_linter

import (
	"Metascan/main/log_templates/Entry"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type DotenvLinter struct {
	path          []string
	output        string
	outputChannel chan []Entry.Entry
}

func New(_path []string, _output string, _outputChannel chan []Entry.Entry) DotenvLinter {
	dc := DotenvLinter{_path, _output, _outputChannel}
	return dc
}

func (dl DotenvLinter) Scan() {
	out, err := exec.Command("dotenv-linter", dl.path...).CombinedOutput()
	outfile, _ := os.OpenFile(dl.output+"/dotenv_linter.txt", os.O_RDWR|os.O_CREATE, 0755)

	_, err2 := outfile.Write(out)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(string(out))
	if err != nil {
		if err.Error() != "exit status 1" {
			log.Fatal(err)
		}
	}

	//body, err := ioutil.ReadFile(dl.output + "/dotenv-linter.txt")
	//if err != nil {
	//	log.Fatalf("unable to read file: %v", err)
	//}
	//
	var entries []Entry.Entry

	dl.outputChannel <- entries
}
