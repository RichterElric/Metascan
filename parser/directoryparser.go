package parser

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strconv"
)

var g int
var extmap = make(map[string][]string)

func visit(path string, di fs.DirEntry, err error) error {
	// skip folder on error
	if err != nil {
		return filepath.SkipDir
	}
	extmap[filepath.Ext(path)] = append(extmap[filepath.Ext(path)], path)
	g += 1
	return nil
}

func GetFiles(basedir string) {
	var a = &g
	// walk all dirs recursively and calls func visit on each one
	err := filepath.WalkDir(basedir, visit)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("parsed " + strconv.Itoa(*a) + " files")
	fmt.Println("map of ext :\n", extmap)
}
