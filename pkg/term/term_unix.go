// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd

package term

import (
	"syscall"
	"unsafe"
)

const (
	_TIOCGWINSZ = 0x5413 // On OSX use 1074295912
)

// NewWindow returns a pointer of Window that contains
// the current terminal window size and position.
func NewWindow() (*Window, error) {
	ws := new(Window)

	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(_TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if errno != 0 || int(r1) == -1 {
		// It returns a zerod value pointer instead of an error.
		// return nil, os.NewSyscallError("NewWindow", errno)
		return &Window{X: 0, Y: 0, Row: 0, Col: 0}, nil
	}
	return ws, nil
}
