package product

import "time"

// Reviews Таблица отзывов к product.Products
type Reviews struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Rating      int64     `json:"rating"`
	ProductId   int64     `json:"productId"`
}
