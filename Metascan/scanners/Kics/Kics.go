package Kics

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

type Kics struct {
	path          string
	output        string
	outputChannel chan string
}

func New(_path string, _output string, _outputChannel chan string) Kics {
	k := Kics{_path, _output, _outputChannel}
	return k
}

func (k Kics) Scan() {
	fmt.Println("TODO: Récupérer le lieu de téléchargement pour l'insérer dans la commande")
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
	for range jsonParsed.S("queries").Children() {
		files := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files")

		if files != nil {
			j := 0
			for range jsonParsed.S("queries", strconv.Itoa(i), "files").Children() {
				issueName := jsonParsed.Path("queries." + strconv.Itoa(i) + ".description").String()

				severity := jsonParsed.Path("queries." + strconv.Itoa(i) + ".severity").String()
				fileName := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files." + strconv.Itoa(j) + ".file_name").String()
				line := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files." + strconv.Itoa(j) + ".line").String()
				description := jsonParsed.Path("queries."+strconv.Itoa(i)+".files."+strconv.Itoa(j)+".actual_value").String() + " at line : " + line
				fix := jsonParsed.Path("queries." + strconv.Itoa(i) + ".files." + strconv.Itoa(j) + ".expected_value").String()

				fmt.Println(issueName, fileName, severity, description, fix)

				j++
			}
		}
		i++
	}

	kicsReturn := "KICS \n"
	k.outputChannel <- kicsReturn + string(body)
}
