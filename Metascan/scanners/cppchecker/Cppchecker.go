package cppchecker

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Cppstruct struct {
	path       string
	cppChannel chan bool
}

func New(_path string, _cppChannel chan bool) Cppstruct {
	k := Cppstruct{_path, _cppChannel}
	return k
}

func Scan(cpps Cppstruct) {
	// We execute the cppcheck binary with all checks and an xml output
	out, err := exec.Command("./bin/cppcheck", "--enable=all", "-q", "--xml", cpps.path).CombinedOutput()
	// We prepare the xml output to write to report
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
	cpps.cppChannel <- true
}
