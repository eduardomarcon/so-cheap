package entity

import "errors"

type Item struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      int16   `json:"amount"`
	Price       float64 `json:"price"`
}

func NewItem(description string, amount int16, price float64) (*Item, error) {
	item := &Item{
		Description: description,
		Amount:      amount,
		Price:       price,
	}
	if err := item.isValid(); err != nil {
		return nil, err
	}
	return item, nil
}

func (i *Item) isValid() error {
	if i.Description == "" {
		return errors.New("description invalid")
	}
	if i.Amount <= 0 {
		return errors.New("amount invalid")
	}
	if i.Price <= 0 {
		return errors.New("price invalid")
	}
	return nil
}
