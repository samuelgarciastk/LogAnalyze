package eraseParam

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

//EraseParameters used to extract the raw log key
//@param in and out are file paths
func EraseParameters(in, out string) {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		panic("File not found!\n" + err.Error())
	}
	reg := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}\.\d{3}`)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if res := reg.FindStringIndex(line); res == nil {
			continue
		}
		fmt.Println(erasingPipeline(line))
	}
}

func erasingPipeline(line string) string {
	res := line[13:]
	res = eraseBrace(res)
	return res
}

func eraseBrace(line string) string {
	reg := regexp.MustCompile(`{[^{}]*?}|\[[^\[\]]*?\]`)
	res := line
	for {
		replaced := reg.ReplaceAllString(res, "")
		if replaced == res {
			break
		}
		res = replaced
	}
	return res
}

// func eraseSquareBracket(line string) string {
// 	reg := regexp.MustCompile(`\[[^[]]*?\]`)
// }
