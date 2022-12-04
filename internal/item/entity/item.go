package entity

type Item struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      int16   `json:"amount"`
	Price       float64 `json:"price"`
}
