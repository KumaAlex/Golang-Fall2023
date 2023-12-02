package data

import (
	"database/sql"
	"errors"
)

var (
	ErrBagNotFound  = errors.New("bag not found")
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	LaptopBags LaptopBagModel
	Tokens     TokenModel
	Users      UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		LaptopBags: LaptopBagModel{DB: db},
		Tokens:     TokenModel{DB: db},
		Users:      UserModel{DB: db},
	}
}
