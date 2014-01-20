// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package term

import (
	"syscall"
	"unsafe"
)

var (
	kernelModule32 = syscall.NewLazyDLL("kernel32.dll")
	procScreenBuf  = kernelModule32.NewProc("GetConsoleScreenBufferInfo")
)

type coord struct {
	x uint16
	y uint16
}

type smallRect struct {
	left   uint16
	top    uint16
	right  uint16
	bottom uint16
}

type consoleScreenBuffer struct {
	size       coord
	cursorPos  coord
	attrs      int32
	window     smallRect
	maxWinSize coord
}

// NewWindow returns a pointer of Window that contains
// the current terminal window size and position.
func NewWindow() (*Window, error) {
	ws := new(Window)

	handle, err := syscall.Open("CONOUT$", syscall.O_RDONLY, 0)
	if err != nil {
		// It returns a zerod value pointer instead of an error.
		return &Window{X: 0, Y: 0, Row: 0, Col: 0}, nil
	}
	defer syscall.Close(handle)

	sb, err := screenBuffer(handle)
	if err != nil {
		return &Window{X: 0, Y: 0, Row: 0, Col: 0}, nil
	}

	ws.Row = sb.size.y
	ws.Col = sb.size.x
	ws.X = sb.window.left
	ws.Y = sb.window.top

	return ws, nil
}

func screenBuffer(handle syscall.Handle) (sb consoleScreenBuffer, err error) {
	rc, _, ec := syscall.Syscall(procScreenBuf.Addr(), 2, uintptr(handle),
		uintptr(unsafe.Pointer(&sb)), 0)
	if rc == 0 {
		err = syscall.Errno(ec)
	}
	return
}
