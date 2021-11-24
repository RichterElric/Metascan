package Kics

import (
	"fmt"
	"os/exec"
)

type Kics struct {
	path string
	output string
}

func New(_path string, _output string) Kics {
	k := Kics{_path, _output}
	return k
}

func (k Kics) Scan() string{
	fmt.Println("TODO: Récupérer le lieu de téléchargement pour l'insérer dans la commande")
	result,error := exec.Command("C:\\Users\\nicol\\Downloads\\kics_1.4.7_windows_x64\\kics.exe",
		"--ci", "scan","-p", k.path, "-o", k.output, "--output-name", "kics.json").Output()
	if error!=nil{
		fmt.Println(error.Error())
	}
	fmt.Println("TODO: Lire le JSON et le renvoyer")
	return string(result)
}

