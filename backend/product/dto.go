package product

import "errors"

type CreateProductDTO struct {
	Name     string `form:"name"`
	Sku      string `form:"sku"`
	Quantity int `form:"quantity"`
	Location string `form:"location"`
	Status   string `form:"status"`
}

func (p *CreateProductDTO) Validate() error {
	if p.Name == "" {
		return errors.New("field name is required")
	}

	if p.Sku == "" {
		return errors.New("field sku is required")
	}

	if p.Quantity < 1 || p.Quantity == 0 {
		return errors.New("field price min 1")
	}

	if p.Location == "" {
		return errors.New("field location is required")
	}

	if p.Status == "" {
		return errors.New("field status is requried")
	}

	return nil
}
