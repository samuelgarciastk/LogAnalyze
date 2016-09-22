package extractor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

//EraseParameters used to extract the raw log key
//@param in and out are file paths
func GenerateTemplate(in, out string) {
	file, err := ioutil.ReadFile(in)
	if err != nil {
		panic("File not found!\n" + err.Error())
	}
	fmt.Println("Open file: " + in)
	lines := strings.Split(string(file), "\n")
	template := make(map[string]bool)
	for _, line := range lines {
		res, ok := erasingPipeline(line)
		if ok == false {
			continue
		}
		template[res] = true
	}
	output(template, out)
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
	reg := regexp.MustCompile(`\d+\.?\d*`)
	parts := strings.Split(line, "-")
	return parts[0] + "-" + reg.ReplaceAllString(parts[1], " ")
}

func trim(line string) string {
	reg := regexp.MustCompile(` +`)
	return strings.Trim(reg.ReplaceAllString(line, " "), " ")
}

func output(template map[string]bool, out string) {
	var buf bytes.Buffer
	for k := range template {
		buf.WriteString(k)
		buf.WriteString("\n")
	}
	err := ioutil.WriteFile(out, buf.Bytes(), 0666)
	if err != nil {
		panic("Write file error!\n" + err.Error())
	}
}
