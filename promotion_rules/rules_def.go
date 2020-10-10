package promotion_rules

import "promotionengine/model"

func NofSame(n int, productName string, promotionPrice int, cart model.Cart) {

	for _, item := range cart.ListItems {

		if item.ProductInCart.ProductName == productName && item.Quantity >= n {
			numberOfPromotions := promotionPrice * (item.Quantity % n)
			item.ProductInCart.Price = numberOfPromotions*promotionPrice + (item.Quantity/n)*item.ProductInCart.Price
		}
	}

}

func CombinationOfTwo(productA string, productB string, cart model.Cart) {

}
