/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

type _Wm struct {
	WParam win.WPARAM
	LParam win.LPARAM
}

// Raw window message parameters.
type Wm struct {
	_Wm
}

type WmActivate struct{ _Wm }

func (p WmActivate) Event() co.WA                           { return co.WA(p.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                      { return p.WParam.HiWord() != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.LParam) }

type WmActivateApp struct{ _Wm }

func (p WmActivateApp) IsBeingActivated() bool { return p.WParam != 0 }
func (p WmActivateApp) ThreadId() uint32       { return uint32(p.LParam) }

type WmAppCommand struct{ _Wm }

func (p WmAppCommand) OwnerWindow() win.HWND     { return win.HWND(p.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.LParam.HiWord() &^ 0xF000) }
func (p WmAppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.LParam.HiWord() & 0xF000) }
func (p WmAppCommand) Keys() co.MK               { return co.MK(p.LParam.LoWord()) }

type _WmChar struct{ _Wm } // base for other messages

func (p _WmChar) CharCode() uint16          { return uint16(p.WParam) }
func (p _WmChar) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p _WmChar) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p _WmChar) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p _WmChar) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p _WmChar) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }
func (p _WmChar) KeyBeingReleased() bool    { return (p.LParam.HiByteHiWord() & 0b1000_0000) != 0 }

type WmChar struct{ _WmChar }

type WmCommand struct{ _Wm }

func (p WmCommand) IsFromMenu() bool         { return p.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return p.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() int32            { return p.ControlId() }
func (p WmCommand) AcceleratorId() int32     { return p.ControlId() }
func (p WmCommand) ControlId() int32         { return int32(p.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() uint16 { return p.WParam.HiWord() }
func (p WmCommand) ControlHwnd() win.HWND    { return win.HWND(p.LParam) }

type WmContextMenu struct{ _Wm }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.LParam.MakePoint() }

type WmCreate struct{ _Wm }

func (p WmCreate) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.LParam))
}

type WmDeadChar struct{ _WmChar }

type WmDropFiles struct{ _Wm }

func (p WmDropFiles) Hdrop() win.HDROP { return win.HDROP(p.WParam) }

// Calls DragQueryFile successively to retrieve all file names, and releases the
// HDROP handle.
func (p WmDropFiles) RetrieveAll() []string {
	count := p.Hdrop().DragQueryFile(0xFFFFFFFF, nil, 0)
	files := make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		pathLen := p.Hdrop().DragQueryFile(i, nil, 0) + 1 // room for terminating null
		pathBuf := make([]uint16, pathLen)
		p.Hdrop().DragQueryFile(i, &pathBuf[0], pathLen)
		files = append(files, syscall.UTF16ToString(pathBuf))
	}
	p.Hdrop().DragFinish()
	sort.Slice(files, func(i, j int) bool { // case insensitive
		return strings.ToUpper(files[i]) < strings.ToUpper(files[j])
	})
	return files
}

type WmGetDlgCode struct{ _Wm }

func (p WmGetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.WParam) }
func (p WmGetDlgCode) IsQuery() bool         { return p.LParam == 0 }
func (p WmGetDlgCode) Msg() *win.MSG         { return (*win.MSG)(unsafe.Pointer(p.LParam)) }
func (p WmGetDlgCode) HasAlt() bool          { return (win.GetAsyncKeyState(co.VK_MENU) & 0x8000) != 0 }
func (p WmGetDlgCode) HasCtrl() bool         { return (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0 }
func (p WmGetDlgCode) HasShift() bool        { return (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0 }

type WmHelp struct{ _Wm }

func (p WmHelp) HelpInfo() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.LParam)) }

type WmHotKey struct{ _Wm }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.LParam.HiWord()) }

type WmInitMenuPopup struct{ _Wm }

func (p WmInitMenuPopup) Hmenu() win.HMENU        { return win.HMENU(p.WParam) }
func (p WmInitMenuPopup) MenuRelativePos() uint16 { return p.LParam.LoWord() }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return p.LParam.HiWord() != 0 }

type _WmKey struct{ _Wm } // base for other messages

func (p _WmKey) VirtualKeyCode() co.VK     { return co.VK(p.WParam) }
func (p _WmKey) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p _WmKey) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p _WmKey) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p _WmKey) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p _WmKey) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }

type WmKeyDown struct{ _WmKey }

type WmKeyUp struct{ _WmKey }

type WmKillFocus struct{ _Wm }

func (p WmKillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.WParam) }

type _WmButton struct{ _Wm } // base for other messages

func (p _WmButton) HasCtrl() bool      { return (co.MK(p.WParam) & co.MK_CONTROL) != 0 }
func (p _WmButton) HasLeftBtn() bool   { return (co.MK(p.WParam) & co.MK_LBUTTON) != 0 }
func (p _WmButton) HasMiddleBtn() bool { return (co.MK(p.WParam) & co.MK_MBUTTON) != 0 }
func (p _WmButton) HasRightBtn() bool  { return (co.MK(p.WParam) & co.MK_RBUTTON) != 0 }
func (p _WmButton) HasShift() bool     { return (co.MK(p.WParam) & co.MK_SHIFT) != 0 }
func (p _WmButton) HasXBtn1() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON1) != 0 }
func (p _WmButton) HasXBtn2() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON2) != 0 }
func (p _WmButton) Pos() win.POINT     { return p.LParam.MakePoint() }

type WmLButtonDblClk struct{ _WmButton }

type WmLButtonDown struct{ _WmButton }

type WmLButtonUp struct{ _WmButton }

type WmMButtonDblClk struct{ _WmButton }

type WmMButtonDown struct{ _WmButton }

type WmMButtonUp struct{ _WmButton }

type WmMenuChar struct{ _Wm }

func (p WmMenuChar) CharCode() uint16      { return p.WParam.LoWord() }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.LParam) }

type WmMenuCommand struct{ _Wm }

func (p WmMenuCommand) ItemIndex() uint16 { return uint16(p.WParam) }
func (p WmMenuCommand) Hmenu() win.HMENU  { return win.HMENU(p.LParam) }

type WmMenuSelect struct{ _Wm }

func (p WmMenuSelect) Item() uint16     { return p.WParam.LoWord() }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.LParam) }

type WmMouseHover struct{ _WmButton }

type WmMouseMove struct{ _WmButton }

type WmMove struct{ _Wm }

func (p WmMove) Pos() win.POINT { return p.LParam.MakePoint() }

type WmNcPaint struct{ _Wm }

func (p WmNcPaint) Hrgn() win.HRGN { return win.HRGN(p.WParam) }

type WmNotify struct{ _Wm }

func (p WmNotify) NmHdr() *win.NMHDR { return (*win.NMHDR)(unsafe.Pointer(p.LParam)) }

type WmPrint struct{ _Wm }

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.LParam) }

type WmRButtonDblClk struct{ _WmButton }

type WmRButtonDown struct{ _WmButton }

type WmRButtonUp struct{ _WmButton }

type WmSetFocus struct{ _Wm }

func (p WmSetFocus) UnfocusedWindow() win.HWND { return win.HWND(p.WParam) }

type WmSetFont struct{ _Wm }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.LParam == 1 }

type WmSize struct{ _Wm }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.LParam.MakeSize() }

type WmSysChar struct{ _WmChar }

type WmSysCommand struct{ _Wm }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.LParam.MakePoint() }

type WmSysDeadChar struct{ _WmChar }

type WmSysKeyDown struct{ _WmKey }

type WmSysKeyUp struct{ _WmKey }