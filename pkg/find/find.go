// Copyright 2014 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(atomaths): Test whether the file is a binary or not.
//   http://stackoverflow.com/a/567938
//   https://en.wikipedia.org/wiki/Magic_number_(programming)
//   https://en.wikipedia.org/wiki/File_format#Magic_number
package find

type FindSet struct {
	files []string
}

func New(path string) *FindSet {
	f := &FindSet{}
	return f
}
