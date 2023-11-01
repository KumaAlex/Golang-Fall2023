package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	LaptopBags LaptopBagModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		LaptopBags: LaptopBagModel{DB: db},
	}
}
