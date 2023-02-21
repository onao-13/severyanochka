package product

// Categories Таблица категории product.Products
type Categories struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}
