package PMD

import (
	"fmt"
	"log"
	"os/exec"
)

type PMD struct {
	path    string
	channel chan bool
}

func New(_path string, _channel chan bool) PMD {
	pmd := PMD{_path, _channel}
	return pmd
}

func (pmd PMD) Scan(lang string) {
	if lang == "java" {
		_, err := exec.Command("./bin/pmd-bin-6.42.0/bin/run.sh", "pmd", "-d", pmd.path, "-R", "rulesets/java/quickstart.xml", "-f", "html", "--report-file", "/opt/scan/metascan_results/pmd_java.html").Output()
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	if lang == "xml" {
		_, err := exec.Command("./bin/pmd-bin-6.42.0/bin/run.sh", "pmd", "-d", pmd.path, "-R", "rulesets/xml/basic.xml", "-f", "html", "--report-file", "/opt/scan/metascan_results/pmd_xml.html").Output()
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	pmd.channel <- true
}
