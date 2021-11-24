package Kics

import (
	"fmt"
	"os/exec"
)

type Kics struct {
	path          string
	output        string
	outputChannel chan string
}

func New(_path string, _output string, _outputChannel chan string) Kics {
	k := Kics{_path, _output, _outputChannel}
	k.checkDependency()
	return k
}

func (k Kics) checkDependency() bool {
	// TODO: if no depentency then download else return
	fmt.Println("TODO: Téléchargement de la dépendance")
	return true
}

func (k Kics) Scan() {
	fmt.Println("TODO: Tester le bon téléchargement de la dépendance")
	fmt.Println("TODO: Récupérer le lieu de téléchargement pour l'insérer dans la commande")
	result, error := exec.Command("D:\\GIT_Perso\\Metascan\\bin\\kicks\\kics.exe",
		"--ci", "scan", "-p", k.path, "-o", k.output, "--output-name", "kics.json").Output()
	if error != nil {
		fmt.Println(error.Error())
	}

	k.outputChannel <- string(result)
}
