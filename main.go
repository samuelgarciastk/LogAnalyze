package main

import (
	. "LogAnalyze/common"
	"LogAnalyze/extractor"
	"LogAnalyze/sql"
	"fmt"
	"sync"
	"time"
)

var dirs = []string{"/Users/stk/Documents/Projects/RawLog/9/creditloan", "/Users/stk/Documents/Projects/RawLog/84/creditloantask"}

//var dirs = []string{"/Users/stk/Documents/Projects/RawLog/84/creditloantask"}
//var dirs = []string{"/Users/stk/Documents/Projects/RawLog/9/creditloan"}

func main() {
	start := time.Now()

	generateLocation()

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func generateLocation() {
	fmt.Println("Generate location list")
	paths := ListFile(dirs...)
	sql.Open()
	sql.Exec("DELETE FROM location")
	var wg sync.WaitGroup
	for _, v := range paths {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			extractor.GenerateLocation(v)
		}(v)
	}
	wg.Wait()
	sql.Close()
	fmt.Println("End")
}

func generateTemplate() {

}

//"UPDATE sqlite_sequence SET seq=0 WHERE name='template'"
