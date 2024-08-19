package schemas

type ClothCreateUpdateSchema struct {
	Name  string  `json:"name" binding:"required,max=255"`
	Color string  `json:"color" binding:"required,max=100"`
	Size  string  `json:"size" binding:"required,max=5"`
	Price float64 `json:"price" binding:"required,min=0"`
	Stock *int    `json:"stock" binding:"omitempty,min=0"`
}