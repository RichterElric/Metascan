package Keyfinder

import (
	"fmt"
	"os/exec"
)

type Keyfinder struct {
	path string
	output string
}

func New(_path string, _output string) Keyfinder {
	k := Keyfinder{_path, _output}
	k.checkDependency()
	return k
}

func (k Keyfinder) checkDependency() bool{
	// TODO: if no depentency then download else return
	// TODO: installer python ?
	exec.Command("pip3 install androguard python-magic PyOpenSSL")
	return true
}

func (k Keyfinder) Scan() string{
	fmt.Println("TODO: Tester le bon téléchargement de la dépendance")
	fmt.Println("TODO: Récupérer le lieu de téléchargement pour l'insérer dans la commande")
	result,error := exec.Command("python3", "C:\\Users\\Nicolas\\Downloads\\keyfinder\\keyfinder.py").Output()
	if error!=nil{
		fmt.Println(error.Error())
	}
	return string(result)
}

