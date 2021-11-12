package parser

import (
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var extmap = make(map[string][]string)
var kicksCheck = "(.*\\.tf$)|(.*\\.json$)|(.*\\.yaml$)|(.*\\.yml$)|(^dockerfile.*)|(.*\\.dockerfile$)"

func visit(path string, di fs.DirEntry, err error) error {
	// skip folder on error
	if err != nil {
		return filepath.SkipDir
	}
	if filepath.Ext(path) != "." && filepath.Ext(path) != ".." {
		match, _ := regexp.Match(kicksCheck, []byte(strings.ToLower(path)))
		if match {
			extmap["kics"] = append(extmap["kics"], path)
		} else if filepath.Ext(path) != "" {
			extmap[filepath.Ext(path)] = append(extmap[filepath.Ext(path)], path)
		}
	}
	return nil
}

func GetFiles(basedir string) map[string][]string {
	// walk all dirs recursively and calls func visit on each one
	err := filepath.WalkDir(basedir, visit)
	if err != nil {
		log.Println(err)
	}
	return extmap
}
