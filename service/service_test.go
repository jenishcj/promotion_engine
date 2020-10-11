package service

import (
	"github.com/golang/mock/gomock"
	mock_promotion_rules "promotionengine/mocks"
	"promotionengine/model"
	"promotionengine/promotion_rules"
	"testing"
)

func TestServiceImpl_calculateTotal(t *testing.T) {
	mckController := gomock.NewController(t)
	promotionRulesMock := mock_promotion_rules.NewMockPromotionRules(mckController)

	promotionRulesMock.EXPECT().CombinationOfTwo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(10)
	promotionRulesMock.EXPECT().CombinationOfTwo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(5)

	type fields struct {
		rules    []model.Rule
		products map[string]model.Product
		pri      promotion_rules.PromotionRules
	}
	type args struct {
		cart model.Cart
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{
				rules:    tt.fields.rules,
				products: tt.fields.products,
				pri:      promotionRulesMock,
			}

			if got := s.calculateTotal(tt.args.cart); got != tt.want {
				t.Errorf("calculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
