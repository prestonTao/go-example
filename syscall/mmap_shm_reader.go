// 上面的程序映射了一块4G的虚拟内存，用来证明mmap没有实际占用4G内存，而是用到了虚拟内存。
// shm_writer创建好共享内存以后，往内存区域写入了一个结构体，shm_reader则读出一个结构体。
// 内嵌的C代码中有一行 ：
// #cgo linux LDFLAGS: -lrt
// 因为mmap在Mac上不需要连接librt，在linux上则需要，所以做了一个条件链接，这是cgo提供的功能。
// 上面代码中还用到一个cgo的技巧，像shm_open和mmap函数在错误时会返回errno，如果我们在go中使用多返回值语法，cgo会自己把错误码转换成错误信息，很方便的功能。
package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_open(char *name) {
    return shm_open(name, O_RDWR);
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

const SHM_NAME = "my_shm"
const SHM_SIZE = 4 * 1000 * 1000 * 1000

type MyData struct {
    Col1 int
    Col2 int
    Col3 int
}

func main() {
    fd, err := C.my_shm_open(C.CString(SHM_NAME))
    if err != nil {
        fmt.Println(err)
        return
    }

    ptr, err := C.mmap(nil, SHM_SIZE, C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
    if err != nil {
        fmt.Println(err)
        return
    }
    C.close(fd)

    data := (*MyData)(unsafe.Pointer(ptr))

    fmt.Println(data)
}