package rules

import "errors"

type Atleast18 struct{}

func (r Atleast18) Evaluate(params RuleParams) (bool, error) {
	if params.Age < 0 {
		return false, errors.New("age is negative")
	}
	return params.Age >= 18, nil
}
