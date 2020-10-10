package model

type Rules struct {
	listRules []func(*Cart)
}

type Cart struct {
	ListItems []Item
	total     float64
}

type Item struct {
	ProductName string
	Price       int
	Quantity    int
}
