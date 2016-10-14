package extractor

import (
	. "LogAnalyze/common"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	outPath = "/Users/stk/Documents/Projects/Reports/exception/"
	data    = make(map[string]map[string]interface{})
)

func GetException(in string) {
	file, err := os.OpenFile(in, os.O_RDONLY, 0666)
	CheckErr(err, "File not found!")
	defer file.Close()
	fmt.Println("Open file: " + in)
	tmp := strings.Split(in, ".")
	name := tmp[len(tmp)-1]

	lock.Lock()
	exception, ok := data[name]
	if !ok {
		exception = make(map[string]interface{})
		data[name] = exception
	}
	lock.Unlock()

	lastMsg := ""
	lastCond := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res, ok := match(scanner.Text())
		if lastCond == 0 && ok == 1 {
			lock.Lock()
			exception[lastMsg] = true
			lock.Unlock()
			//fmt.Println(lastMsg)
		}
		lastMsg = res
		lastCond = ok
	}
}

func match(line string) (string, int) {
	regDate := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}\.\d{3}`)
	if regDate.FindStringIndex(line) == nil {
		return "", 1
	}
	reg := regexp.MustCompile(`.*?(WARN|ERROR|WARNING) (.*?) - `)
	res := reg.FindStringSubmatch(line)
	if len(res) < 1 {
		return "", 2
	}
	return strings.Trim(res[2], " "), 0
}

func WriteException() {
	for date, exception := range data {
		out := outPath + date
		outFile, err := os.Create(out)
		CheckErr(err, "Create file failed!")
		fmt.Println("Create out file: " + out)
		bw := bufio.NewWriter(outFile)
		for key := range exception {
			bw.WriteString(key)
			bw.WriteString("\n")
		}
		bw.Flush()
		fmt.Println("Write file end")
		outFile.Close()
	}
}
