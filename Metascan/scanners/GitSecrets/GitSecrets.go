package GitSecrets

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type GitSecrets struct {
	path    string    //path of the git repo
	channel chan bool //not used
}

func New(_path string, _channel chan bool) (GitSecrets, error) {
	//creation of the object gitsecrets

	g := GitSecrets{_path, _channel}

	//verification of the smooth operation of gitsecrets
	cmd := exec.Command("git", "secrets")
	cmd.Dir = g.path
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return g, errors.New("gitsecrets not installed")
	}

	//we get the current directory because we will need to get a file in it later (forbidden_patterns)
	dirMeta, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return g, errors.New("can't get project's repertory")
	}

	//installation of hooks for gitsecrets
	cmd = exec.Command("git", "secrets", "--install", "-f")
	cmd.Dir = g.path
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return g, errors.New("can't install hook")
	}

	//we add the forbidden patterns
	cmd = exec.Command("git", "secrets", "--add-provider", "--", "cat", fmt.Sprintf("%s/scanners/GitSecrets/forbidden_patterns", dirMeta))
	cmd.Dir = g.path
	_, err = cmd.Output()
	if err != nil {
		//this cmd command return an error if it has already been executed at least one time before
		//fmt.Println(err)
		//return g, errors.New("can't add forbidden patterns' provider")
	}
	//we rewrite ALL the patterns because gitsecrets is buggy as heck
	//we open the git config file to copy in it at the end of the file
	fileConf, err := os.OpenFile(fmt.Sprintf("%s/.git/config", g.path), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return g, errors.New("can't open git config")
	}
	//we open the patterns file to read it
	filePatterns, err := os.Open(fmt.Sprintf("%s/scanners/GitSecrets/forbidden_patterns", dirMeta))
	if err != nil {
		fmt.Println(err)
		return g, errors.New("can't open patterns' file")
	}
	//we write line by line the patterns
	scannerPatterns := bufio.NewScanner(filePatterns)
	for scannerPatterns.Scan() {
		if _, err = fileConf.WriteString(fmt.Sprintf("\tpatterns = %s\n", scannerPatterns.Text())); err != nil {
			fmt.Println(err)
			return g, errors.New("can't write in git config")
		}
	}
	//check if the scanner is ok
	if err = scannerPatterns.Err(); err != nil {
		fmt.Println(err)
		return g, errors.New("scanner of patterns file is corrupted")
	}
	//we close the 2 files
	err = filePatterns.Close()
	if err != nil {
		fmt.Println(err)
		return g, errors.New("can't close patterns' file")
	}
	err = fileConf.Close()
	if err != nil {
		fmt.Println(err)
		return g, errors.New("can't close config file")
	}

	return g, nil
}

func (g GitSecrets) Scan() {
	//we scan the whole repo
	cmd := exec.Command("git", "secrets", "--scan", "-r")
	cmd.Dir = g.path
	//the result is in the error stream
	var errb bytes.Buffer
	cmd.Stderr = &errb
	//we run the command and return the result
	cmd.Run()
	//we output the result
	outfile, _ := os.OpenFile("/opt/scan/metascan_results/gitsecret.txt", os.O_RDWR|os.O_CREATE, 0755)
	outfile.WriteString(errb.String())
	err := outfile.Close()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	g.channel <- true
}
