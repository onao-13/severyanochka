package site

// SpecialOffers Таблица специальных предложений
type SpecialOffers struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}
