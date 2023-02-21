package product

import "database/sql"

// TODO: CHANGE NULL STRING TO STRING
// Products Таблица продуктов
type Products struct {
	Id         int64          `json:"id"`
	Name       sql.NullString `json:"name"`
	Price      float64        `json:"price"`
	ImageUrl   sql.NullString `json:"imageUrl"`
	Article    int64          `json:"article"` //Артикул продукта
	Rating     float64        `json:"rating"`
	Brand      sql.NullString `json:"brand"`
	Country    sql.NullString `json:"country"`
	Weight     int64          `json:"weight"`
	Sale       int64          `json:"sale"`
	Reviews    []Reviews      `json:"reviews"`
	CategoryId int64          `json:"categoryId"`
}
