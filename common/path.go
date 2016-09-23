package common

import (
	"io/ioutil"
	"strings"
)

//ListFile is used to get all file paths(except hidden files) from the given directory
//'path' is the directory path
func ListFile(dirs ...string) []string {
	var paths []string
	for _, path := range dirs {
		file, err := ioutil.ReadDir(path)
		CheckErr(err, "Read file path failed!")
		for _, f := range file {
			if !f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
				paths = append(paths, path+"/"+f.Name())
			}
		}
	}
	return paths
}
