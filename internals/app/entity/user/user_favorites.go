package user

// UserFavorites Таблица избранных product.Products Users
type UserFavorites struct {
	Id        int64
	UserId    int64
	ProductId int64
}
