package extractor

import (
	. "LogAnalyze/common"
	"LogAnalyze/sql"
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//GenerateLocation generates the location list
//'in' is the file path used to generate
func GenerateLocation(in string) {
	//file, err := ioutil.ReadFile(in)
	file, err := os.Open(in)
	CheckErr(err, "File not found!")
	defer file.Close()
	fmt.Println("Open file: " + in)

	br := bufio.NewReader(file)
	var location []string
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		res, ok := getLocation(line)
		if ok == false {
			continue
		}
		location = append(location, res)
	}

	//lines := strings.Split(string(file), "\n")
	//for _, line := range lines {
	//	res, ok := getLocation(line)
	//	if ok == false {
	//		continue
	//	}
	//	location = append(location, res)
	//}
	sql.Insert(location, "INSERT INTO location(location) values(?)")
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
