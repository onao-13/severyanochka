package product

import "database/sql"

// Products Таблица продуктов
type Products struct {
	Id       int64
	Name     sql.NullString
	Price    float64
	ImageUrl sql.NullString
	Article  int64 //Артикул продукта
	Rating   float64
	Brand    sql.NullString
	Country  sql.NullString
	Weight   int64
	Sale     int64
	//TODO: FIX THIS!!!
	//Reviews    []Reviews
	//CategoryId Categories
}
