// ConstomRouter project main.go
package main

import (
	"fmt"
	//"time"
	//"reflect"
	//"path"
	//"path/filepath"
	//"strings"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"reflect"
	"unsafe"
)

var ()

func init() {
}

func main() {

	hash := sha256.New()
	hash.Write([]byte("preston"))
	md := hash.Sum(nil)
	nodeId := hex.EncodeToString(md)
	fmt.Println("nodeId为：", nodeId)

	// x := new(big.Int).SetBytes([]byte(nodeId))
	// x := new(big.Int).SetBytes([]byte("________________________________________________________________"))
	// x := new(big.Int).SetBytes([]byte("////////////////////////////////////////////////////////////////"))
	x := new(big.Int).SetBytes([]byte("a"))
	fmt.Println(x.String())
	fmt.Println(x.BitLen() - 1)
	temp := big.NewInt(1)
	//向左位移操作
	temp = temp.Lsh(temp, uint(x.BitLen()-1))
	fmt.Println(temp)
	//与 操作
	y := new(big.Int).And(temp, x)

	cmp := temp.Cmp(y)
	fmt.Println(cmp)
	//异或操作
	fmt.Println(new(big.Int).AndNot(x, temp).String())
	//异或操作后就去掉了最高位
	fmt.Println(new(big.Int).AndNot(x, temp).BitLen() - 1)

	s := []byte("http://golang.org/pkg/encoding/base64/#variables")

	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(s))

	fmt.Println("==================================================")
	fmt.Println("bigInt为0的二进制长度为：", big.NewInt(0).BitLen())

	fmt.Println("==================================================")
	haha := new(big.Int).SetBytes([]byte("h"))
	fmt.Println(haha.String())
	fmt.Println("haha长度为：", haha.BitLen())
	temp = big.NewInt(1)
	temp = temp.Lsh(temp, uint(haha.BitLen()-1))
	haha = new(big.Int).AndNot(haha, temp)
	fmt.Println(haha.String())

	a := new(big.Int).SetBytes([]byte("a"))
	fmt.Println(a.String())
	Print(a)
	b := new(big.Int).SetBytes([]byte("b"))
	fmt.Println(b.String())
	Print(b)
	c := new(big.Int).SetBytes([]byte("c"))
	fmt.Println(c.String())
	Print(c)
	d := new(big.Int).SetBytes([]byte("d"))
	fmt.Println(d.String())
	Print(d)
	e := new(big.Int).SetBytes([]byte("e"))
	fmt.Println(e.String())
	Print(e)
	f := new(big.Int).SetBytes([]byte("f"))
	fmt.Println(f.String())
	Print(f)
	tempInt := big.NewInt(1)
	tempInt = tempInt.Lsh(tempInt, uint(6))

	Print(tempInt)

	fmt.Println("/////////////////////////////////////////")
	fmt.Println(new(big.Int).Xor(b, c).BitLen())
	fmt.Println("/////////////////////////////////////////")
	fmt.Println(new(big.Int).Xor(new(big.Int).SetBytes([]byte("27511e620b42e8fbec37edf4bfc765d490f326137a40a51837184f61b8aae39f")),
		new(big.Int).SetBytes([]byte("27511e620b42e8fbec37edf4bfc765d490f326137a40a51837184f61b8aae39c"))).BitLen())

	fmt.Println(len([]string{}))

	fmt.Println("++++++++++++++++++++++++++++")
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:2]
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:3]
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr2:", arr2)

	fmt.Println("++++++++++++++++++++++++++++")
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3[:2], arr3[2:])
	fmt.Println(arr3)
	for i, a := range arr3 {
		fmt.Println(i, a)
	}

	ids := []string{}
	ids = append(ids, "nihao")
	for i, value := range ids {
		fmt.Println(i, value)
	}

	fmt.Println(new(big.Int).SetBytes([]byte("d")).String())
	fmt.Println(new(big.Int).Xor(new(big.Int).SetBytes([]byte("d")), big.NewInt(1)))

	testNodeId, err := new(big.Int).SetString("5981959483235454346322815529427907596121900416505857619628533582499268460193490593175551753472696668660452304293204513519804205382177707412299303284324709", 10)
	if err {
		fmt.Println(err)
	}
	fmt.Println(string(testNodeId.Bytes()))

	ChineseByte("法克")
}

func Print(findInt *big.Int) {
	fmt.Println("==================================")
	bi := ""

	// findInt := new(big.Int).SetBytes([]byte(nodeId))
	lenght := findInt.BitLen()
	for i := 0; i < lenght; i++ {
		tempInt := findInt
		findInt = new(big.Int).Div(tempInt, big.NewInt(2))
		mod := new(big.Int).Mod(tempInt, big.NewInt(2))
		bi = mod.String() + bi
	}
	fmt.Println(bi)
	fmt.Println("==================================")
}

func ChineseByte(str string) {
	// a := "abc"                          // 字符串 "abc"
	b := (*uintptr)(unsafe.Pointer(&str)) // b 存储 a 的地址
	var c []byte
	// d 将 c 的结构用 reflect.SliceHeader 表示
	d := (*reflect.SliceHeader)((unsafe.Pointer(&c)))
	d.Cap = len(str)
	d.Len = len(str)
	d.Data = *b    // *b 存储字符串首元素地址
	fmt.Println(c) // [97 98 99]
}
