package model

import (
	"log"
	"promotionengine/promotion_rules"
)

type Service interface {
	initialize() error
	runScenario(string)
}

type ServiceImpl struct {
	rules []Rules
}

func (s *ServiceImpl) RunScenario(scenario string) {

	log.Println("Running Scenario " + scenario)

}

func (s *ServiceImpl) Initialize() error {

	log.Println("Promotion Engine Initializing")

	//parse and store the rules csv
	r, err := promotion_rules.Parse_rules_file_csv("rules.csv")
	if err != nil {
		return err
	}

	s.rules = r
	log.Println("Promotion Engine Initialized")

	return nil
}
