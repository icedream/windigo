package ui

import (
	"unsafe"
	"wingows/api"
	c "wingows/consts"
)

type windowMainSetup struct {
	wasInit bool // default to false

	ClassName        string      // Optional, defaults to a hash generated by WNDCLASSEX parameters. Passed to RegisterClassEx.
	ClassStyle       c.CS        // Window class style, passed to RegisterClassEx.
	HIcon            api.HICON   // Icon associated with the window, passed to RegisterClassEx.
	HCursor          api.HCURSOR // Window cursor, passed to RegisterClassEx.
	HBrushBackground api.HBRUSH  // Window background brush, passed to RegisterClassEx.
	HIconSmall       api.HICON   // Small icon associated with the window, passed to RegisterClassEx.

	Title   string    // The title of the window, passed to CreateWindowEx.
	Width   uint32    // Initial width of the window, passed to CreateWindowEx.
	Height  uint32    // Initial height of the window, passed to CreateWindowEx.
	Style   c.WS      // Window style, passed to CreateWindowEx.
	ExStyle c.WS_EX   // Window extended style, passed to CreateWindowEx.
	HMenu   api.HMENU // Main window menu, passed to CreateWindowEx.

	CmdShow c.SW // Passed to ShowWindow, defaults to SW_SHOW.
}

func (me *windowMainSetup) initOnce() {
	if !me.wasInit {
		me.wasInit = true

		me.ClassStyle = c.CS_DBLCLKS

		me.Width = 600 // arbitrary dimensions
		me.Height = 500
		me.Style = c.WS_CAPTION | c.WS_SYSMENU | c.WS_CLIPCHILDREN | c.WS_BORDER
		me.ExStyle = c.WS_EX(0)

		me.CmdShow = c.SW_SHOW
	}
}

func (me *windowMainSetup) genWndClassEx(hInst api.HINSTANCE) *api.WNDCLASSEX {
	wcx := api.WNDCLASSEX{}

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.HInstance = hInst
	wcx.Style = me.ClassStyle
	wcx.HIcon = me.HIcon
	wcx.HIconSm = me.HIconSmall

	if me.HCursor != 0 {
		wcx.HCursor = me.HCursor
	} else {
		wcx.HCursor = api.HINSTANCE(0).LoadCursor(c.IDC_ARROW)
	}

	if me.HBrushBackground != 0 {
		wcx.HbrBackground = me.HBrushBackground
	} else {
		wcx.HbrBackground = api.NewBrushFromSysColor(c.COLOR_BTNFACE)
	}

	if me.ClassName == "" {
		me.ClassName = wcx.Hash() // generate hash after all other fields are set
	}
	wcx.LpszClassName = api.StrToUtf16Ptr(me.ClassName)

	return &wcx
}
