package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

//调用proj1.mycom的定义体
func main() {
	CoInitialize(0)
	defer CoUninitialize()

	handle, err := CreateObject(`proj1.mycom`)
	if err == nil {
		zobj := (*IMycom)(handle)

		fmt.Println("gettime:", zobj.Gettime())
		fmt.Println("gettime1:", zobj.Gettime1())
	} else {
		fmt.Println(err)
	}

}

type IMycomVtbl struct {
	IDispatchVtbl
	Gettime  uintptr
	Gettime1 uintptr
}

type IMycom struct {
	LpVtbl *IMycomVtbl
}

func (obj *IMycom) Gettime() string {
	//var zx *int32 = new(int32)
	var zx unsafe.Pointer
	syscall.Syscall(obj.LpVtbl.Gettime, 2,
		0,
		uintptr(unsafe.Pointer(&zx)),
		0)

	//zxd := *(*DBLSTR)(unsafe.Pointer(uintptr(*zx)))
	zxd := *(*DBLSTR)(zx)

	return zxd.String()
}
func (obj *IMycom) Gettime1() int {
	var zx unsafe.Pointer
	syscall.Syscall(obj.LpVtbl.Gettime1, 2,
		0,
		uintptr(unsafe.Pointer(&zx)),
		0)

	return int(uintptr(zx))
}

type DBLSTR struct {
	Data [40]uint16
}

func (b DBLSTR) String() string {
	z := ""
	for _, zv := range b.Data {
		if zv == 0 {
			break
		}
		z += string(zv)
	}
	return z
}

//mycom 定义体结束

func CreateObject(zclass string) (unsafe.Pointer, error) {
	var zo REFCLSID = new(CLSID)

	if hr := CLSIDFromProgID(zclass, zo); FAILED(hr) {
		return nil, Error{fmt.Sprintf("error:%v", hr)}
	}

	var zobj unsafe.Pointer

	if hr := CoCreateInstance(zo, nil, CLSCTX_LOCAL_SERVER|CLSCTX_INPROC_SERVER,
		IID_IDispatch, &zobj); FAILED(hr) {
		return nil, Error{fmt.Sprintf("error:%v", hr)}
	}
	return zobj, nil
}

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type IID GUID
type CLSID GUID
type REFIID *IID
type REFCLSID *CLSID

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type IUnknown struct {
	LpVtbl *IUnknownVtbl
}

type IDispatch struct {
	LpVtbl *IDispatchVtbl
}

type IDispatchVtbl struct {
	IUnknownVtbl
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

type (
	BOOL    int32
	HRESULT int32
)

const (
	CLSCTX_INPROC_SERVER          = 0x1
	CLSCTX_INPROC_HANDLER         = 0x2
	CLSCTX_LOCAL_SERVER           = 0x4
	CLSCTX_INPROC_SERVER16        = 0x8
	CLSCTX_REMOTE_SERVER          = 0x10
	CLSCTX_INPROC_HANDLER16       = 0x20
	CLSCTX_RESERVED1              = 0x40
	CLSCTX_RESERVED2              = 0x80
	CLSCTX_RESERVED3              = 0x100
	CLSCTX_RESERVED4              = 0x200
	CLSCTX_NO_CODE_DOWNLOAD       = 0x400
	CLSCTX_RESERVED5              = 0x800
	CLSCTX_NO_CUSTOM_MARSHAL      = 0x1000
	CLSCTX_ENABLE_CODE_DOWNLOAD   = 0x2000
	CLSCTX_NO_FAILURE_LOG         = 0x4000
	CLSCTX_DISABLE_AAA            = 0x8000
	CLSCTX_ENABLE_AAA             = 0x10000
	CLSCTX_FROM_DEFAULT_CONTEXT   = 0x20000
	CLSCTX_ACTIVATE_32_BIT_SERVER = 0x40000
	CLSCTX_ACTIVATE_64_BIT_SERVER = 0x80000
	CLSCTX_ENABLE_CLOAKING        = 0x100000
	CLSCTX_PS_DLL                 = 0x80000000
	CLSCTX_ALL                    = CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER | CLSCTX_LOCAL_SERVER | CLSCTX_REMOTE_SERVER
)

var IID_IDispatch REFIID = &IID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
var libole32 uintptr = MustLoadLibrary("ole32.dll")
var clsidFromString uintptr = MustGetProcAddress(libole32, "CLSIDFromString")
var clsidFromProgID uintptr = MustGetProcAddress(libole32, "CLSIDFromProgID")
var coInitialize uintptr = MustGetProcAddress(libole32, "CoInitialize")
var coCreateInstance uintptr = MustGetProcAddress(libole32, "CoCreateInstance")
var coUninitialize uintptr = MustGetProcAddress(libole32, "CoUninitialize")

func CoUninitialize() {
	syscall.Syscall(uintptr(coUninitialize), 0, 0, 0, 0)
}
func FAILED(hr HRESULT) bool {
	return hr < 0
}

func CoCreateInstance(rclsid REFCLSID, pUnkOuter *IUnknown, dwClsContext uint32, riid REFIID, ppv *unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.Syscall6(coCreateInstance, 5,
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppv)),
		0)

	return HRESULT(ret)
}

func CoInitialize(zpre int) HRESULT {
	zret, _, _ := syscall.Syscall(coInitialize, 1,
		uintptr(zpre),
		0, 0)
	return HRESULT(zret)
}

/*Example:
if hr:=CLSIDFromProgID(`proj1.mycom`, zo);!FAILED(hr) {

	fmt.Println("Progid:", *zo)
} else {
	fmt.Println("Error!")
}
*/
func CLSIDFromProgID(z string, zo REFCLSID) HRESULT {
	zret, _, _ := syscall.Syscall(clsidFromProgID, 2,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(z))),
		uintptr(unsafe.Pointer(zo)), 0)
	return HRESULT(zret)

}

/*Example:
var ziidIDispatch REFIID = new(IID)
fmt.Println(CLSIDFromString(`{08AD76BE-923F-45B4-B14B-04008CA6188F}`, ziidIDispatch))
fmt.Println("refiid:", *ziidIDispatch)

*/
func CLSIDFromString(z string, zo REFIID) HRESULT {
	zret, _, _ := syscall.Syscall(clsidFromString, 2,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(z))),
		uintptr(unsafe.Pointer(zo)), 0)
	return HRESULT(zret)
}

func MustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}

	return uintptr(lib)
}

func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}
