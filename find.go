// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import "fmt"

type finder struct {
	files []string
}

func (f *finder) String() string {
	s := []byte("files: ")
	for _, v := range f.files {
		s = append(s, fmt.Sprintf("%s, ", v)...)
	}
	s = s[:len(s)-2]
	return string(s)
}
