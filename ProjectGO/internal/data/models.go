package data

import (
	"database/sql"
	"errors"
)

var (
	ErrBagNotFound = errors.New("bag not found")
)

type Models struct {
	LaptopBags LaptopBagModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		LaptopBags: LaptopBagModel{DB: db},
	}
}
