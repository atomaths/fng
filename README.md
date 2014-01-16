## Simple find and grep command line tool

The reuslt of fng is easy to see output.

This repository also includes [`term`](http://godoc.org/github.com/atomaths/fng/term)
package that can get the size of the current terminal window.

### Installation

```bash
$ go get github.com/atomaths/fng
```

### Usage

```bash
$ fng *.go string_to_find
$ fng *.go *.c "string to find"
```

### API reference
[godoc.org/github.com/atomaths/fng](http://godoc.org/github.com/atomaths/fng)
