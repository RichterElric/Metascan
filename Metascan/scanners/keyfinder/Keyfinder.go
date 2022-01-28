package Keyfinder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Keyfinder struct {
	path    string
	channel chan bool
}

func New(_path string, _channel chan bool) Keyfinder {
	k := Keyfinder{_path, _channel}
	return k
}

func (k Keyfinder) Scan() {
	result, err := exec.Command("python3", "./bin/keyfinder/keyfinder.py", "-k", k.path).Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//we output the result
	outfile, err := os.OpenFile("/opt/scan/metascan_results/keyfinder.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	outfile.Write(result)
	err = outfile.Close()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	k.channel <- true
}
