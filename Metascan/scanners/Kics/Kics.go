package Kics

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
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
	cmd := exec.Command("bin\\kics\\kics",
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

	kicsReturn := "KICS \n"
	k.outputChannel <- kicsReturn + string(body)
}
