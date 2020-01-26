package win32

// Simple types https://docs.microsoft.com/en-us/windows/desktop/WinProg/windows-data-types
type (
	ATOM                  uint16
	BOOL                  int16
	COLORREF              uint32
	DPI_AWARENESS_CONTEXT uint32
	DWORD                 uint32
	HACCEL                uintptr
	HANDLE                uintptr
	HBITMAP               uintptr
	HBRUSH                uintptr
	HCURSOR               uintptr
	HDC                   uintptr
	HFONT                 uintptr
	HGDIOBJ               uintptr
	HGLOBAL               uintptr
	HINSTANCE             uintptr
	HMENU                 uintptr
	HPEN                  uintptr
	HMONITOR              uintptr
	HWND                  uintptr
	LPACCEL               uintptr
	LPCWSTR               *uint16
	LPVOID                uintptr
	LRESULT               uintptr
	LPARAM                uintptr
	WORD                  uint16
	WPARAM                uintptr
)
