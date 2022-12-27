package store

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	Id     string
	Title  string
	Amount float64
	Date   time.Time
}

var store map[uuid.UUID]Payload

func NewStore() {
	id1, id2, id3, id4 := uuid.New(), uuid.New(), uuid.New(), uuid.New()
	store = map[uuid.UUID]Payload{
		id1: {
			Id:     id1.String(),
			Title:  "Toilet Paper",
			Amount: 94.12,
			Date:   time.Date(2020, 7, 14, 0, 0, 0, 0, time.Local),
		},
		id2: {
			Id:     id2.String(),
			Title:  "New TV",
			Amount: 799.49,
			Date:   time.Date(2021, 2, 12, 0, 0, 0, 0, time.Local),
		},
		id3: {
			Id:     id3.String(),
			Title:  "Car Insurance",
			Amount: 294.67,
			Date:   time.Date(2021, 2, 28, 0, 0, 0, 0, time.Local),
		},
		id4: {
			Id:     id4.String(),
			Title:  "New Desk (Wooden)",
			Amount: 450,
			Date:   time.Date(2021, 5, 12, 0, 0, 0, 0, time.Local),
		},
	}
}

func storeInMemory(id uuid.UUID, input Payload) {
	store[id] = input
}

func SaveExpense(expense Payload) {
	newId := uuid.New()
	storeInMemory(newId, Payload{
		Id:     newId.String(),
		Title:  expense.Title,
		Amount: expense.Amount,
		Date:   expense.Date,
	})
}

func FetchExpenses() []Payload {
	expenses := []Payload{}
	for _, payload := range store {
		expenses = append(expenses, payload)
	}
	return expenses
}
