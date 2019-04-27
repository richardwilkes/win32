package win32

//nolint:golint
// https://docs.microsoft.com/en-us/windows/desktop/WinProg/windows-data-types
type (
	ATOM                  uint16
	BOOL                  int16
	DWORD                 uint32
	HACCEL                uintptr
	HANDLE                uintptr
	HBITMAP               uintptr
	HCURSOR               uintptr
	HDC                   uintptr
	HMENU                 uintptr
	HINSTANCE             uintptr
	HMONITOR              uintptr
	HWND                  uintptr
	LPACCEL               uintptr
	LPCWSTR               *uint16
	LPVOID                uintptr
	LRESULT               uintptr
	LPARAM                uintptr
	WORD                  uint16
	WPARAM                uintptr
	DPI_AWARENESS_CONTEXT uint32
)
