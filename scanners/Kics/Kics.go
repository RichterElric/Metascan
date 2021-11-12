package Kics

import (
	"fmt"
	"os/exec"
)

type Kics struct {
	path string
}

func New(_path string) Kics {
	k := Kics{_path}
	return k
}

func (k Kics) GetDependency() bool{
	fmt.Println("TODO: Téléchargement de la dépendance")
	return true
}

func (k Kics) Scan() string{
	result,error := exec.Command("C:\\Users\\Nicolas\\Downloads\\kics_1.4.7_windows_x64\\kics.exe", "help -p" + k.path).Output()
	if error!=nil{
		fmt.Println(error.Error())
	}
	return string(result)
}

