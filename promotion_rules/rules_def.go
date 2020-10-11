package promotion_rules

import (
	"math"
	"promotionengine/model"
)

type PromotionRules interface {
	NofSame(n int, productName string, promotionPrice int, cart model.Cart) int
	CombinationOfTwo(productA string, productB string, promotionPrice int, cart model.Cart) int
}

type PromotionRulesImpl struct {
}

func (pri *PromotionRulesImpl) NofSame(n int, productName string, promotionPrice int, cart model.Cart) int {

	if val, ok := cart.Items[productName]; ok {
		discountPerPromotion := val.ProductInCart.Price*n - promotionPrice
		return discountPerPromotion * (val.Quantity / n)
	}
	return 0
}

func (pri *PromotionRulesImpl) CombinationOfTwo(productA string, productB string, promotionPrice int, cart model.Cart) int {

	if productAItem, okA := cart.Items[productA]; okA {
		if productBItem, okB := cart.Items[productB]; okB {
			discountPerPromotion := (productAItem.ProductInCart.Price + productBItem.ProductInCart.Price) - promotionPrice
			quantityProductsOnPromotion := math.Min(float64(productAItem.Quantity), float64(productBItem.Quantity))
			return discountPerPromotion * int(quantityProductsOnPromotion)
		}
	}
	return 0
}
