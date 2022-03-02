package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienummediatypes
type IEnumMediaTypes interface {
	com.IUnknown

	// ⚠️ You must defer IEnumMediaTypes.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-clone
	Clone() IEnumMediaTypes

	// Calls Skip() until the end of the enum to retrieve the actual number of
	// media types, then calls Reset().
	Count() int

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-next
	Next(mt *AM_MEDIA_TYPE) bool

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-reset
	Reset()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-skip
	Skip(numMediaTypes int) bool
}

type _IEnumMediaTypes struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IEnumMediaTypes.Release().
func NewIEnumMediaTypes(base com.IUnknown) IEnumMediaTypes {
	return &_IEnumMediaTypes{IUnknown: base}
}

func (me *_IEnumMediaTypes) Clone() IEnumMediaTypes {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Clone, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumMediaTypes(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}

func (me *_IEnumMediaTypes) Count() int {
	count := int(0)
	for {
		gotOne := me.Skip(1)
		if gotOne {
			count++
		} else {
			me.Reset()
			return count
		}
	}
}

func (me *_IEnumMediaTypes) Next(mt *AM_MEDIA_TYPE) bool {
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Next, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		1, uintptr(unsafe.Pointer(&mt)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IEnumMediaTypes) Reset() {
	syscall.Syscall(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Reset, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)
}

func (me *_IEnumMediaTypes) Skip(numMediaTypes int) bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Skip, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(uint32(numMediaTypes)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
