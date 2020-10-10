package service

import (
	"log"
	"promotionengine/input_parser"
	"promotionengine/model"
	"promotionengine/promotion_rules"
)

type Service interface {
	initialize() error
	runScenarioA()
	runScenarioB()
	runScenarioC()
}

type ServiceImpl struct {
	rules    []model.Rule
	products map[string]model.Product
}

func (s *ServiceImpl) RunScenarioA() {

	log.Println("Running Scenario A")

	var cart = model.Cart{}
	cart.ListItems = []model.Item{
		{
			ProductInCart: s.products["A"],
			Quantity:      1,
		},
		{
			ProductInCart: s.products["B"],
			Quantity:      1,
		},
		{
			ProductInCart: s.products["C"],
			Quantity:      1,
		},
	}

	total := s.calculateTotal(cart)
	log.Printf("Scenario A total is %d", total)

}

func (s *ServiceImpl) RunScenarioB() {

	log.Println("Running Scenario B")

	var cart = model.Cart{}
	cart.ListItems = []model.Item{
		{
			ProductInCart: s.products["A"],
			Quantity:      5,
		},
		{
			ProductInCart: s.products["B"],
			Quantity:      5,
		},
		{
			ProductInCart: s.products["C"],
			Quantity:      1,
		},
	}

	total := s.calculateTotal(cart)
	log.Printf("Scenario B total is %d", total)

}

func (s *ServiceImpl) RunScenarioC() {

	log.Println("Running Scenario C")

	var cart = model.Cart{}
	cart.ListItems = []model.Item{
		{
			ProductInCart: s.products["A"],
			Quantity:      3,
		},
		{
			ProductInCart: s.products["B"],
			Quantity:      5,
		},
		{
			ProductInCart: s.products["C"],
			Quantity:      1,
		},
		{
			ProductInCart: s.products["D"],
			Quantity:      1,
		},
	}

	total := s.calculateTotal(cart)
	log.Printf("Scenario C total is %d", total)

}

func (s *ServiceImpl) Initialize() error {

	log.Println("Promotion Engine Initializing")
	var err error
	//parse and store the rules csv
	s.rules, err = promotion_rules.InitializeRules()
	if err != nil {
		return err
	}

	s.products, err = input_parser.InitializeProductDB("input/product.csv")
	if err != nil {
		return err
	}

	log.Println("Promotion Engine Initialized")

	return nil
}

func (s *ServiceImpl) calculateTotal(cart model.Cart) int {
	total := 0
	for _, v := range cart.ListItems {
		total += v.Quantity * v.ProductInCart.Price
	}

	return total
}
