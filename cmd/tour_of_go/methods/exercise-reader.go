/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, fmt.Errorf("not a valid buffer")
	}
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}
func main() {
	reader.Validate(MyReader{})
}
