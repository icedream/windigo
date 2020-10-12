/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native list view control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/list-view-controls-overview
type ListView struct {
	_ControlNativeBase
	contextMenu *Menu // if set, will be shown with right-click
}

// Appends a new column, returns the new column.
// Column width will be adjusted to the current system DPI.
func (me *ListView) AddColumn(text string, width uint) *ListViewColumn {
	_, _, width, _ = _Ui.MultiplyDpi(0, 0, width, 0)
	textBuf := win.Str.ToUint16Slice(text)
	lvc := win.LVCOLUMN{
		Mask:    co.LVCF_TEXT | co.LVCF_WIDTH,
		PszText: &textBuf[0],
		Cx:      int32(width),
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTCOLUMN, 0xffff,
		win.LPARAM(unsafe.Pointer(&lvc)))
	if int(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTCOLUMN failed \"%s\".", text))
	}
	return me.Column(uint(newIdx))
}

// Appends many columns at once.
// Column width will be adjusted to the current system DPI.
func (me *ListView) AddColumns(texts []string, widths []uint) *ListView {
	if len(texts) != len(widths) {
		panic("ColumnAdd texts/widths mismatch.")
	}

	for i := 0; i < len(texts); i++ {
		me.AddColumn(texts[i], widths[i])
	}
	return me
}

// Adds a new item, returns the new item.
func (me *ListView) AddItem(text string) *ListViewItem {
	textBuf := win.Str.ToUint16Slice(text)
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT,
		PszText: &textBuf[0],
		IItem:   0x0fff_ffff, // insert as the last one
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTITEM, 0,
		win.LPARAM(unsafe.Pointer(&lvi)))
	if int(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return me.Item(uint(newIdx))
}

// Adds a new item specifying the text of many columns at once.
func (me *ListView) AddItemWithColumns(textsOfEachColumn []string) *ListView {
	newItem := me.AddItem(textsOfEachColumn[0])
	for i := 1; i < len(textsOfEachColumn); i++ {
		newItem.SetSubItemText(uint(i), textsOfEachColumn[i])
	}
	return me
}

// Adds a new item returns the new item.
//
// Before call this method, attach an image list and load its icons.
func (me *ListView) AddItemWithIcon(
	text string, iconIndex int) *ListViewItem {

	textBuf := win.Str.ToUint16Slice(text)
	lvi := win.LVITEM{
		Mask:    co.LVIF_TEXT | co.LVIF_IMAGE,
		PszText: &textBuf[0],
		IImage:  int32(iconIndex),
		IItem:   0x0fff_ffff, // insert as the last one
	}
	newIdx := me.sendLvmMessage(co.LVM_INSERTITEM, 0,
		win.LPARAM(unsafe.Pointer(&lvi)))
	if int(newIdx) == -1 {
		panic(fmt.Sprintf("LVM_INSERTITEM failed \"%s\".", text))
	}
	return me.Item(uint(newIdx))
}

// Adds many items at once.
func (me *ListView) AddItems(texts []string) *ListView {
	for i := range texts {
		me.AddItem(texts[i])
	}
	return me
}

// Returns the column at the given index.
//
// Does not perform bound checking.
func (me *ListView) Column(index uint) *ListViewColumn {
	return &ListViewColumn{
		owner: me,
		index: index,
	}
}

// Calls CreateWindowEx(). This is a basic method: no styles are provided by
// default, you must inform all of them.
//
// Position and size will be adjusted to the current system DPI.
func (me *ListView) Create(
	parent Window, ctrlId, x, y int, width, height uint,
	exStyles co.WS_EX, styles co.WS,
	lvExStyles co.LVS_EX, lvStyles co.LVS) *ListView {

	me.installSubclass()

	x, y, width, height = _Ui.MultiplyDpi(x, y, width, height)
	me._ControlNativeBase.create(exStyles,
		"SysListView32", "", styles|co.WS(lvStyles),
		x, y, width, height, parent, ctrlId)

	if lvExStyles != co.LVS_EX_NONE {
		me.SetExtendedStyle(true, lvExStyles)
	}
	return me
}

// Calls CreateWindowEx() with LVS_REPORT | LVS_NOSORTHEADER | LVS_SHOWSELALWAYS | LVS_SHAREIMAGELISTS.
//
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateReport(
	parent Window, ctrlId, x, y int, width, height uint) *ListView {

	return me.Create(parent, ctrlId, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.LVS_EX_FULLROWSELECT,
		co.LVS_REPORT|co.LVS_NOSORTHEADER|co.LVS_SHOWSELALWAYS|co.LVS_SHAREIMAGELISTS)
}

