/*
Exercise: rot13Reader
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot.r.Read(bytes)
	for i := 0; i < len(bytes); i++ {
		ch := bytes[i]
		if ch >= 'A' && ch <= 'Z' {
			rot13 := ((ch - 'A' + 13) % 26) + 'A'
			bytes[i] = rot13
		} else if ch >= 'a' && ch <= 'z' {
			rot13 := ((ch - 'a' + 13) % 26) + 'a'
			bytes[i] = rot13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
