package rules

import "errors"

type NotMoreThan3Cards struct{}

func (r NotMoreThan3Cards) Evaluate(params RuleParams) (bool, error) {
	if params.NumberOfCreditCards < 0 {
		return false, errors.New("number of cards can't be negative")
	}
	return params.NumberOfCreditCards <= 3, nil
}
