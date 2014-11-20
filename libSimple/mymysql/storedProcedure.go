package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
)

func main() {
	simple1()
}

func simple1() {
	conn := mysql.New("tcp", "", "127.0.0.1:3306", "root", "root", "test")
	err := conn.Connect()
	defer conn.Close()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	str := fmt.Sprintf(`set @result = 0`)

	res, err := conn.Start(str)
	if err != nil {
		fmt.Println("存储过程调用失败")
	}

	str = fmt.Sprintf(`call taop(%d,@result)`, 1)

	res, err = conn.Start(str)
	if err != nil {
		fmt.Println("存储过程调用失败")
	}

	str = fmt.Sprintf(`select @result`)

	res, err = conn.Start(str)
	if err != nil {
		fmt.Println("存储过程调用失败")
	}
	rows, _ := res.GetRows()
	for _, row := range rows {
		fmt.Println(row.Int(0))
	}

}

func simple2() {
	conn := mysql.New("tcp", "", "127.0.0.1:3306", "root", "root", "morefun_accountdata3")
	err := conn.Connect()
	defer conn.Close()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	// str := fmt.Sprintf(`set @iisfcm = 1`)
	// res, err := conn.Start(str)
	// if err != nil {
	// 	fmt.Println("存储过程调用出错", err.Error())
	// }
	// str = fmt.Sprintf(`set @iserverid = 0`)
	// res, err = conn.Start(str)
	// if err != nil {
	// 	fmt.Println("存储过程调用出错", err.Error())
	// }
	str := fmt.Sprintf(`call mfxy_accountlogin(2,"fsdferf2er","1111111111","","haha",0,9,0)`)
	res, err := conn.Start(str)
	if err != nil {
		fmt.Println("存储过程调用出错", err.Error())
	}

	row, _ := res.GetRow()
	fmt.Println(row.Str(0))
}
