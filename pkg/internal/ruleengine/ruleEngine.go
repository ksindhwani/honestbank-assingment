package ruleengine

import (
	"fmt"

	"github.com/honestbank/tech-assignment-backend-engineer/pkg/internal/ruleengine/rules"
)

type RuleEngine struct {
	rules []rules.Rule
}

func PrepareNewRuleEngine() RuleEngine {
	re := RuleEngine{}
	re.rules = append(re.rules, rules.PreApprovedPhoneNumbers{})
	re.rules = append(re.rules, rules.EarnMoreThan100000{})
	re.rules = append(re.rules, rules.Atleast18{})
	re.rules = append(re.rules, rules.NotMoreThan3Cards{})
	re.rules = append(re.rules, rules.LowCreditRiskScore{})
	re.rules = append(re.rules, rules.NotPoliticallyExposed{})
	re.rules = append(re.rules, rules.PhoneNumberInAllowedArea{})
	return re
}

func (re *RuleEngine) ValidateCreditCardApplication(params rules.RuleParams) (bool, error) {
	for _, rule := range re.rules {
		ok, err := rule.Evaluate(params)
		if err != nil {
			return false, fmt.Errorf("Error: %w", err)
		}

		_, isPreApprovedPhoneNumberRule := rule.(rules.PreApprovedPhoneNumbers)
		if isPreApprovedPhoneNumberRule && ok {
			return true, nil
		}

		if !isPreApprovedPhoneNumberRule && !ok {
			return false, nil
		}

	}
	return true, nil
}