// Calls CreateWindowEx() with LVS_REPORT | LVS_NOSORTHEADER | LVS_SHOWSELALWAYS | LVS_SORTASCENDING | LVS_SHAREIMAGELISTS.
//
// Position and size will be adjusted to the current system DPI.
func (me *ListView) CreateSortedReport(
	parent Window, ctrlId, x, y int, width, height uint) *ListView {

	return me.Create(parent, ctrlId, x, y, width, height,
		co.WS_EX_CLIENTEDGE,
		co.WS_CHILD|co.WS_GROUP|co.WS_TABSTOP|co.WS_VISIBLE,
		co.LVS_EX_FULLROWSELECT,
		co.LVS_REPORT|co.LVS_NOSORTHEADER|co.LVS_SHOWSELALWAYS|co.LVS_SORTASCENDING|co.LVS_SHAREIMAGELISTS)
}

// Retrieves the number of columns with LVM_GETHEADER and HDM_GETITEMCOUNT.
func (me *ListView) ColumnCount() uint {
	hHeader := win.HWND(me.sendLvmMessage(co.LVM_GETHEADER, 0, 0))
	if hHeader == 0 {
		panic("LVM_GETHEADER failed.")
	}

	count := hHeader.SendMessage(co.WM(co.HDM_GETITEMCOUNT), 0, 0)
	if int(count) == -1 {
		panic("HDM_GETITEMCOUNT failed.")
	}
	return uint(count)
}

// Deletes all items with LVM_DELETEALLITEMS.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-deleteallitems
func (me *ListView) DeleteAllItems() *ListView {
	ret := me.sendLvmMessage(co.LVM_DELETEALLITEMS, 0, 0)
	if ret == 0 {
		panic("LVM_DELETEALLITEMS failed.")
	}
	return me
}

// Deletes many items at once.
//
// Assumes the items are sorted by index.
func (me *ListView) DeleteItems(items []ListViewItem) *ListView {
	for i := len(items) - 1; i >= 0; i-- {
		if items[i].owner != me {
			panic("Cannot delete an item from another list view.")
		}
		items[i].Delete()
	}
	return me
}

// Retrieves extended styles with LVM_GETEXTENDEDLISTVIEWSTYLE.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getextendedlistviewstyle
func (me *ListView) ExtendedStyle() co.LVS_EX {
	return co.LVS_EX(me.sendLvmMessage(co.LVM_GETEXTENDEDLISTVIEWSTYLE, 0, 0))
}

// Searches for an item with the given exact text, case-insensitive.
//
// Returns nil if not found.
func (me *ListView) FindItem(text string) *ListViewItem {
	buf := win.Str.ToUint16Slice(text)
	lvfi := win.LVFINDINFO{
		Flags: co.LVFI_STRING,
		Psz:   &buf[0],
	}
	wp := -1
	idx := me.sendLvmMessage(co.LVM_FINDITEM,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvfi)))
	if int(idx) == -1 {
		return nil // not found
	}
	return &ListViewItem{
		owner: me,
		index: uint(idx),
	}
}

// Retrieves the currently focused item, or nil if none.
func (me *ListView) FocusedItem() *ListViewItem {
	idx := -1
	idx = int(me.sendLvmMessage(co.LVM_GETNEXTITEM,
		win.WPARAM(idx), win.LPARAM(co.LVNI_FOCUSED)))
	if idx == -1 {
		return nil
	}
	return me.Item(uint(idx))
}

// Sends LVM_HITTEST to determine the item at specified position, if any. Pos
// coordinates must be relative to list view.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-hittest
func (me *ListView) HitTest(pos win.POINT) *win.LVHITTESTINFO {
	lvhti := win.LVHITTESTINFO{
		Pt: pos,
	}
	wp := -1 // Vista: retrieve iGroup and iSubItem
	me.sendLvmMessage(co.LVM_HITTEST,
		win.WPARAM(wp), win.LPARAM(unsafe.Pointer(&lvhti)))
	return &lvhti
}

// Retrieves the associated HIMAGELIST by sending LVM_GETIMAGELIST.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getimagelist
func (me *ListView) ImageList(typeImgList co.LVSIL) win.HIMAGELIST {
	return win.HIMAGELIST(
		me.sendLvmMessage(co.LVM_GETIMAGELIST, win.WPARAM(typeImgList), 0),
	)
}

// Sends LVM_ISGROUPVIEWENABLED.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-isgroupviewenabled
func (me *ListView) IsGroupViewEnabled() bool {
	return me.sendLvmMessage(co.LVM_ISGROUPVIEWENABLED, 0, 0) >= 0
}

// Returns the item at the given index.
//
// Does not perform bound checking.
//
// Note: When an item is deleted, all other items may become invalid, because
// they keep the sequential index.
func (me *ListView) Item(index uint) *ListViewItem {
	return &ListViewItem{
		owner: me,
		index: index,
	}
}

