package PyLint

import (
	"os/exec"
)

//A __init__.py FILE IS NEEDED
type PyLint struct {
	path    string
	channel chan bool
}

func New(_path string, _channel chan bool) PyLint {
	p := PyLint{_path, _channel}
	return p
}

func (p PyLint) Scan() {
	_, err := exec.Command("pylint", p.path, "-f", "json", "--output", "/opt/scan/metascan_results/pylint.json").Output()
	if err != nil {
		//fmt.Println(err) //The func send an error even if it worked well
		//log.Fatal(err)  //The func send an error even if it worked well
	}
	_, err = exec.Command("pylint", p.path, "--output", "/opt/scan/metascan_results/pylint.txt").Output()
	if err != nil {
		//fmt.Println(err) //The func send an error even if it worked well
		//log.Fatal(err)  //The func send an error even if it worked well
	}
	p.channel <- true
}
