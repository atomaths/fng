// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package term provides current terminal window informations.
package term

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	_TIOCGWINSZ = 0x5413 // On OSX use 1074295912
)

type Window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

var win *Window = nil

func init() {
	ws, err := NewWindow()
	if err != nil {
		panic(err)
	}
	win = ws
}

func NewWindow() (*Window, error) {
	ws := new(Window)

	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(_TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if errno != 0 || int(r1) == -1 {
		return nil, os.NewSyscallError("NewWindow", errno)
	}
	return ws, nil
}

// Width returns the window width size.
func (w *Window) Width() uint16 {
	return w.Col
}

// Width returns the window width size for the standard window.
func Width() uint16 {
	return win.Width()
}
