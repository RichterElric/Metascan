package Keyfinder

import (
	"fmt"
	"os/exec"
)

type Keyfinder struct {
	path   string
	output string
}

func New(_path string, _output string) Keyfinder {
	k := Keyfinder{_path, _output}
	return k
}

func (k Keyfinder) Scan() string {
	fmt.Println("TODO: Tester le bon téléchargement de la dépendance")
	fmt.Println("TODO: Récupérer le lieu de téléchargement pour l'insérer dans la commande")
	result, error := exec.Command("python", "C:\\Users\\nicol\\Downloads\\keyfinder-master\\keyfinder.py", "-k", k.path).Output()
	if error != nil {
		fmt.Println(error.Error())
	}
	return string(result)
}
