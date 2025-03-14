package product

type ProductInput struct {
	Name     string `json:"name" binding:"required"`
	Sku      string `json:"sku" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Location string `json:"location" binding:"required"`
	Status   string `json:"status" binding:"required"`
}
