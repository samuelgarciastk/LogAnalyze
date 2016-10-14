package main

import (
	. "LogAnalyze/common"
	"LogAnalyze/extractor"
	"fmt"
	"sync"
	"time"
)

var dirs = []string{"/Users/stk/Documents/Projects/RawLog/sit10.1.5.9/creditloan", "/Users/stk/Documents/Projects/RawLog/sit10.1.5.84/creditloantask"}

//var dirs = []string{"/Users/stk/Documents/Projects/RawLog/sit10.1.5.9/creditloan"}

func main() {
	start := time.Now()

	//getTemplate()
	getException()

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func getTemplate() {
	fmt.Println("Generate location list")
	paths := ListFile(dirs...)
	//sql.Open()
	//sql.Exec("DROP TABLE IF EXISTS template",
	//"CREATE TABLE template (id INTEGER PRIMARY KEY AUTOINCREMENT,location TEXT NOT NULL UNIQUE,regexp TEXT)")
	var wg sync.WaitGroup
	for _, v := range paths {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			extractor.GenerateTemplate(v)
		}(v)
	}
	wg.Wait()
	//sql.Close()
	extractor.WriteFile()
	fmt.Println("End")
}

func getException() {
	fmt.Println("Generate exception list")
	paths := ListFile(dirs...)
	//var wg sync.WaitGroup
	//for _, v := range paths {
	//	wg.Add(1)
	//	go func(v string) {
	//		defer wg.Done()
	//		extractor.GetException(v)
	//	}(v)
	//}
	//wg.Wait()
	for _, v := range paths {
		extractor.GetException(v)
	}
	extractor.WriteException()
	fmt.Println("End")
}
