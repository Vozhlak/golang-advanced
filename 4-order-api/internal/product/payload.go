package product

import "github.com/lib/pq"

type RequestProduct struct {
	Name        string         `json:"name" validate:"required,max=100,min=2"`
	Description string         `json:"description" validate:"required,max=1000,min=10"`
	Images      pq.StringArray `json:"images" validate:"omitempty,dive,url"`
	Price       float64        `json:"price" validate:"required,gt=0"`
}

type RequestUpdateProduct struct {
	Name        string         `json:"name,omitempty" validate:"omitempty,max=100,min=2"`
	Description string         `json:"description,omitempty" validate:"omitempty,max=1000,min=10"`
	Images      pq.StringArray `json:"images,omitempty" validate:"omitempty,dive,url"`
	Price       float64        `json:"price,omitempty" validate:"omitempty,gt=0"`
}
