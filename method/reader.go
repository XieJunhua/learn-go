package main

import (
	"fmt"
	"io"
	"strings"
	"golang.org/x/tour/reader"
)

type MyReader struct{
}

func (m MyReader) Read(b []byte) (i int, e error) {
    // for x := range b {
    //     b[x] = 'A'
    // }
    // return len(b), nil

	b[0] = byte('A')
    return 1, nil
}

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}	
	}

	reader.Validate(MyReader{})
}

