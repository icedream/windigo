/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type HINSTANCE HANDLE

// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-duplicateicon
func (hInst HINSTANCE) DuplicateIcon(hIcon HICON) HICON {
	ret, _, _ := syscall.Syscall(proc.DuplicateIcon.Addr(), 2,
		uintptr(hInst), uintptr(hIcon), 0)
	if ret == 0 {
		panic("DuplicateIcon failed.")
	}
	return HICON(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclassinfoexw
func (hInst HINSTANCE) GetClassInfoEx(className *uint16,
	destBuf *WNDCLASSEX) ATOM {

	ret, _, lerr := syscall.Syscall(proc.GetClassInfoEx.Addr(), 3,
		uintptr(hInst),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(destBuf)))
	if ret == 0 {
		panic(co.ERROR(lerr).Format("GetClassInfoEx failed."))
	}
	return ATOM(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
//
// Pass an empty string to get own process handle.
func GetModuleHandle(moduleName string) HINSTANCE {
	ret, _, _ := syscall.Syscall(proc.GetModuleHandle.Addr(), 1,
		uintptr(unsafe.Pointer(StrToPtrBlankIsNil(moduleName))),
		0, 0)
	return HINSTANCE(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadcursorw
func (hInst HINSTANCE) LoadCursor(lpCursorName co.IDC) HCURSOR {
	ret, _, lerr := syscall.Syscall(proc.LoadCursor.Addr(), 2,
		uintptr(hInst), uintptr(lpCursorName), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("LoadCursor failed."))
	}
	return HCURSOR(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadiconw
func (hInst HINSTANCE) LoadIcon(lpIconName co.IDI) HICON {
	ret, _, lerr := syscall.Syscall(proc.LoadIcon.Addr(), 2,
		uintptr(hInst), uintptr(lpIconName), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("LoadIcon failed."))
	}
	return HICON(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadmenuw
func (hInst HINSTANCE) LoadMenu(lpMenuName int32) HMENU {
	ret, _, lerr := syscall.Syscall(proc.LoadMenu.Addr(), 2,
		uintptr(hInst), uintptr(lpMenuName), 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("LoadMenu failed."))
	}
	return HMENU(ret)
}
