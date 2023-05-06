package test

type CardBook struct {
	ID       int
	Name     string
	Authors  []string
	Price    string
	Discount int
	Photo    string
	Amount   int
}

type order struct {
	med    *mediator
	id     int
	date   string
	adress string
	status string
	books  []CardBook
}

func NewOrd(med *mediator, id int, date string, adress string, status string) order {
	o := order{med, id, date, adress, status, []CardBook{}}
	return o
}

func (o *order) GetBooks() []CardBook {
	return o.books
}

func (o *order) SetBooks(books []CardBook) {
	o.books = books
}
