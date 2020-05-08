package consts

type GA uint32 // GetAncestor

const (
	GA_PARENT    GA = 1
	GA_ROOT      GA = 2
	GA_ROOTOWNER GA = 3
)

type GWLP int32 // GetWindowLongPtr offsets

const (
	GWLP_STYLE      GWLP = -16
	GWLP_EXSTYLE    GWLP = -20
	GWLP_WNDPROC    GWLP = -4
	GWLP_HINSTANCE  GWLP = -6
	GWLP_HWNDPARENT GWLP = -8
	GWLP_USERDATA   GWLP = -21
	GWLP_ID         GWLP = -12
)

type HDM WM // list view header message

const (
	hDM_FIRST HDM = 0x1200

	HDM_GETITEMCOUNT HDM = hDM_FIRST + 0
	HDM_INSERTITEM   HDM = hDM_FIRST + 10
	HDM_DELETEITEM   HDM = hDM_FIRST + 2
	HDM_GETITEM      HDM = hDM_FIRST + 11
	HDM_SETITEM      HDM = hDM_FIRST + 12
	HDM_LAYOUT       HDM = hDM_FIRST + 5
)

type ID uint16 // dialog box command ID

const (
	IDOK       ID = 1
	IDCANCEL   ID = 2
	IDABORT    ID = 3
	IDRETRY    ID = 4
	IDIGNORE   ID = 5
	IDYES      ID = 6
	IDNO       ID = 7
	IDCLOSE    ID = 8
	IDHELP     ID = 9
	IDTRYAGAIN ID = 10
	IDCONTINUE ID = 11
	IDTIMEOUT  ID = 32000
)

type IDC uintptr // LoadCursor

const (
	IDC_ARROW       IDC = 32512
	IDC_IBEAM       IDC = 32513
	IDC_WAIT        IDC = 32514
	IDC_CROSS       IDC = 32515
	IDC_UPARROW     IDC = 32516
	IDC_SIZENWSE    IDC = 32642
	IDC_SIZENESW    IDC = 32643
	IDC_SIZEWE      IDC = 32644
	IDC_SIZENS      IDC = 32645
	IDC_SIZEALL     IDC = 32646
	IDC_NO          IDC = 32648
	IDC_HAND        IDC = 32649
	IDC_APPSTARTING IDC = 32650
	IDC_HELP        IDC = 32651
	IDC_PIN         IDC = 32671
	IDC_PERSON      IDC = 32672
)