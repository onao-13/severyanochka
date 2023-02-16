package site

import "time"

// Articles Таблица статей
type Articles struct {
	Id          int64
	Title       string
	Description string
	Date        time.Time
	ImageUrl    string
}
