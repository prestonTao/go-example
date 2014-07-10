package main

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"
)

const (
	CLSCTX_INPROC_SERVER = 1
	CLSCTX_LOCAL_SERVER  = 4

	DISPATCH_METHOD      = 1
	DISPATCH_PROPERTYGET = 2
	DISPATCH_PROPERTYPUT = 4

	VT_EMPTY    = 0x0
	VT_NULL     = 0x1
	VT_I2       = 0x2
	VT_I4       = 0x3
	VT_R4       = 0x4
	VT_R8       = 0x5
	VT_CY       = 0x6
	VT_DATE     = 0x7
	VT_BSTR     = 0x8
	VT_DISPATCH = 0x9
	VT_ERROR    = 0xa
	VT_BOOL     = 0xb
	VT_VARIANT  = 0xc
	VT_UNKNOWN  = 0xd
	VT_DECIMAL  = 0xe
	VT_I1       = 0x10
	VT_UI1      = 0x11
	VT_UI2      = 0x12
	VT_UI4      = 0x13
	VT_I8       = 0x14
	VT_UI8      = 0x15
	VT_INT      = 0x16
	VT_UINT     = 0x17
	VT_VOID     = 0x18
	VT_VECTOR   = 0x1000
	VT_ARRAY    = 0x2000
	VT_BYREF    = 0x4000
	VT_RESERVED = 0x8000
)

type DISPPARAMS struct {
	rgvarg            uintptr
	rgdispidNamedArgs uintptr
	cArgs             uint32
	cNamedArgs        uint32
}

type VARIANT struct {
	VT         uint16
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	Val        int64
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type IDispatch struct {
	lpVtbl *pIDispatchVtbl
}

type pIDispatchVtbl struct {
	pQueryInterface   uintptr
	pAddRef           uintptr
	pRelease          uintptr
	pGetTypeInfoCount uintptr
	pGetTypeInfo      uintptr
	pGetIDsOfNames    uintptr
	pInvoke           uintptr
}

var (
	ole32, _    = syscall.LoadLibrary("ole32.dll")
	oleaut32, _ = syscall.LoadLibrary("oleaut32.dll")
	kernel32, _ = syscall.LoadLibrary("kernel32.dll")

	procCoInitialize, _     = syscall.GetProcAddress(ole32, "CoInitialize")
	procCoUninitialize, _   = syscall.GetProcAddress(ole32, "CoUninitialize")
	procCoCreateInstance, _ = syscall.GetProcAddress(ole32, "CoCreateInstance")
	procCLSIDFromProgID, _  = syscall.GetProcAddress(ole32, "CLSIDFromProgID")

	procVariantInit, _    = syscall.GetProcAddress(oleaut32, "VariantInit")
	procSysAllocString, _ = syscall.GetProcAddress(oleaut32, "SysAllocString")
	procSysFreeString, _  = syscall.GetProcAddress(oleaut32, "SysFreeString")

	procGetUserDefaultLCID, _ = syscall.GetProcAddress(kernel32, "GetUserDefaultLCID")

	IID_NULL      = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
	IID_IDispatch = &GUID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
)

func CoInitialize(p uintptr) (hr uintptr) {
	hr, _, _ = syscall.Syscall(uintptr(procCoInitialize), 1, p, 0, 0)
	return
}

func CoUninitialize() {
	syscall.Syscall(uintptr(procCoUninitialize), 0, 0, 0, 0)
}

func CLSIDFromProgID(progId string) (clsid *GUID, hr uintptr) {
	var guid GUID
	hr, _, _ = syscall.Syscall(uintptr(procCLSIDFromProgID), 2,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(progId))),
		uintptr(unsafe.Pointer(&guid)), 0)
	clsid = &guid
	return
}

func CreateInstance(clsid *GUID, iid *GUID) (unk *IDispatch, hr uintptr) {
	hr, _, _ = syscall.Syscall6(uintptr(procCoCreateInstance), 5,
		uintptr(unsafe.Pointer(clsid)),
		0,
		CLSCTX_INPROC_SERVER,
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&unk)),
		0)
	return
}

func (disp *IDispatch) GetIDsOfNames(names []string) (dispid []int32, hr uintptr) {
	wnames := make([]*uint16, len(names))
	for i := 0; i < len(names); i++ {
		wnames[i] = syscall.StringToUTF16Ptr(names[i])
	}
	dispid = make([]int32, len(names))
	hr, _, _ = syscall.Syscall6(disp.lpVtbl.pGetIDsOfNames, 6,
		uintptr(unsafe.Pointer(disp)),
		uintptr(unsafe.Pointer(IID_NULL)),
		uintptr(unsafe.Pointer(&wnames[0])),
		uintptr(len(names)),
		uintptr(GetUserDefaultLCID()),
		uintptr(unsafe.Pointer(&dispid[0])))
	return
}