// Retrieves the number of items with LVM_GETITEMCOUNT.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getitemcount
func (me *ListView) ItemCount() uint {
	count := me.sendLvmMessage(co.LVM_GETITEMCOUNT, 0, 0)
	if int(count) == -1 {
		panic("LVM_GETITEMCOUNT failed.")
	}
	return uint(count)
}

// Scrolls the contents with LVM_SCROLL.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-scroll
func (me *ListView) Scroll(pxHorz, pxVert int) *ListView {
	ret := me.sendLvmMessage(co.LVM_SCROLL,
		win.WPARAM(pxHorz), win.LPARAM(pxVert))
	if ret == 0 {
		panic("LVM_SCROLL failed.")
	}
	return me
}

// Selects or deselects all items at once.
func (me *ListView) SelectAllItems(isSelected bool) *ListView {
	state := co.LVIS_NONE
	if isSelected {
		state = co.LVIS_SELECTED
	}

	lvi := win.LVITEM{
		State:     state,
		StateMask: co.LVIS_SELECTED,
	}
	idx := -1
	ret := me.sendLvmMessage(co.LVM_SETITEMSTATE,
		win.WPARAM(idx), win.LPARAM(unsafe.Pointer(&lvi)))
	if ret == 0 {
		panic("LVM_SETITEMSTATE failed.")
	}
	return me
}

// Retrieves the number of selected items with LVM_GETSELECTEDCOUNT.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getselectedcount
func (me *ListView) SelectedItemCount() uint {
	count := me.sendLvmMessage(co.LVM_GETSELECTEDCOUNT, 0, 0)
	if int(count) == -1 {
		panic("LVM_GETSELECTEDCOUNT failed.")
	}
	return uint(count)
}

// Retrieves the currently selected items, sorted by index, with LVM_GETNEXTITEM.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getnextitem
func (me *ListView) SelectedItems() []ListViewItem {
	items := make([]ListViewItem, 0, me.SelectedItemCount())
	idx := -1
	for {
		idx = int(me.sendLvmMessage(co.LVM_GETNEXTITEM,
			win.WPARAM(idx), win.LPARAM(co.LVNI_SELECTED)))
		if idx == -1 {
			break
		}
		items = append(items, *me.Item(uint(idx)))
	}
	return items
}

// Retrieves the texts of the selected items, under the given column, if any.
func (me *ListView) SelectedItemTexts(columnIndex uint) []string {
	selItems := me.SelectedItems()
	texts := make([]string, 0, len(selItems))
	for _, item := range selItems {
		texts = append(texts, item.SubItemText(columnIndex))
	}
	return texts
}

// Defines a menu to be shown as the context menu for the list view.
func (me *ListView) SetContextMenu(popupMenu *Menu) *ListView {
	me.contextMenu = popupMenu
	return me
}

// Sets or unsets one or more extended styles with LVM_SETEXTENDEDLISTVIEWSTYLE.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setextendedlistviewstyle
func (me *ListView) SetExtendedStyle(isSet bool, exStyle co.LVS_EX) *ListView {
	mask := exStyle
	if !isSet {
		mask = 0
	}
	me.sendLvmMessage(co.LVM_SETEXTENDEDLISTVIEWSTYLE,
		win.WPARAM(mask), win.LPARAM(exStyle))
	return me
}

// Sends LVM_SETIMAGELIST, returning the previously associated image list.
//
// Note that if the list view was created with LVS_SHAREIMAGELISTS, the image
// list must be manually destroyed.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setimagelist
func (me *ListView) SetImageList(typeImgList co.LVSIL,
	hImgList win.HIMAGELIST) win.HIMAGELIST {

	return win.HIMAGELIST(
		me.sendLvmMessage(co.LVM_SETIMAGELIST,
			win.WPARAM(typeImgList), win.LPARAM(hImgList)),
	)
}

// Sends WM_SETREDRAW to enable or disable UI updates.
//
// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-setredraw
func (me *ListView) SetRedraw(allowRedraw bool) *ListView {
	me.hwnd.SendMessage(co.WM_SETREDRAW,
		win.WPARAM(_Ui.BoolToUint32(allowRedraw)), 0)
	return me
}

// Sets the current view with LVM_SETVIEW.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-setview
func (me *ListView) SetView(view co.LV_VIEW) *ListView {
	if int(me.sendLvmMessage(co.LVM_SETVIEW, 0, 0)) == -1 {
		panic("LVM_SETVIEW failed.")
	}
	return me
}

