package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (v int, err error) {
	for el := range b {
		b[el] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
