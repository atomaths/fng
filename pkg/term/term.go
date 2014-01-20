// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package term provides the size and position of the current terminal window.
// This package supports Linux, Mac and Windows.
package term

// A Window contains window's row(height), col(width) pixel size and x, y point.
// Note that field declaration order is important.(Row, Col following by X, Y)
// The reason for this is from syscall.Syscall in the NewWindow function.
type Window struct {
	Row, Col uint16
	X, Y     uint16
}

var win *Window = nil

func init() {
	ws, err := NewWindow()
	if err != nil {
		panic(err)
	}
	win = ws
}

// Width returns the window column size.
func (w *Window) Width() uint16 {
	return w.Col
}

// Width returns the window column size for the standard window.
func Width() uint16 {
	return win.Width()
}
