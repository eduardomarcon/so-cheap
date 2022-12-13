package database

import (
	"database/sql"
	"fmt"
	"so-cheap/internal/config"
	"so-cheap/internal/order/entity"
)

type OrderRepository interface {
	Insert(order entity.Order) (int, error)
	Update(order entity.Order) error
	UpdateStatus(id uint64, status entity.OrderStatus) error
	Delete(id uint64) error
	FindOne(id uint64) (entity.Order, error)
	FindAll() ([]entity.Order, error)
	FindAllByStatus(status entity.OrderStatus) ([]entity.Order, error)
}

type PGOrder struct {
	db *sql.DB
}

func (p PGOrder) Insert(order entity.Order) (int, error) {
	query := `insert into orders (id_user, status, total) values ($1, $2, $3) RETURNING id`
	lastInsertId := 0
	var total float64 = 0
	for _, item := range order.Itens {
		total += item.Total
	}
	err := p.db.QueryRow(query, order.IDUser, entity.WaitingPayment, total).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	query = `insert into orders_itens (id_order, id_item, quantity, total) values `
	numOfFields := 4
	params := make([]interface{}, len(order.Itens)*numOfFields)
	for i, item := range order.Itens {
		pos := i * numOfFields
		paramPlus0 := pos + 0
		params[paramPlus0] = lastInsertId
		paramPlus0++
		paramPlus1 := pos + 1
		params[paramPlus1] = item.IDItem
		paramPlus1++
		paramPlus2 := pos + 2
		params[paramPlus2] = item.Quantity
		paramPlus2++
		paramPlus3 := pos + 3
		params[paramPlus3] = item.Total
		paramPlus3++
		query += fmt.Sprintf(`($%d,$%d,$%d,$%d),`, paramPlus0, paramPlus1, paramPlus2, paramPlus3)
	}

	query = query[:len(query)-1]
	_, err = p.db.Exec(query, params...)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (p PGOrder) Update(order entity.Order) error {
	query := `update orders set id_user = $2, status = $3, total = $4 where id = $1`
	_, err := p.db.Exec(query, order.ID, order.IDUser, order.Status, order.Total)
	if err != nil {
		return err
	}
	return nil
}

func (p PGOrder) UpdateStatus(id uint64, status entity.OrderStatus) error {
	query := `update orders set status = $2 where id = $1`
	_, err := p.db.Exec(query, id, status)
	if err != nil {
		return err
	}
	return nil
}

func (p PGOrder) Delete(id uint64) error {
	query := `delete from orders_itens where id_order = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	query = `delete from orders where id = $1`
	_, err = p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p PGOrder) FindOne(id uint64) (entity.Order, error) {
	order := entity.Order{}
	query := `select * from orders where id = $1`
	row := p.db.QueryRow(query, id)
	err := row.Scan(&order.ID, &order.IDUser, &order.Status, &order.Total)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (p PGOrder) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	query := `select * from orders`
	row, err := p.db.Query(query)
	if err != nil {
		return orders, err
	}

	var order entity.Order
	for row.Next() {
		err = row.Scan(&order.ID, &order.IDUser, &order.Status, &order.Total)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (p PGOrder) FindAllByStatus(status entity.OrderStatus) ([]entity.Order, error) {
	var orders []entity.Order
	query := `select * from orders where status = $1`
	row, err := p.db.Query(query, status)
	if err != nil {
		return orders, err
	}

	var order entity.Order
	for row.Next() {
		err = row.Scan(&order.ID, &order.IDUser, &order.Status, &order.Total)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func NewOrderRepository() (OrderRepository, error) {
	connection, err := config.OpenConnection()
	if err != nil {
		return nil, err
	}

	repository := &PGOrder{db: connection}
	return repository, nil
}
