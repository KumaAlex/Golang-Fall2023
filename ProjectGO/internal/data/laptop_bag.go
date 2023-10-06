package data

import (
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
