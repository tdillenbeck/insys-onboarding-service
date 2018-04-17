// +build !windows

package config

import (
	"syscall"
	"unsafe"

	"weavelab.xyz/wlib/color"
)

func terminalWidth() int {
	ws := &winsize{}
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		color.DisableColor()
	}

	w := int(ws.Col)
	if w <= 0 {
		w = 70
	} else if w > 200 {
		w = 200
	}

	return w
}
