package promotion_rules

import (
	"promotionengine/model"
)

//TODO: Futurely can be driven by a CSV the rules in the form as example csv rules.csv
func InitializeRules() ([]model.Rule, error) {
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

	return []model.Rule{r1, r2, r3}, nil
}
