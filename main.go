package main

import (
	"promotionengine/model"
)

func main() {
	s := model.ServiceImpl{}
	s.Initialize()

	s.RunScenario("ScenarioA")
}
