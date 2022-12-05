package entity

type OrderStatus uint8

const (
	_                          = iota
	WaitingPayment OrderStatus = iota
	Payed
	Sended
	Delivered
)

type Order struct {
	ID     uint64       `json:"id"`
	IDUser uint64       `json:"idUser"`
	Status OrderStatus  `json:"status"`
	Total  float64      `json:"total"`
	Itens  []OrderItens `json:"itens"`
}

type OrderItens struct {
	IDItem   uint64  `json:"id"`
	Quantity uint8   `json:"quantity"`
	Total    float64 `json:"total"`
}
