package entity

import "errors"

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(description string, email string) (*User, error) {
	item := &User{
		Name:  description,
		Email: email,
	}
	if err := item.isValid(); err != nil {
		return nil, err
	}
	return item, nil
}

func (i *User) isValid() error {
	if i.Name == "" {
		return errors.New("description invalid")
	}
	if i.Email == "" {
		return errors.New("amount invalid")
	}
	return nil
}
