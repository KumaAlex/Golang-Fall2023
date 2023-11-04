package data

import (
	"ProjectGO/internal/validator"
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	Version      int32     `json:"version"`
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
		RETURNING id, created_at, version`

	args := []interface{}{laptopBag.Brand, laptopBag.Color, laptopBag.Weight, pq.Array(laptopBag.Dimensions)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return l.DB.QueryRowContext(ctx, query, args...).Scan(&laptopBag.ID, &laptopBag.CreatedAt, &laptopBag.Version)
}

func (l LaptopBagModel) Get(id int64) (*LaptopBag, error) {
	if id < 1 {
		return nil, ErrBagNotFound
	}

	query := `
		SELECT id, Created_at, Brand, Model, Color, Material, Compartments, Weight, Dimensions, Version
		FROM laptopbags
		WHERE id = $1`

	var laptopbag LaptopBag

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := l.DB.QueryRowContext(ctx, query, id).Scan(
		&laptopbag.ID,
		&laptopbag.CreatedAt,
		&laptopbag.Brand,
		&laptopbag.Model,
		&laptopbag.Color,
		&laptopbag.Material,
		&laptopbag.Compartments,
		&laptopbag.Weight,
		pq.Array(&laptopbag.Dimensions),
		&laptopbag.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrBagNotFound
		default:
			return nil, err
		}
	}

	return &laptopbag, nil
}

func (l LaptopBagModel) Update(laptopBag *LaptopBag) error {
	query := `
		UPDATE laptopBags
		SET brand = $1, color = $2, weight = $3, dimensions = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version`

	args := []interface{}{
		laptopBag.Brand,
		laptopBag.Color,
		laptopBag.Weight,
		pq.Array(laptopBag.Dimensions),
		laptopBag.ID,
		laptopBag.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := l.DB.QueryRowContext(ctx, query, args...).Scan(&laptopBag.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (l LaptopBagModel) Delete(id int64) error {
	if id < 1 {
		return ErrBagNotFound
	}

	query := `
		DELETE FROM laptopbags
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := l.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrBagNotFound
	}
	return nil
}

func (l LaptopBagModel) GetAll(brand string, color string, filters Filters) ([]*LaptopBag, Metadata, error) {

	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, brand, model, color, material, compartments, weight, dimensions, version
		FROM laptopbags
		WHERE (to_tsvector('simple', brand) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (to_tsvector('simple', color) @@ plainto_tsquery('simple', $2) OR $2 = '')
		ORDER BY %s %s, id ASC 
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{brand, color, filters.limit(), filters.offset()}

	rows, err := l.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalBags := 0
	laptopBags := []*LaptopBag{}

	for rows.Next() {

		var laptopbag LaptopBag

		err := rows.Scan(
			&totalBags,
			&laptopbag.ID,
			&laptopbag.CreatedAt,
			&laptopbag.Brand,
			&laptopbag.Model,
			&laptopbag.Color,
			&laptopbag.Material,
			&laptopbag.Compartments,
			&laptopbag.Weight,
			pq.Array(&laptopbag.Dimensions),
			&laptopbag.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		laptopBags = append(laptopBags, &laptopbag)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalBags, filters.Page, filters.PageSize)

	return laptopBags, metadata, nil
}
