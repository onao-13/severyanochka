package user

// Users Таблица пользователей
type Users struct {
	Id        int64
	Name      string
	Number    string
	Password  string
	Favorites []UserFavorites
	Basket    []UserBasket
}
