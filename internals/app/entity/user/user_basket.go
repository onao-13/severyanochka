package user

// UserBasket Таблица product.Products, которые добавил Users в корзину
type UserBasket struct {
	Id        int64
	UserId    int64
	ProductId int64
}
