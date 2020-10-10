package promotion_rules

import "promotionengine/model"

func NofSame(n int, productName string, promotionPrice int, cart model.Cart) {

	for _, item := range cart.ListItems {

		if item.ProductName == productName && item.Quantity >= n {
			numberOfPromotions := promotionPrice * (item.Quantity % n)
			item.Price = numberOfPromotions*promotionPrice + (item.Quantity/n)*item.Price
		}
	}

}

func CombinationOfTwo(productA string, productB string, cart model.Cart) {

}
