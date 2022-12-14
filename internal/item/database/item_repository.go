package database

import (
	"database/sql"
	"so-cheap/internal/config"
	"so-cheap/internal/item/entity"
)

type ItemRepository interface {
	Insert(item entity.Item) (int, error)
	Update(item entity.Item) error
	Delete(id int64) error
	FindOne(id int64) (entity.Item, error)
	FindAll() ([]entity.Item, error)
}

type PGItem struct {
	db *sql.DB
}

func (p PGItem) Insert(item entity.Item) (int, error) {
	query := `insert into itens (description, amount, price) values ($1, $2, $3) RETURNING id`
	lastInsertId := 0
	err := p.db.QueryRow(query, item.Description, item.Amount, item.Price).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (p PGItem) Update(item entity.Item) error {
	query := `update itens set description = $2, amount = $3, price = $4 where id = $1`
	_, err := p.db.Exec(query, item.ID, item.Description, item.Amount, item.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p PGItem) Delete(id int64) error {
	query := `delete from itens where id = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p PGItem) FindOne(id int64) (entity.Item, error) {
	item := entity.Item{}
	query := `select * from itens where id = $1`
	row := p.db.QueryRow(query, id)
	err := row.Scan(&item.ID, &item.Description, &item.Amount, &item.Price)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (p PGItem) FindAll() ([]entity.Item, error) {
	var itens []entity.Item
	query := `select * from itens`
	row, err := p.db.Query(query)
	if err != nil {
		return itens, err
	}

	var item entity.Item
	for row.Next() {
		err = row.Scan(&item.ID, &item.Description, &item.Amount, &item.Price)
		if err != nil {
			return itens, err
		}
		itens = append(itens, item)
	}
	return itens, nil
}

func NewItemRepository() (ItemRepository, error) {
	connection, err := config.OpenConnection()
	if err != nil {
		return nil, err
	}

	repository := &PGItem{db: connection}
	return repository, nil
}
