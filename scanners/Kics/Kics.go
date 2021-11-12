package Kics

import "fmt"

type Kics struct {
	path string
}

func New(_path string) Kics {
	k := Kics{_path}
	return k
}

func (Kics) GetDependency() bool{
	fmt.Println("TODO: Téléchargement de la dépendance")
	return true
}

func (Kics) Scan() string{
	return "Results !"
}

