package proc

import (
	"syscall"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	AdjustWindowRectEx            = user32.NewProc("AdjustWindowRectEx")
	AllowSetForegroundWindow      = user32.NewProc("AllowSetForegroundWindow")
	AppendMenu                    = user32.NewProc("AppendMenuW")
	BeginDeferWindowPos           = user32.NewProc("BeginDeferWindowPos")
	BeginPaint                    = user32.NewProc("BeginPaint")
	BroadcastSystemMessage        = user32.NewProc("BroadcastSystemMessageW")
	CallNextHookEx                = user32.NewProc("CallNextHookEx")
	CheckMenuItem                 = user32.NewProc("CheckMenuItem")
	CheckMenuRadioItem            = user32.NewProc("CheckMenuRadioItem")
	ChildWindowFromPoint          = user32.NewProc("ChildWindowFromPoint")
	ChildWindowFromPointEx        = user32.NewProc("ChildWindowFromPointEx")
	ClientToScreen                = user32.NewProc("ClientToScreen")
	CloseClipboard                = user32.NewProc("CloseClipboard")
	CopyAcceleratorTable          = user32.NewProc("CopyAcceleratorTableW")
	CopyIcon                      = user32.NewProc("CopyIcon")
	CountClipboardFormats         = user32.NewProc("CountClipboardFormats")
	CreateAcceleratorTable        = user32.NewProc("CreateAcceleratorTableW")
	CreateDialogParam             = user32.NewProc("CreateDialogParamW")
	CreateIconIndirect            = user32.NewProc("CreateIconIndirect")
	CreateMenu                    = user32.NewProc("CreateMenu")
	CreatePopupMenu               = user32.NewProc("CreatePopupMenu")
	CreateWindowEx                = user32.NewProc("CreateWindowExW")
	DefDlgProc                    = user32.NewProc("DefDlgProcW")
	DeferWindowPos                = user32.NewProc("DeferWindowPos")
	DefWindowProc                 = user32.NewProc("DefWindowProcW")
	DeleteMenu                    = user32.NewProc("DeleteMenu")
	DestroyAcceleratorTable       = user32.NewProc("DestroyAcceleratorTable")
	DestroyCaret                  = user32.NewProc("DestroyCaret")
	DestroyCursor                 = user32.NewProc("DestroyCursor")
	DestroyIcon                   = user32.NewProc("DestroyIcon")
	DestroyMenu                   = user32.NewProc("DestroyMenu")
	DestroyWindow                 = user32.NewProc("DestroyWindow")
	DialogBoxParam                = user32.NewProc("DialogBoxParamW")
	DispatchMessage               = user32.NewProc("DispatchMessageW")
	DrawIcon                      = user32.NewProc("DrawIcon")
	DrawIconEx                    = user32.NewProc("DrawIconEx")
	DrawMenuBar                   = user32.NewProc("DrawMenuBar")
	EmptyClipboard                = user32.NewProc("EmptyClipboard")
	EnableMenuItem                = user32.NewProc("EnableMenuItem")
	EnableWindow                  = user32.NewProc("EnableWindow")
	EndDeferWindowPos             = user32.NewProc("EndDeferWindowPos")
	EndDialog                     = user32.NewProc("EndDialog")
	EndMenu                       = user32.NewProc("EndMenu")
	EndPaint                      = user32.NewProc("EndPaint")
	EnumChildWindows              = user32.NewProc("EnumChildWindows")
	EnumClipboardFormats          = user32.NewProc("EnumClipboardFormats")
	EnumDisplayDevices            = user32.NewProc("EnumDisplayDevicesW")
	EnumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	EnumWindows                   = user32.NewProc("EnumWindows")
	FindWindow                    = user32.NewProc("FindWindowW")
	GetAncestor                   = user32.NewProc("GetAncestor")
	GetAsyncKeyState              = user32.NewProc("GetAsyncKeyState")
	GetCaretPos                   = user32.NewProc("GetCaretPos")
	GetClassInfoEx                = user32.NewProc("GetClassInfoExW")
	GetClassLongPtr               = user32.NewProc("GetClassLongPtrW")
	GetClassName                  = user32.NewProc("GetClassNameW")
	GetClientRect                 = user32.NewProc("GetClientRect")
	GetClipboardOwner             = user32.NewProc("GetClipboardOwner")
	GetClipboardSequenceNumber    = user32.NewProc("GetClipboardSequenceNumber")
	GetCursorPos                  = user32.NewProc("GetCursorPos")
	GetDC                         = user32.NewProc("GetDC")
	GetDesktopWindow              = user32.NewProc("GetDesktopWindow")
	GetDialogBaseUnits            = user32.NewProc("GetDialogBaseUnits")
	GetDlgCtrlID                  = user32.NewProc("GetDlgCtrlID")
	GetDlgItem                    = user32.NewProc("GetDlgItem")
	GetFocus                      = user32.NewProc("GetFocus")
	GetForegroundWindow           = user32.NewProc("GetForegroundWindow")
	GetGUIThreadInfo              = user32.NewProc("GetGUIThreadInfo")
	GetIconInfo                   = user32.NewProc("GetIconInfo")
	GetIconInfoEx                 = user32.NewProc("GetIconInfoExW")
	GetInputState                 = user32.NewProc("GetInputState")
	GetLastActivePopup            = user32.NewProc("GetLastActivePopup")
	GetMenu                       = user32.NewProc("GetMenu")
	GetMenuDefaultItem            = user32.NewProc("GetMenuDefaultItem")
	GetMenuItemCount              = user32.NewProc("GetMenuItemCount")
	GetMenuItemID                 = user32.NewProc("GetMenuItemID")
	GetMenuItemInfo               = user32.NewProc("GetMenuItemInfoW")
	GetMessage                    = user32.NewProc("GetMessageW")
	GetMessageExtraInfo           = user32.NewProc("GetMessageExtraInfo")
	GetMessagePos                 = user32.NewProc("GetMessagePos")
	GetMessageTime                = user32.NewProc("GetMessageTime")
	GetMonitorInfo                = user32.NewProc("GetMonitorInfoW")
	GetNextDlgGroupItem           = user32.NewProc("GetNextDlgGroupItem")
	GetNextDlgTabItem             = user32.NewProc("GetNextDlgTabItem")
	GetOpenClipboardWindow        = user32.NewProc("GetOpenClipboardWindow")
	GetParent                     = user32.NewProc("GetParent")
	GetPhysicalCursorPos          = user32.NewProc("GetPhysicalCursorPos")
	GetProcessDefaultLayout       = user32.NewProc("GetProcessDefaultLayout")
	GetQueueStatus                = user32.NewProc("GetQueueStatus")
	GetScrollInfo                 = user32.NewProc("GetScrollInfo")
	GetShellWindow                = user32.NewProc("GetShellWindow")
	GetSubMenu                    = user32.NewProc("GetSubMenu")
	GetSysColor                   = user32.NewProc("GetSysColor")
	GetSysColorBrush              = user32.NewProc("GetSysColorBrush")
	GetSystemMenu                 = user32.NewProc("GetSystemMenu")
	GetSystemMetrics              = user32.NewProc("GetSystemMetrics")
	GetTopWindow                  = user32.NewProc("GetTopWindow")
	GetWindow                     = user32.NewProc("GetWindow")
	GetWindowDC                   = user32.NewProc("GetWindowDC")
	GetWindowRect                 = user32.NewProc("GetWindowRect")
	GetWindowText                 = user32.NewProc("GetWindowTextW")
	GetWindowTextLength           = user32.NewProc("GetWindowTextLengthW")
	GetWindowThreadProcessId      = user32.NewProc("GetWindowThreadProcessId")
	HideCaret                     = user32.NewProc("HideCaret")
	HiliteMenuItem                = user32.NewProc("HiliteMenuItem")
	InSendMessage                 = user32.NewProc("InSendMessage")
	InSendMessageEx               = user32.NewProc("InSendMessageEx")
	InsertMenuItem                = user32.NewProc("InsertMenuItemW")
	InvalidateRect                = user32.NewProc("InvalidateRect")
	IsChild                       = user32.NewProc("IsChild")
	IsClipboardFormatAvailable    = user32.NewProc("IsClipboardFormatAvailable")
	IsDialogMessage               = user32.NewProc("IsDialogMessageW")
	IsDlgButtonChecked            = user32.NewProc("IsDlgButtonChecked")
	IsGUIThread                   = user32.NewProc("IsGUIThread")
	IsIconic                      = user32.NewProc("IsIconic")
	IsWindow                      = user32.NewProc("IsWindow")
	IsWindowEnabled               = user32.NewProc("IsWindowEnabled")
	IsWindowVisible               = user32.NewProc("IsWindowVisible")
	IsZoomed                      = user32.NewProc("IsZoomed")
	KillTimer                     = user32.NewProc("KillTimer")
	LoadAccelerators              = user32.NewProc("LoadAcceleratorsW")
	LoadCursor                    = user32.NewProc("LoadCursorW")
	LoadIcon                      = user32.NewProc("LoadIconW")
	LoadImage                     = user32.NewProc("LoadImageW")
	LoadMenu                      = user32.NewProc("LoadMenuW")
	LockSetForegroundWindow       = user32.NewProc("LockSetForegroundWindow")
	LogicalToPhysicalPoint        = user32.NewProc("LogicalToPhysicalPoint")
	MapDialogRect                 = user32.NewProc("MapDialogRect")
	MenuItemFromPoint             = user32.NewProc("MenuItemFromPoint")
	MessageBox                    = user32.NewProc("MessageBoxW")
	MonitorFromPoint              = user32.NewProc("MonitorFromPoint")
	MonitorFromRect               = user32.NewProc("MonitorFromRect")
	MonitorFromWindow             = user32.NewProc("MonitorFromWindow")
	MoveWindow                    = user32.NewProc("MoveWindow")
	OpenClipboard                 = user32.NewProc("OpenClipboard")
	PeekMessage                   = user32.NewProc("PeekMessageW")
	PhysicalToLogicalPoint        = user32.NewProc("PhysicalToLogicalPoint")
	PostMessage                   = user32.NewProc("PostMessageW")
	PostQuitMessage               = user32.NewProc("PostQuitMessage")
	PostThreadMessage             = user32.NewProc("PostThreadMessageW")
	RealChildWindowFromPoint      = user32.NewProc("RealChildWindowFromPoint")
	RealGetWindowClass            = user32.NewProc("RealGetWindowClassW")
	RegisterClassEx               = user32.NewProc("RegisterClassExW")
	RegisterWindowMessage         = user32.NewProc("RegisterWindowMessageW")
	ReleaseDC                     = user32.NewProc("ReleaseDC")
	RemoveMenu                    = user32.NewProc("RemoveMenu")
	ReplyMessage                  = user32.NewProc("ReplyMessage")
	ScreenToClient                = user32.NewProc("ScreenToClient")
	SendMessage                   = user32.NewProc("SendMessageW")
	SendMessageTimeout            = user32.NewProc("SendMessageTimeoutW")
	SetClipboardData              = user32.NewProc("SetClipboardData")
	SetFocus                      = user32.NewProc("SetFocus")
	SetForegroundWindow           = user32.NewProc("SetForegroundWindow")
	SetMenu                       = user32.NewProc("SetMenu")
	SetMenuDefaultItem            = user32.NewProc("SetMenuDefaultItem")
	SetMenuInfo                   = user32.NewProc("SetMenuInfo")
	SetMenuItemBitmaps            = user32.NewProc("SetMenuItemBitmaps")
	SetMenuItemInfo               = user32.NewProc("SetMenuItemInfoW")
	SetMessageExtraInfo           = user32.NewProc("SetMessageExtraInfo")
	SetParent                     = user32.NewProc("SetParent")
	SetProcessDefaultLayout       = user32.NewProc("SetProcessDefaultLayout")
	SetProcessDPIAware            = user32.NewProc("SetProcessDPIAware")
	SetProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")
	SetScrollInfo                 = user32.NewProc("SetScrollInfo")
	SetSystemCursor               = user32.NewProc("SetSystemCursor")
	SetTimer                      = user32.NewProc("SetTimer")
	SetUserObjectInformation      = user32.NewProc("SetUserObjectInformationW")
	SetWindowDisplayAffinity      = user32.NewProc("SetWindowDisplayAffinity")
	SetWindowPos                  = user32.NewProc("SetWindowPos")
	SetWindowsHookEx              = user32.NewProc("SetWindowsHookExW")
	SetWindowText                 = user32.NewProc("SetWindowTextW")
	ShowCaret                     = user32.NewProc("ShowCaret")
	ShowWindow                    = user32.NewProc("ShowWindow")
	SystemParametersInfo          = user32.NewProc("SystemParametersInfoW")
	TrackPopupMenu                = user32.NewProc("TrackPopupMenu")
	TranslateAccelerator          = user32.NewProc("TranslateAcceleratorW")
	TranslateMessage              = user32.NewProc("TranslateMessage")
	UnhookWindowsHookEx           = user32.NewProc("UnhookWindowsHookEx")
	UnregisterClass               = user32.NewProc("UnregisterClassW")
	UpdateWindow                  = user32.NewProc("UpdateWindow")
)
