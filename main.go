package main

import (
	"log"
	"promotionengine/service"
)

func main() {
	var s = service.ServiceImpl{}

	err := s.Initialize()
	if err != nil {
		log.Fatal(err)
	} else {

		//Running scenarios
		s.RunScenarioA()
		s.RunScenarioB()
		s.RunScenarioC()
	}
}
