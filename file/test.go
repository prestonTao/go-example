package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func WriteFile(content []string) {
	f, err := os.OpenFile("sameName.csv", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	CheckErr(err)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)

	writer.Write(content)

	writer.Flush()
}

var exsit []string = []string{}

func Exsit(name string) bool {
	size := len(exsit)
	for i := 0; i < size; i++ {
		if exsit[i] == name {
			return true
		}
	}
	return false
}

func findSame(path string) {
	fullPath := GetFullPath(path)
	records := getStrings(fullPath)

	size := len(records)
	for i := 0; i < size; i++ {
		rec := records[i]
		recNoDate := rec[:len(rec)-19]
		flag := 1
		for i2 := i + 1; i2 < size; i2++ {
			if recNoDate == records[i2][:len(records[i2])-19] && !Exsit(recNoDate) {
				WriteFile([]string{strconv.Itoa(flag), records[i2]})
				flag++
			}
		}
		if flag > 1 && !Exsit(recNoDate) {
			exsit = append(exsit, recNoDate)
			WriteFile([]string{strconv.Itoa(flag), rec})
		}
	}

	size = len(exsit)
	f, err := os.OpenFile("debug.csv", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	CheckErr(err)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)
	for i := 0; i < size; i++ {
		writer.Write([]string{exsit[i]})
	}
	writer.Flush()
}

func getStrings(path string) []string {
	f, err := os.Open(path)
	CheckErr(err)
	defer f.Close()

	records := []string{}

	reader := csv.NewReader(f)
	record := []string{}
	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		} else if nil != err {
			fmt.Println(err)
			return nil
		}

		records = append(records, record[0])
	}
	return records
}

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path, _ = os.Getwd()
	}
	findSame(path)

	fmt.Println("done!")
}
