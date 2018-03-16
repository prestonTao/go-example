package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	example2()
}

func example1() {
	path := "file/data_1"

	err := os.Truncate(path, int64(100))
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err.Error())
	}

	size, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(size)

	buf := bytes.NewBuffer([]byte("t"))
	fmt.Println(buf.Bytes())
	file.WriteAt(buf.Bytes(), 40)

}

func example2() {
	buf := bytes.NewBuffer([]byte("taopopootaohongefifnlvnknihaoa"))
	fmt.Println(buf.Bytes(), buf.Len())
	result, e := binary.Varint(buf.Bytes())
	fmt.Println(result, e)
}


//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func example3() {
    //也许大部分基本的文件读取任务是将文件内容读取到内存中。
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))
    //你经常会想对于一个文件是怎么读并且读取到哪一部分进行更多的控制。对于这个任务，从使用 os.Open打开一个文件获取一个 os.File 值开始。
    f, err := os.Open("/tmp/dat")
    check(err)
    //从文件开始位置读取一些字节。这里最多读取 5 个字节，并且这也是我们实际读取的字节数。
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
    //你也可以 Seek 到一个文件中已知的位置并从这个位置开始进行读取。
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
    //io 包提供了一些可以帮助我们进行文件读取的函数。例如，上面的读取可以使用 ReadAtLeast 得到一个更健壮的实现。
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
    //没有内置的回转支持，但是使用 Seek(0, 0) 实现。
    _, err = f.Seek(0, 0)
    check(err)
    //bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
    //任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）。
    f.Close()
}

func example4() {
    //开始，这里是展示如写入一个字符串（或者只是一些字节）到一个文件。
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)
    //对于更细粒度的写入，先打开一个文件。
    f, err := os.Create("/tmp/dat2")
    check(err)
    //打开文件后，习惯立即使用 defer 调用文件的 Close操作。
    defer f.Close()
    //你可以写入你想写入的字节切片
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)
    //WriteString 也是可用的。
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)
    //调用 Sync 来将缓冲区的信息写入磁盘。
    f.Sync()
    //bufio 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)
    //使用 Flush 来确保所有缓存的操作已写入底层写入器。
    w.Flush()
}

/*
运行这端文件写入代码。
$ go run writing-files.go
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes
然后检查写入文件的内容。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered
*/