package header

import (
	"fmt"
	"http-header/mime"
	"testing"
)

func Test_Header(t *testing.T) {

	h := NewHeaders().SetAccept(mime.JPEG, mime.BMP).Build()
	fmt.Printf("Headers: %#v\n", h)
}
