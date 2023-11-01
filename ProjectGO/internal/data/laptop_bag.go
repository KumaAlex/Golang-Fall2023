package data

import (
	"ProjectGO/internal/validator"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
)

type LaptopBag struct {
	ID           int64     `json:"id,omitempty"`
	CreatedAt    time.Time `json:"-"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	Color        string    `json:"color"`
	Material     string    `json:"material,omitempty"`
	Compartments int32     `json:"compartments,omitempty"`
	Weight       Weight    `json:"weight,omitempty"`
	Dimensions   []float32 `json:"dimensions,omitempty"`
}

func ValidateLaptopBag(v *validator.Validator, laptopBag *LaptopBag) {
	v.Check(laptopBag.Brand != "", "brand", "must be provided")
	v.Check(len(laptopBag.Brand) <= 500, "brand", "must not be more than 500 bytes long")
	v.Check(laptopBag.Color != "", "color", "must be provided")
	v.Check(len(laptopBag.Color) <= 64, "color", "must not be more than 64 bytes long")
	v.Check(laptopBag.Weight != 0, "weight", "must be provided")
	v.Check(laptopBag.Weight >= 0, "weight", "must be greater than 0")
	v.Check(laptopBag.Dimensions != nil, "dimensions", "must be provided")
	v.Check(len(laptopBag.Dimensions) == 3, "genres", "must contain exactly 3 dimensions")
}

type LaptopBagModel struct {
	DB *sql.DB
}

func (l LaptopBagModel) Insert(laptopBag *LaptopBag) error {
	query := `
		INSERT INTO laptopBags (brand, color, weight, dimensions)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []interface{}{laptopBag.Brand, laptopBag.Color, laptopBag.Weight, pq.Array(laptopBag.Dimensions)}

	return l.DB.QueryRow(query, args...).Scan(&laptopBag.ID, &laptopBag.CreatedAt)
}

func (l LaptopBagModel) Get(id int64) (*LaptopBag, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, Created_at, Brand, Model, Color, Material, Compartments, Weight, Dimensions
		FROM laptopbags
		WHERE id = $1`

	var laptopbag LaptopBag

	err := l.DB.QueryRow(query, id).Scan(
		&laptopbag.ID,
		&laptopbag.CreatedAt,
		&laptopbag.Brand,
		&laptopbag.Model,
		&laptopbag.Color,
		&laptopbag.Material,
		&laptopbag.Compartments,
		&laptopbag.Weight,
		pq.Array(&laptopbag.Dimensions),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &laptopbag, nil
}

func (l LaptopBagModel) Update(laptopBag *LaptopBag) error {
	query := `
		UPDATE laptopBags
		SET brand = $1, color = $2, weight = $3, dimensions = $4
		WHERE id = $5
		RETURNING ID`

	args := []interface{}{
		laptopBag.Brand,
		laptopBag.Color,
		laptopBag.Weight,
		pq.Array(laptopBag.Dimensions),
		laptopBag.ID,
	}
	// Use the QueryRow() method to execute the query, passing in the args slice as a
	// variadic parameter and scanning the new version value into the movie struct.
	return l.DB.QueryRow(query, args...).Scan(&laptopBag.ID)

}

func (l LaptopBagModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM laptopbags
		WHERE id = $1`

	result, err := l.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
