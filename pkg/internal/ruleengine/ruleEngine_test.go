package ruleengine

import (
	"errors"
	"fmt"
	"testing"

	"github.com/honestbank/tech-assignment-backend-engineer/pkg/internal/ruleengine/rules"
	"github.com/stretchr/testify/assert"
)

func TestPrepareRuleEngine(t *testing.T) {
	tests := []struct {
		Name           string
		ExpectedResult RuleEngine
	}{
		{
			Name: "Test Rule Engine is prepared properly",
			ExpectedResult: RuleEngine{
				rules: []rules.Rule{
					rules.PreApprovedPhoneNumbers{},
					rules.EarnMoreThan100000{},
					rules.Atleast18{},
					rules.NotMoreThan3Cards{},
					rules.LowCreditRiskScore{},
					rules.NotPoliticallyExposed{},
					rules.PhoneNumberInAllowedArea{},
				},
			},
		},
	}

	for _, test := range tests {
		result := PrepareNewRuleEngine()
		assert.Equal(t, test.ExpectedResult, result, test.Name)
	}
}

func TestValidateCreditCardApplication(t *testing.T) {
	tests := []struct {
		Name           string
		Input          rules.RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "All parmas are good and valid",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Age is less than 18",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 18,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Age is negative",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 -18,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  fmt.Errorf("Error: %w", errors.New("age is negative")),
		},
		{
			Name: "Income is less than 100000",
			Input: rules.RuleParams{
				Income:              34560,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Income is negative",
			Input: rules.RuleParams{
				Income:              -100,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  fmt.Errorf("Error: %w", errors.New("income is negative")),
		},
		{
			Name: "Number of credit cards are 4 ",
			Input: rules.RuleParams{
				Income:              234500,
				Age:                 20,
				NumberOfCreditCards: 4,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Number of credit card is negative",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: -1,
				PoliticallyExposed:  false,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  fmt.Errorf("Error: %w", errors.New("number of cards can't be negative")),
		},
		{
			Name: "Is Politically Exposed",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  true,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is in area code 8",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "876-376-123",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is in invalid area code",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  false,
				PhoneNumber:         "-576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  fmt.Errorf("Error: %w", errors.New("area code should be in between 0-9")),
		},
		{
			Name: "Age is less than 18 and politically exposed",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 17,
				NumberOfCreditCards: 1,
				PoliticallyExposed:  true,
				PhoneNumber:         "576-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Credit Score is medium",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 2,
				PoliticallyExposed:  false,
				PhoneNumber:         "476-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is in pre-approved list, Credit Score is medium",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 2,
				PoliticallyExposed:  false,
				PhoneNumber:         "323-912-6988",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Phone number is not in pre-approved list, Credit Score is medium",
			Input: rules.RuleParams{
				Income:              239700,
				Age:                 20,
				NumberOfCreditCards: 2,
				PoliticallyExposed:  false,
				PhoneNumber:         "416-376-123",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		ruleEngine := RuleEngine{
			rules: []rules.Rule{
				rules.PreApprovedPhoneNumbers{},
				rules.EarnMoreThan100000{},
				rules.Atleast18{},
				rules.NotMoreThan3Cards{},
				rules.LowCreditRiskScore{},
				rules.NotPoliticallyExposed{},
				rules.PhoneNumberInAllowedArea{},
			},
		}
		rules.APPROVED_PHONE_NUMBERS_FILE_NAME = "pre-approved-phone-numbers_test.txt"
		result, err := ruleEngine.ValidateCreditCardApplication(test.Input)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)
	}
}
