package response

// ProductResponse Обертка product.Products для отправки на сайт
type ProductResponse struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	ImageUrl string  `json:"imageUrl"`
	Price    float64 `json:"price"`
	Sale     int64   `json:"sale"`
	Rating   float64 `json:"rating"`
	County   string  `json:"county"`
}
