package lightweight

import (
	"fmt"
	"testing"
)

func TestLight(t *testing.T) {
	light := NewLight()
	book1 := light.GetBook(32)
	book2 := light.GetBook(33)
	book2 = light.GetBook(33)
	book1.SetPostion("main")
	fmt.Printf("Value %v\n", book1)
	book1.Display()
	fmt.Printf("Value %v\n", book2)
	book2.Display()
}
