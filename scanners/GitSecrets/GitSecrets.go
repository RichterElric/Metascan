package GitSecrets

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type GitSecrets struct {
	path   string //path of the git repo
	output string //not used
}

func New(_path string, _output string) (GitSecrets, error) {
	//creation of the object gitsecrets
	g := GitSecrets{_path, _output}

	//verification of the smooth operation of gitsecrets
	cmd := exec.Command("cmd", "/C", "git", "secrets")
	cmd.Dir = g.path
	_, error := cmd.Output()
	if error != nil {
		fmt.Println(error.Error())
		return g, errors.New("gitsecrets not installed")
	}

	//we get the current directory because we will need to get a file in it later (forbidden_patterns)
	dirMeta, error := os.Getwd()
	if error != nil {
		fmt.Println(error.Error())
		return g, errors.New("can't get project's repertory")
	}

	//installation of hooks for gitsecrets
	cmd = exec.Command("cmd", "/C", "git", "secrets", "--install")
	cmd.Dir = g.path
	_, error = cmd.Output()
	if error != nil {
		fmt.Println(error.Error())
		//return g, errors.New("can't install hook")
	}

	//we add the forbidden patterns
	cmd = exec.Command("cmd", "/C", "git", "secrets", "--add-provider", "--", "cat", fmt.Sprintf("%s\\scanners\\GitSecrets\\forbidden_patterns", dirMeta))
	cmd.Dir = g.path
	_, error = cmd.Output()
	if error != nil {
		fmt.Println(error.Error())
		//return g, errors.New("can't add forbidden patterns' provider")
	}
	return g, nil
}

func (g GitSecrets) Scan() string {
	//we scan the whole repo
	cmd := exec.Command("cmd", "/C", "git", "secrets", "--scan", "-r")
	cmd.Dir = g.path
	//the result is in the error stream
	var errb bytes.Buffer
	cmd.Stderr = &errb
	//we run the command and return the result
	cmd.Run()
	return errb.String()
}
