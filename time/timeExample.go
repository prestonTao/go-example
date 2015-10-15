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
	eTest7()
}

/*
	得到一个一秒钟的心跳
*/
func heartbeat() {
	for range time.NewTicker(1 * time.Second).C {
		fmt.Println("haha")
	}
}

//在一定时间之后做某个事情
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

//计算一段时间相隔秒数
func eTest6() {
	t := time.Now()
	time.Sleep(time.Second)
	interval := time.Now().Sub(t)
	done := interval > 1
	fmt.Println("---", done)

}

/*
	计算当天半夜12点
*/
func eTest7() {
	str := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(str[:10])

	newTime, _ := time.Parse("2006-01-02 15:04:05", str[:10]+" 00:00:00")
	fmt.Println(newTime.Add(time.Hour * 24))
}
