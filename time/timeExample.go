package main

import (
	"fmt"
	"time"
)

func main() {
	//eTest1()
	//eTest2()
	// eTest3()
	// //睡眠3秒钟
	// time.Sleep(time.Second * 3)

	// eTest5()
	heartbeat()
}

/*
	得到一个一秒钟的心跳
*/
func heartbeat() {
	for range time.NewTicker(1 * time.Second).C {
		fmt.Println("haha")
	}
}

func eTest2() {
	fmt.Println("the 1")
	tc := time.After(time.Second)
	fmt.Println("the 2")
	<-tc
	fmt.Println("the 3")
	fmt.Println(tc)
}

//计算程序运行的时间
func eTest3() {
	fmt.Println("the 1")
	t1 := time.Now()
	fmt.Println("the 2")
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Nanoseconds()) //计算出微毫秒，十亿份之一秒
}

/*
	创建一个时间
*/
func eTest4() {
	t := time.Now()
	fmt.Println(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local))
	fmt.Println(t.Sub(time.Date(0, 0, 1, 0, 0, 0, 0, time.Local)))
}

//计算昨天的这个时候
func eTest5() {
	t := time.Now()
	yesterday := time.Unix(t.Unix()-(60*60*24), int64(t.Nanosecond()))
	fmt.Println(yesterday)
}

func Parse_time() {
	the_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}
}
