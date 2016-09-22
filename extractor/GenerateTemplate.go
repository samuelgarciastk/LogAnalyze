package extractor

import (
	. "LogAnalyze/common"
	"LogAnalyze/sql"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

//EraseParameters used to extract the raw log key
//in and out are file paths
func GenerateTemplate(in string) {
	file, err := ioutil.ReadFile(in)
	CheckErr(err, "File not found!")
	fmt.Println("Open file: " + in)
	lines := strings.Split(string(file), "\n")
	var events []string
	for _, line := range lines {
		res, ok := erasingPipeline(line)
		if ok == false {
			continue
		}
		events = append(events, res)
	}
	sql.Insert(events)
	fmt.Println(len(events))
}

func erasingPipeline(line string) (string, bool) {
	res, ok := extractMsg(line)
	if ok == false {
		return "", false
	}
	res = eraseBrace(res)
	res = eraseColon(res)
	res = eraseEqual(res)
	res = eraseNum(res)
	res = trim(res)
	return res, true
}

func extractMsg(line string) (string, bool) {
	reg := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}\.\d{3}.*?(INFO|WARN|ERROR) (.*? - .*)`)
	res := reg.FindStringSubmatch(line)
	if len(res) < 1 {
		return "", false
	}
	return res[2], true
}

func eraseBrace(line string) string {
	reg := regexp.MustCompile(`\{[^\{\}]*?}|\[[^\[\]]*?]`)
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

func eraseColon(line string) string {
	reg := regexp.MustCompile(`(:|：).*?($|,|，| )`)
	return reg.ReplaceAllString(line, " ")
}

func eraseEqual(line string) string {
	reg := regexp.MustCompile(`([^=])=[^=]*?($|,|，| )`)
	return reg.ReplaceAllString(line, "$1 ")
}

func eraseNum(line string) string {
	reg := regexp.MustCompile(`\d+\.?\d* | \d+\.?\d*`)
	parts := strings.Split(line, " - ")
	return parts[0] + " - " + reg.ReplaceAllString(parts[1], " ")
}

func trim(line string) string {
	reg := regexp.MustCompile(` +|	+`)
	return strings.Trim(reg.ReplaceAllString(line, " "), " ")
}
