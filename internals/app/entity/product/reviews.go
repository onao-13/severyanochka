package product

import "time"

// Reviews Таблица отзывов к product.Products
type Reviews struct {
	Id          int64
	Name        string
	Description string
	Date        time.Time
	Rating      int64
	ProductId   int64
}
