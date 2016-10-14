package extractor

import (
	. "LogAnalyze/common"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	location = make(map[string]interface{})
	out      = "/Users/stk/Documents/Projects/Reports/template"
	lock     = sync.Mutex{}
)

//GenerateLocation generates the location list
//'in' is the file path used to generate
func GenerateTemplate(in string) {
	file, err := os.OpenFile(in, os.O_RDONLY, 0666)
	CheckErr(err, "File not found!")
	defer file.Close()
	fmt.Println("Open file: " + in)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res, ok := getLocation(scanner.Text())
		if ok == false {
			continue
		}
		lock.Lock()
		location[res] = true
		lock.Unlock()
		fmt.Println(res)
	}
	fmt.Println(in + ": " + strconv.Itoa(len(location)))
}

func getLocation(line string) (string, bool) {
	reg := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}\.\d{3}.*?(INFO|WARN|ERROR) (.*?) - (.*)`)
	res := reg.FindStringSubmatch(line)
	if len(res) < 1 {
		return "", false
	}
	return strings.Trim(res[2], " "), true
}

func WriteFile() {
	file, err := os.Create(out)
	CheckErr(err, "Create file failed!")
	defer file.Close()
	fmt.Println("Begin write file")
	bw := bufio.NewWriter(file)
	for key := range location {
		bw.WriteString(key)
		bw.WriteString("\n")
	}
	bw.Flush()
	fmt.Println("Write file end")
}
