package internal

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SignupRepository(data *GetSubRequest) error {
	query := fmt.Sprintf("INSERT INTO users(name, email, password) VALUES('%s','%s','%s')",
		data.Name,
		data.Email,
		data.Password,
	)

	_, err := r.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) InsertRepository() error {
	query := "INSERT INTO users(name, email, password) VALUES('node','node@gmail.com','12345')"
	_, err := r.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
