package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func updateOrderToSended(db *sql.DB, id uint64) error {
	_, err := db.Exec("update orders set status = $2 where id = $1", id, 3)
	if err != nil {
		return err
	}
	return nil
}
