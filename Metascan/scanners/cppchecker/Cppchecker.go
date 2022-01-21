package cppchecker

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Cppstruct struct {
	path string
}

func New(_path string) Cppstruct {
	k := Cppstruct{_path}
	return k
}

func Scan(cpps Cppstruct) {
	out, err := exec.Command("./bin/cppcheck", "--enable=all", "-q", "--xml", "/opt/scan").CombinedOutput()

	outfile, _ := os.OpenFile("/opt/scan/cpperr.xml", os.O_RDWR|os.O_CREATE, 0755)
	outfile.Write(out)

	err2 := outfile.Close()
	if err2 != nil {
		log.Fatal(err2)
	}
	if err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}
}
