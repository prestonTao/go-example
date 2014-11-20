package main

import (
	"../mgo"
	"../mgo/bson"
	"fmt"
	"os"
	//"strings"
)

func main() {
	//importData()
	find()
}

func find() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("users")

	result := User{}
	err = c.Find(bson.M{"email": "chenmeng@dist.com.cn"}).One(&result)
	//result := []User{}
	//err = c.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

//func findTest(){

//}

func importData() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("users")

	//###########################################################

	//f, e := os.Open("D:/ProgramFiles/mongodb-win32-x86_64-2.0.7/2000W/test.csv")
	f, e := os.Open("D:/workspaces/go/mongoDB/hotel2000WdataImport/test.csv")

	if e != nil {
		fmt.Println(e)
	}

	defer f.Close()
	//bytes := make([]byte, 400000000) //400M
	bytes := make([]byte, 10240) //400M

	n, e := f.Read(bytes)
	if e != nil {
		fmt.Println(e)
	}
	runes := []rune(string(bytes[:n+1]))
	countRow := 1
	start := 0
	first := true
	rowNum := 1
	user := new(User)
	for i, r := range runes {
		//fmt.Println(i, " rune:", r, "---------string:", string(r))
		if r == 13 || r == 10 {
			if r == 10 {
				//fmt.Println(r)
				continue
			}
			if first {
				first = false
			} else {
				value := string(runes[start:i])
				user.Id = value
				//fmt.Println("Id:", value)
				fmt.Println("插入第：", countRow, "条数据")
				countRow++
				c.Insert(user)
			}
			user = new(User)
			start = i + 1
			rowNum = 1
			continue
		}
		if string(r) == "," {
			if start == i {
				start = i + 1
				rowNum += 1
				continue
			}
			//value := strings.Trim(string(runes[start:i]), " ")
			value := string(runes[start:i])
			//fmt.Println(value)
			switch rowNum {
			case 1:
				user.Name = value
				//fmt.Println("Name:", value)
			case 2:
				user.CardNo = value
				//fmt.Println("CardNo:", value)
			case 3:
				user.Descriot = value
				//fmt.Println("Descriot:", value)
			case 4:
				user.CtfTp = value
				//fmt.Println("CtfTp:", value)
			case 5:
				user.CtfId = value
				//fmt.Println("CtfId:", value)
			case 6:
				user.Gender = value
				//fmt.Println("Gender:", value)
			case 7:
				user.Birthday = value
				//fmt.Println("Birthday:", value)
			case 8:
				user.Address = value
				//fmt.Println("Address:", value)
			case 9:
				user.Zip = value
				//fmt.Println("Zip:", value)
			case 10:
				user.Dirty = value
				//fmt.Println("Dirty:", value)
			case 11:
				user.District1 = value
				//fmt.Println("District1:", value)
			case 12:
				user.District2 = value
				//fmt.Println("District2:", value)
			case 13:
				user.District3 = value
				//fmt.Println("District3:", value)
			case 14:
				user.District4 = value
				//fmt.Println("District4:", value)
			case 15:
				user.District5 = value
				//fmt.Println("District5:", value)
			case 16:
				user.District6 = value
				//fmt.Println("District6:", value)
			case 17:
				user.FirstNm = value
				//fmt.Println("FirstNm:", value)
			case 18:
				user.LastNm = value
				//fmt.Println("LastNm:", value)
			case 19:
				user.Duty = value
				//fmt.Println("Duty:", value)
			case 20:
				user.Mobile = value
				//fmt.Println("Mobile:", value)
			case 21:
				user.Tel = value
				//fmt.Println("Tel:", value)
			case 22:
				user.Fax = value
				//fmt.Println("Fax:", value)
			case 23:
				user.EMail = value
				//fmt.Println("EMail:", value)
			case 24:
				user.Nation = value
				//fmt.Println("Nation:", value)
			case 25:
				user.Taste = value
				//fmt.Println("Taste:", value)
			case 26:
				user.Education = value
				//fmt.Println("Education:", value)
			case 27:
				user.Company = value
				//fmt.Println("Company:", value)
			case 28:
				user.CTel = value
				//fmt.Println("CTel:", value)
			case 29:
				user.CAddress = value
				//fmt.Println("CAddress:", value)
			case 30:
				user.CZip = value
				//fmt.Println("CZip:", value)
			case 31:
				user.Family = value
				//fmt.Println("Family:", value)
			case 32:
				user.Version = value
				//fmt.Println("Version:", value)
			case 33:
				user.Id = value
				//fmt.Println("Id:", value)
			}
			start = i + 1
			rowNum += 1
		}
		if r == 0 {
			value := string(runes[start:i])
			user.Id = value
			//fmt.Println("Id:", value)
			fmt.Println("插入第：", countRow, "条数据")
			countRow++
			c.Insert(user)
		}
	} //end for

	//for i := 0; i < len(users); i++ {
	//	fmt.Println(users[i])
	//}
	fmt.Println("csv文件分析完成")
	//######################################

	//for i := 0; i < len(users); i++ {
	//	err = c.Insert(users[i])
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	fmt.Println("插入完成")

}

type User struct {
	Name      string
	CardNo    string
	Descriot  string
	CtfTp     string
	CtfId     string
	Gender    string
	Birthday  string //生日
	Address   string
	Zip       string
	Dirty     string
	District1 string
	District2 string //可能是年龄
	District3 string
	District4 string
	District5 string
	District6 string
	FirstNm   string
	LastNm    string
	Duty      string
	Mobile    string //电话：携程网147541773担保
	Tel       string //座机？
	Fax       string
	EMail     string
	Nation    string //名族
	Taste     string
	Education string
	Company   string
	CTel      string
	CAddress  string
	CZip      string
	Family    string
	Version   string //时间
	Id        string //id
}
