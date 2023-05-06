package proxy

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProxy(t *testing.T) {
	var bk, bk2 book
	bk = NewProxyBook(32)
	bk.info()
	bk2 = NewProxyBook(32)
	fmt.Printf("Compare result before bk2 dowloads a book: %d == %d ? %t\n",
		bk.getBookId(), bk2.getBookId(), cmp.Equal(bk.getBookId(), bk2.getBookId()))
	bk2.info()
	fmt.Printf("Compare result after bk2 dowloads a book: %d == %d ? %t\n",
		bk.getBookId(), bk2.getBookId(), cmp.Equal(bk.getBookId(), bk2.getBookId()))
}
