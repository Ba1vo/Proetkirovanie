package state

import (
	"fmt"
	"time"
)

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
	draft       *draftState
	confirm     *awaitingConfirmState
	shipping    *shippingState
	state       state
	id          int
	date        string
	arrivalDate string
	adress      string
	books       []CardBook
}

type state interface {
	forward(o *order)
	changeInfo(ID int, Adress string, Books []CardBook, o *order)
	cancel(o *order)
	orderInfo(o *order)
}

func NewOrder(ID int, Adress string, Books []CardBook) order {
	o := order{}
	draft := draftState{}
	confirm := awaitingConfirmState{}
	shipping := shippingState{}
	o.draft = &draft
	o.confirm = &confirm
	o.shipping = &shipping
	o.id = ID
	o.adress = Adress
	o.books = Books
	o.state = o.draft
	return o
}

func (o *order) Forward() { o.state.forward(o) }

func (o *order) ChangeInfo(ID int, Adress string, Books []CardBook) {
	o.state.changeInfo(ID, Adress, Books, o)
}

func (o *order) Cancel() { o.state.cancel(o) }

func (o *order) OrderInfo() { o.state.orderInfo(o) }

type draftState struct {
}

func (d *draftState) forward(o *order) {
	o.date = time.Now().Format("01-02-2006")
	o.state = o.confirm
}

func (d *draftState) changeInfo(ID int, Adress string, Books []CardBook, o *order) {
	o.id = ID
	o.adress = Adress
	o.books = Books
}

func (d *draftState) cancel(o *order) { fmt.Println("Nothing to cancel") }

func (d *draftState) orderInfo(o *order) {
	fmt.Printf("Adress: %s, contents: %v\n", o.adress, o.books)
}

type awaitingConfirmState struct {
}

func (d *awaitingConfirmState) forward(o *order) {
	o.arrivalDate = time.Now().AddDate(0, 0, 7).Format("01-02-2006")
	o.state = o.shipping
}

func (d *awaitingConfirmState) changeInfo(ID int, Adress string, Books []CardBook, o *order) {
	fmt.Println("Can't change info on this stage")
}

func (d *awaitingConfirmState) cancel(o *order) { o.state = o.draft }

func (d *awaitingConfirmState) orderInfo(o *order) {
	fmt.Printf("Order created on %s, Adress: %s, contents: %v\n", o.date, o.adress, o.books)
}

type shippingState struct {
}

func (d *shippingState) forward(o *order) {
	fmt.Println("Can't move order to the next stage")
}

func (d *shippingState) changeInfo(ID int, Adress string, Books []CardBook, o *order) {
	fmt.Println("Too late lad, order is on it's way")
}

func (d *shippingState) cancel(o *order) { o.state = o.draft }

func (d *shippingState) orderInfo(o *order) {
	fmt.Printf("Order created on %s, Adress: %s, contents: %v, will arrive on %s\n",
		o.date, o.adress, o.books, o.arrivalDate)
}
