package model

type Cart struct {
	ListItems map[string]Item
	total     float64
}

type Item struct {
	ProductInCart Product
	Quantity      int
}

type Product struct {
	ProductName string
	Price       int
}
