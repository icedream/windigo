package proc

import (
	"syscall"
)

var (
	uxtheme = syscall.NewLazyDLL("uxtheme.dll")

	CloseThemeData                        = uxtheme.NewProc("CloseThemeData")
	DrawThemeBackground                   = uxtheme.NewProc("DrawThemeBackground")
	GetThemePosition                      = uxtheme.NewProc("GetThemePosition")
	GetThemeSysColorBrush                 = uxtheme.NewProc("GetThemeSysColorBrush")
	GetThemeSysFont                       = uxtheme.NewProc("GetThemeSysFont")
	GetThemeTextMetrics                   = uxtheme.NewProc("GetThemeTextMetrics")
	IsAppThemed                           = uxtheme.NewProc("IsAppThemed")
	IsCompositionActive                   = uxtheme.NewProc("IsCompositionActive")
	IsThemeActive                         = uxtheme.NewProc("IsThemeActive")
	IsThemeBackgroundPartiallyTransparent = uxtheme.NewProc("IsThemeBackgroundPartiallyTransparent")
	IsThemeDialogTextureEnabled           = uxtheme.NewProc("IsThemeDialogTextureEnabled")
	IsThemePartDefined                    = uxtheme.NewProc("IsThemePartDefined")
	OpenThemeData                         = uxtheme.NewProc("OpenThemeData")
)
