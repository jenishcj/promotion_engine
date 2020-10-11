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

	promotionRulesMock.EXPECT().NofSame(3, "A", 130, gomock.Any()).Return(20)
	promotionRulesMock.EXPECT().NofSame(2, "B", 45, gomock.Any()).Return(30)
	promotionRulesMock.EXPECT().CombinationOfTwo("C", "D", 30, gomock.Any()).Return(5)
	rules := getRules()
	products := getProducts()

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
		{
			name: "Scenario C",
			fields: fields{
				rules:    rules,
				products: products,
				pri:      promotionRulesMock,
			},
			args: args{cart: model.Cart{Items: map[string]model.Item{
				"A": {
					ProductInCart: products["A"],
					Quantity:      3,
				},
				"B": {
					ProductInCart: products["B"],
					Quantity:      5,
				},
				"C": {
					ProductInCart: products["C"],
					Quantity:      1,
				},
				"D": {
					ProductInCart: products["D"],
					Quantity:      1,
				},
			}},
			},
			want: 280,
		},
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

func getProducts() map[string]model.Product {

	productA := model.Product{
		ProductName: "A",
		Price:       50,
	}

	productB := model.Product{
		ProductName: "B",
		Price:       30,
	}

	productC := model.Product{
		ProductName: "C",
		Price:       20,
	}

	productD := model.Product{
		ProductName: "D",
		Price:       15,
	}

	return map[string]model.Product{"A": productA, "B": productB, "C": productC, "D": productD}

}

func getRules() []model.Rule {
	//nOfSame,3,A,130
	r1 := model.Rule{
		RuleFuncName: "nOfSame",
		FuncParams:   []interface{}{3, "A", 130},
	}

	//nOfSame,2,B,45
	r2 := model.Rule{
		RuleFuncName: "nOfSame",
		FuncParams:   []interface{}{2, "B", 45},
	}

	//combinationOfTwo,C,D,30
	r3 := model.Rule{
		RuleFuncName: "combinationOfTwo",
		FuncParams:   []interface{}{"C", "D", 30},
	}

	return []model.Rule{r1, r2, r3}
}
