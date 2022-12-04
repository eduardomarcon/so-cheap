package database

import (
	"database/sql"
	"so-cheap/internal/config"
	"so-cheap/internal/user/entity"
)

type UserRepository interface {
	Insert(user entity.User) (int, error)
	Update(user entity.User) error
	Delete(id int64) error
	FindOne(id int64) (entity.User, error)
	FindAll() ([]entity.User, error)
}

type PGUser struct {
	db *sql.DB
}

func (p PGUser) Insert(user entity.User) (int, error) {
	query := `insert into users (name, email) values ($1, $2) RETURNING id`
	lastInsertId := 0
	err := p.db.QueryRow(query, user.Name, user.Email).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (p PGUser) Update(user entity.User) error {
	query := `update users set name = $2, email = $3 where id = $1`
	_, err := p.db.Exec(query, user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p PGUser) Delete(id int64) error {
	query := `delete from users where id = $1`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p PGUser) FindOne(id int64) (entity.User, error) {
	user := entity.User{}
	query := `select * from users where id = $1`
	row := p.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p PGUser) FindAll() ([]entity.User, error) {
	var users []entity.User
	query := `select * from users`
	row, err := p.db.Query(query)
	if err != nil {
		return users, err
	}

	var user entity.User
	for row.Next() {
		err = row.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func NewUserRepository() (UserRepository, error) {
	connection, err := config.OpenConnection()
	if err != nil {
		return nil, err
	}

	repository := &PGUser{db: connection}
	return repository, nil
}
