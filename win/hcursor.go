/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"syscall"
	"wingows/co"
	"wingows/win/proc"
)

type HCURSOR HANDLE

func (hCursor HCURSOR) DestroyCursor() {
	ret, _, lerr := syscall.Syscall(proc.DestroyCursor.Addr(), 1,
		uintptr(hCursor), 0, 0)
	if ret == 0 {
		panic(co.ERROR(lerr).Format("DestroyCursor failed."))
	}
}