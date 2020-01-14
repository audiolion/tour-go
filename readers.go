/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(bytes []byte) (int, error) {
	length := len(bytes)
	for i := 0; i < length; i++ {
		bytes[i] = 'A'
	}
	return length, nil
}

func main() {
	reader.Validate(MyReader{})
}
