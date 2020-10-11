package promotion_rules

import (
	"promotionengine/model"
	"testing"
)

func TestCombinationOfTwo(t *testing.T) {

	productA := model.Product{
		ProductName: "A",
		Price:       15,
	}

	productB := model.Product{
		ProductName: "B",
		Price:       45,
	}

	productC := model.Product{
		ProductName: "C",
		Price:       34,
	}

	type args struct {
		productA       string
		productB       string
		promotionPrice int
		cart           model.Cart
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Normal Scenario", args: args{
			productA:       "A",
			productB:       "B",
			promotionPrice: 50,
			cart: model.Cart{ListItems: map[string]model.Item{
				"A": {
					ProductInCart: productA,
					Quantity:      1,
				},
				"B": {
					ProductInCart: productB,
					Quantity:      1,
				},
			}},
		},
			want: 10,
		},
		{name: "No discount applied scenario", args: args{
			productA:       "A",
			productB:       "B",
			promotionPrice: 50,
			cart: model.Cart{ListItems: map[string]model.Item{
				"A": {
					ProductInCart: productA,
					Quantity:      1,
				},
				"C": {
					ProductInCart: productC,
					Quantity:      1,
				},
			}},
		},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationOfTwo(tt.args.productA, tt.args.productB, tt.args.promotionPrice, tt.args.cart); got != tt.want {
				t.Errorf("CombinationOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNofSame(t *testing.T) {
	productA := model.Product{
		ProductName: "A",
		Price:       15,
	}

	productB := model.Product{
		ProductName: "B",
		Price:       45,
	}

	productC := model.Product{
		ProductName: "C",
		Price:       34,
	}

	type args struct {
		n              int
		productName    string
		promotionPrice int
		cart           model.Cart
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Normal Scenario", args: args{
			n:              2,
			productName:    "A",
			promotionPrice: 10,
			cart: model.Cart{ListItems: map[string]model.Item{
				"A": {
					ProductInCart: productA,
					Quantity:      2,
				},
				"B": {
					ProductInCart: productB,
					Quantity:      1,
				},
			}},
		},
			want: 20,
		},
		{name: "No discount applied scenario", args: args{
			n:              2,
			productName:    "B",
			promotionPrice: 50,
			cart: model.Cart{ListItems: map[string]model.Item{
				"A": {
					ProductInCart: productA,
					Quantity:      1,
				},
				"C": {
					ProductInCart: productC,
					Quantity:      1,
				},
			}},
		},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NofSame(tt.args.n, tt.args.productName, tt.args.promotionPrice, tt.args.cart); got != tt.want {
				t.Errorf("NofSame() = %v, want %v", got, tt.want)
			}
		})
	}
}
