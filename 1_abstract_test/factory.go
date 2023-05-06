package abstract

type publisher interface {
	printDoc() writtenDoc
	printPoster() poster
}

func NewFactory(name string) publisher {
	if name == "AST" {
		ast := NewAST()
		return &ast
	}
	if name == "BNJ" {
		bnj := NewBNJ()
		return &bnj
	}
	return nil
}

type writtenDoc interface {
	getAuthor() string
	getName() string
	getPrice() string
	getPhoto() string
	getPublisher() string
	setAuthor(string)
	setName(string)
	setPrice(string)
	setDiscount(int)
	setPhoto(string)
}

type poster interface {
	getSize() string
	getPrice() string
	getPhoto() string
	getPublisher() string
	setPrice(string)
	setDiscount(int)
	setPhoto(string)
}
