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

	// time.Sleep(time.Second * time.Duration(2))
	// eTest9()

	start := time.Now()
	str := start.Format("2006-01-02")
	newTime, _ := time.Parse("2006-01-02", str)
	end := newTime.Unix() + int64(60*60*24)

	str = newTime.Format("2006-01-02 15:04:05")
	fmt.Println(str, newTime.Unix(), start.Unix(), end)
	//		fmt.Println(newTime.)
	fmt.Println((end - time.Now().Unix()) / 60 / 60)

	// time.Sleep(time.Duration((end - time.Now().Unix())) * time.Second)
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
	创建一个整天时间
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
<<<<<<< HEAD
=======

>>>>>>> 0b1e128aa7ba22081d97fdaa35fd5bc8117f93ac
}

/*
	判断一个时间点是否在一个时间段
*/
func eTest8() {
	// t := time.Now()
	// fmt.Println(time.Date(2015, 11, 27, 13, 5, 26, 0, time.Local))
	point := time.Date(2015, 11, 27, 13, 5, 26, 0, time.Local)
	fmt.Println("需要判断的时间点：", point)

	start := time.Date(2015, 11, 27, 0, 0, 0, 0, time.Local)
	fmt.Println("开始时间点：", start)

	end := time.Date(2015, 11, 28, 0, 0, 0, 0, time.Local)
	fmt.Println("结束时间点：", end)

	if point.Unix() >= start.Unix() && point.Unix() < end.Unix() {
		fmt.Println("在这个时间段")
	} else {
		fmt.Println("不在这个时间段")
	}
}

/*
	半夜12点执行一次
*/
func eTest9() {
	Loop()
}

/*
	每天半夜12点清空所有验证码
*/
func Loop() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	tomorrow := today.Add(time.Hour * 24)
	disparity := tomorrow.Sub(now)

	time.Sleep(disparity)
	//do somethin
	// codeMap = make(map[string][]CodeOne, 0)
	go Loop()

}
