package abstract

import (
	"fmt"
	"testing"
)

func TestAbstract(t *testing.T) {
	bnj := NewFactory("BNJ")
	ast := NewFactory("AST")
	doc1 := ast.printDoc()
	doc2 := bnj.printPoster()
	fmt.Printf("Type %T, Publisher: %s\n", doc1, doc1.getPublisher())
	fmt.Printf("Type %T, Publisher: %s\n", doc2, doc2.getPublisher())
}