// Returns the width of a string using list view current font, with
// LVM_GETSTRINGWIDTH.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getstringwidth
func (me *ListView) StringWidth(text string) uint {
	ret := me.sendLvmMessage(co.LVM_GETSTRINGWIDTH,
		0, win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(text))))
	if ret == 0 {
		panic("LVM_GETSTRINGWIDTH failed.")
	}
	return uint(ret)
}

// Retrieves the topmost visible item with LVM_GETTOPINDEX, or nil if none.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-gettopindex
func (me *ListView) TopMostVisibleItem() *ListViewItem {
	idx := int(me.sendLvmMessage(co.LVM_GETTOPINDEX, 0, 0))
	if idx == -1 {
		return nil
	}
	return me.Item(uint(idx))
}

// Retrieves current view with LVM_GETVIEW.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/lvm-getview
func (me *ListView) View() co.LV_VIEW {
	return co.LV_VIEW(me.sendLvmMessage(co.LVM_GETVIEW, 0, 0))
}

// Adds all subclass message handlers; must be called before creation.
func (me *ListView) installSubclass() {
	me.OnSubclassMsg().WmRButtonDown(func(p WmMouse) {
		// WM_RBUTTONUP doesn't work, only NM_RCLICK on parent.
		// https://stackoverflow.com/a/30206896
		me.showContextMenu(true, p.HasCtrl(), p.HasShift())
	})

	me.OnSubclassMsg().WmGetDlgCode(func(p WmGetDlgCode) co.DLGC {
		if !p.IsQuery() && p.VirtualKeyCode() == 'A' && p.HasCtrl() { // Ctrl+A to select all items
			me.SelectAllItems(true)
			return co.DLGC_WANTCHARS

		} else if !p.IsQuery() && p.VirtualKeyCode() == co.VK_RETURN { // send Enter key to parent
			code := co.LVN_KEYDOWN
			nmlvk := win.NMLVKEYDOWN{
				Hdr: win.NMHDR{
					HWndFrom: me.Hwnd(),
					Code:     uint32(code),
					IdFrom:   uintptr(me.Id()),
				},
				WVKey: co.VK_RETURN,
			}
			me.Hwnd().GetAncestor(co.GA_PARENT).
				SendMessage(co.WM_NOTIFY,
					win.WPARAM(me.Hwnd()), win.LPARAM(unsafe.Pointer(&nmlvk)))
			return co.DLGC_WANTALLKEYS

		} else if !p.IsQuery() && p.VirtualKeyCode() == co.VK_APPS { // context menu key
			me.showContextMenu(false, p.HasCtrl(), p.HasShift())
		}

		return co.DLGC(
			me.Hwnd().DefSubclassProc(co.WM_GETDLGCODE,
				p.Raw().WParam, p.Raw().LParam),
		)
	})
}

// Shows the popup menu anchored at cursor pos.
//
// This function will block until the menu disappears.
func (me *ListView) showContextMenu(followCursor, hasCtrl, hasShift bool) {
	if me.contextMenu.Hmenu() == 0 {
		return
	}

	menuPos := win.POINT{} // menu anchor coords, relative to list view

	if followCursor { // usually when fired by a right-click
		menuPos = win.GetCursorPos()         // relative to screen
		me.Hwnd().ScreenToClientPt(&menuPos) // now relative to list view
		lvhti := me.HitTest(menuPos)         // to find item below cursor, if any

		if lvhti.IItem != -1 { // an item was right-clicked
			if !hasCtrl && !hasShift {
				clickedItem := me.Item(uint(lvhti.IItem))
				if !clickedItem.IsSelected() {
					me.SelectAllItems(false)
					clickedItem.Select(true)
				}
				clickedItem.Focus()
			}
		} else if !hasCtrl && !hasShift { // no item was right-clicked
			me.SelectAllItems(false)
		}
		me.Hwnd().SetFocus() // because a right-click won't set the focus by itself

	} else { // usually fired with the context keyboard key
		focusedItem := me.FocusedItem()
		if focusedItem != nil && focusedItem.IsVisible() { // there is a focused item, and it's visible
			rcItem := focusedItem.Rect(co.LVIR_BOUNDS)
			menuPos.X = rcItem.Left + 16 // arbitrary
			menuPos.Y = rcItem.Top + (rcItem.Bottom-rcItem.Top)/2
		} else { // no item is focused and visible
			menuPos.X = 6 // arbitrary
			menuPos.Y = 10
		}
	}

	me.contextMenu.ShowAtPoint(menuPos, me.Hwnd().GetParent(), me.Hwnd())
}

// Syntactic sugar.
func (me *ListView) sendLvmMessage(msg co.LVM,
	wParam win.WPARAM, lParam win.LPARAM) uintptr {

	return me.Hwnd().SendMessage(co.WM(msg), wParam, lParam)
}
