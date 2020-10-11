package promotion_rules

import (
	"math"
	"promotionengine/model"
)

func NofSame(n int, productName string, promotionPrice int, cart model.Cart) int {

	if val, ok := cart.ListItems[productName]; ok {
		discountPerPromotion := val.ProductInCart.Price*n - promotionPrice
		return discountPerPromotion * (val.Quantity / n)
	}
	return 0
}

func CombinationOfTwo(productA string, productB string, promotionPrice int, cart model.Cart) int {

	if productAItem, okA := cart.ListItems[productA]; okA {
		if productBItem, okB := cart.ListItems[productB]; okB {
			discountPerPromotion := (productAItem.ProductInCart.Price + productBItem.ProductInCart.Price) - promotionPrice
			quantityProductsOnPromotion := math.Min(float64(productAItem.Quantity), float64(productBItem.Quantity))
			return discountPerPromotion * int(quantityProductsOnPromotion)
		}
	}
	return 0
}