func (disp *IDispatch) Invoke(dispid int32, dispatch int16, params ...interface{}) (result *VARIANT, hr uintptr) {
	var dispparams DISPPARAMS
	var vargs []VARIANT
	if len(params) > 0 {
		vargs = make([]VARIANT, len(params))
		for i, v := range params {
			n := len(params) - i - 1
			VariantInit(&vargs[n])
			switch v.(type) {
			case bool:
				if v.(bool) {
					vargs[n] = VARIANT{VT_BOOL, 0, 0, 0, 0xffff}
				} else {
					vargs[n] = VARIANT{VT_BOOL, 0, 0, 0, 0}
				}
			case int16:
				vargs[n] = VARIANT{VT_I2, 0, 0, 0, int64(v.(int16))}
			case float32:
				vargs[n] = VARIANT{VT_R4, 0, 0, 0, int64(v.(float32))}
			case float64:
				vargs[n] = VARIANT{VT_R8, 0, 0, 0, int64(v.(float64))}
			case byte:
				vargs[n] = VARIANT{VT_I1, 0, 0, 0, int64(v.(byte))}
			case uint16:
				vargs[n] = VARIANT{VT_UI2, 0, 0, 0, int64(v.(int16))}
			case int, int32:
				vargs[n] = VARIANT{VT_UI4, 0, 0, 0, int64(v.(int))}
			case uint, uint32:
				vargs[n] = VARIANT{VT_UI4, 0, 0, 0, int64(v.(uint))}
			case string:
				vargs[n] = VARIANT{VT_BSTR, 0, 0, 0, int64(uintptr(unsafe.Pointer(SysAllocString(v.(string)))))}
			case *IDispatch:
				vargs[n] = VARIANT{VT_DISPATCH, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*IDispatch))))}
			case *bool:
				vargs[n] = VARIANT{VT_BOOL | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*bool))))}
			case *byte:
				vargs[n] = VARIANT{VT_I1 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*byte))))}
			case *int16:
				vargs[n] = VARIANT{VT_I2 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*int16))))}
			case *uint16:
				vargs[n] = VARIANT{VT_UI2 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*uint16))))}
			case *int, *int32:
				vargs[n] = VARIANT{VT_I4 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*int))))}
			case *uint, *uint32:
				vargs[n] = VARIANT{VT_UI4 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*uint))))}
			case *float32:
				vargs[n] = VARIANT{VT_R4 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*float32))))}
			case *float64:
				vargs[n] = VARIANT{VT_R8 | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*float64))))}
			case *string:
				vargs[n] = VARIANT{VT_BSTR | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*string))))}
			case **IDispatch:
				vargs[n] = VARIANT{VT_DISPATCH | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(**IDispatch))))}
			case *VARIANT:
				vargs[n] = VARIANT{VT_VARIANT | VT_BYREF, 0, 0, 0, int64(uintptr(unsafe.Pointer(v.(*VARIANT))))}
			case nil:
				vargs[n] = VARIANT{VT_NULL, 0, 0, 0, 0}
			default:
				panic("unknown type")
			}
		}
		dispparams.rgvarg = uintptr(unsafe.Pointer(&vargs[0]))
		dispparams.cArgs = uint32(len(params))
	}

	result = new(VARIANT)
	VariantInit(result)
	hr, _, _ = syscall.Syscall9(disp.lpVtbl.pInvoke, 9,
		uintptr(unsafe.Pointer(disp)),
		uintptr(dispid),
		uintptr(unsafe.Pointer(IID_NULL)),
		uintptr(GetUserDefaultLCID()),
		uintptr(dispatch),
		uintptr(unsafe.Pointer(&dispparams)),
		uintptr(unsafe.Pointer(result)),
		0,
		0)

	for _, varg := range vargs {
		if varg.VT == VT_BSTR && varg.Val != 0 {
			SysFreeString(((*int16)(unsafe.Pointer(uintptr(varg.Val)))))
		}
	}
	return
}

func GetUserDefaultLCID() (lcid uint32) {
	ret, _, _ := syscall.Syscall(uintptr(procGetUserDefaultLCID), 0, 0, 0, 0)
	lcid = uint32(ret)
	return
}

func VariantInit(v *VARIANT) (hr uintptr) {
	hr, _, _ = syscall.Syscall(uintptr(procVariantInit), 1, uintptr(unsafe.Pointer(v)), 0, 0)
	return
}

func SysAllocString(v string) (ss *int16) {
	pss, _, _ := syscall.Syscall(uintptr(procSysAllocString), 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v))), 0, 0)
	ss = (*int16)(unsafe.Pointer(pss))
	return
}

func SysFreeString(v *int16) (hr uintptr) {
	hr, _, _ = syscall.Syscall(uintptr(procSysFreeString), 1, uintptr(unsafe.Pointer(v)), 0, 0)
	return
}

func main() {
	CoInitialize(0)

	//clsid, _ := CLSIDFromProgID("Shell.Application")
	//disp, _ := CreateInstance(clsid, IID_IDispatch)
	//dispid, _ := disp.GetIDsOfNames([]string{"BrowseForFolder"})
	//disp.Invoke(dispid[0], DISPATCH_METHOD, 0, "Hello, COM(Go) World!", 0, 36)

	clsid, _ := CLSIDFromProgID("proj1.mycom")
	disp, _ := CreateInstance(clsid, IID_IDispatch)
	dispid, _ := disp.GetIDsOfNames([]string{"gettime"})

	fmt.Println(dispid)
	z, _ := disp.Invoke(dispid[0], DISPATCH_METHOD)

	fmt.Println(reflect.TypeOf(z))

	zx := *(*[40]byte)(unsafe.Pointer(uintptr(z.Val)))
	fmt.Println("value:", string(zx[0:40]))
	CoUninitialize()
}

type BX struct {
	Data [40]byte
}

func (b BX) String() string {
	return ""
}
