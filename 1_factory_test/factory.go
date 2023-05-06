package factory

import (
	"fmt"
)

func NewItem(id string) writtenDoc {
	switch id {
	case "book":
		return NewBook()
	case "magazine":
		return NewMagazine()
	default:
		fmt.Errorf("wrong type")
		return nil
	}
}
