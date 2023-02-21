package site

import "time"

// Articles Таблица статей
type Articles struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	ImageUrl    string    `json:"imageUrl"`
}
