package data

import (
	"ProjectGO/internal/validator"
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

func ValidateMovie(v *validator.Validator, laptopBag *LaptopBag) {
	v.Check(laptopBag.Brand != "", "brand", "must be provided")
	v.Check(len(laptopBag.Brand) <= 500, "brand", "must not be more than 500 bytes long")
	v.Check(laptopBag.Color != "", "color", "must be provided")
	v.Check(len(laptopBag.Color) <= 64, "color", "must not be more than 64 bytes long")
	v.Check(laptopBag.Weight != 0, "weight", "must be provided")
	v.Check(laptopBag.Weight >= 0, "weight", "must be greater than 0")
	v.Check(laptopBag.Dimensions != nil, "dimensions", "must be provided")
	v.Check(len(laptopBag.Dimensions) == 3, "genres", "must contain exactly 3 dimensions")
}
