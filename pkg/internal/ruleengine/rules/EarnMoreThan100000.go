package rules

import "errors"

type EarnMoreThan100000 struct{}

func (r EarnMoreThan100000) Evaluate(params RuleParams) (bool, error) {
	if params.Income < 0 {
		return false, errors.New("income is negative")
	}
	return params.Income > 100000, nil
}
