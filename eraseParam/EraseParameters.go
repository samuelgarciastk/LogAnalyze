package eraseParam

import (
	"LogAnalyze/common"
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
		line = braceErase(line[13:])
		fmt.Println(line)
	}
}

func braceErase(line string) string {
	var stack common.Stack
	for i, v := range line {
		switch v {
		case '{':
			stack.Push(i)
		case '}':
			index, ok := stack.Pop()
			if ok == false {
				continue
			}
			line = line[:index.(int)] + line[i+1:]
		}
	}
	return line
}
