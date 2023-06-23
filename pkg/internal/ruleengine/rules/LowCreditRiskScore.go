package rules

import (
	"errors"

	"github.com/honestbank/tech-assignment-backend-engineer/pkg/internal/risk"
)

type LowCreditRiskScore struct{}

func (r LowCreditRiskScore) Evaluate(params RuleParams) (bool, error) {
	if params.Age < 0 || params.NumberOfCreditCards < 0 {
		return false, errors.New("age and numbers of cards should be non negative")
	}
	return risk.CalculateCreditRisk(params.Age, params.NumberOfCreditCards) == risk.LOW, nil
}
