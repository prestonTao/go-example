package main

import (
	_ "../../lib/mysql-master"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// insertDB()
	New()
}

func createTable() {

}
func insertDB() {
	db, err := sql.Open("mysql", "root:root@/meteorological?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare(`INSERT t_soilmoisture SET CID=?,CNODENUMBER=?,CDATE=?,CAIRTEMPERATURE=?,CAIRMOISTURE=?,
		CSOILTEMPERATURE=?,CSOILMOISTURE=?,CCO2CONCENTRATIONS=?,CLUX=?,CWINDSPEED=?,CPRECIPITATIONRAIN=?`)
	checkErr(err)
	res, err := stmt.Exec(1, 6001, "2014-03-12 16:01:24", 20.27, 24.22, 8.56, 46.92, 0.00, 0.00, 0.00, 0.00)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	log.Println(id)
	db.Close()
}

func New() {
	importData := ImportData{line: make(chan string, 100), over: make(chan string, 1)}
	go importData.readText()
	importData.analysisLine()
	over := <-importData.over
	over = <-importData.over
	log.Println(over)
}

type ImportData struct {
	line chan string //读到的行
	over chan string //处理完成
}

func (this *ImportData) readText() {
	f, e := os.Open("station1.txt")
	if e != nil {
		log.Println(e)
	}
	defer f.Close()
	bytes := make([]byte, 10240000) //400M
	n, e := f.Read(bytes)
	if e != nil {
		fmt.Println(e)
	}
	var line string
	oldIndex := 0
	str := bytes[:n+1]
	for i, r := range str {
		// log.Println(i)
		if r == 13 {
			line = string(str[oldIndex:i])
			this.line <- line
		}
		if r == 10 {
			oldIndex = i
		}
		if r == 0 {
			line = string(str[oldIndex:i])
			this.line <- line
		}
	}
	this.over <- "ok"
}
func (this *ImportData) analysisLine() {
	lineNumber := 0
	for {
		line := <-this.line
		strs := strings.Split(line, ",")
		if _, e := strconv.Atoi(strs[0]); e != nil {
			if lineNumber != 0 {
				log.Println("这里有中文，行号为：", lineNumber)
				return
			}
		}
		fmt.Println(line)
	}
}

type SoilMoisture struct {
	id                int
	airMoisture       float32 // 空气水分
	airTemperature    float32 // 空气温度
	co2concentrations float32 // 二氧化碳浓度
	date              string  //时间
	lux               float32 // 光照
	nodeNumber        int     //节点编号
	precipitationRain float32 // 降水量
	soilMoisture      float32 // 土壤水分
	soilTemperature   float32 // 土壤温度
	windSpeed         float32 // 风速
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
