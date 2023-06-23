package rules

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneNumberInAllowedAreaRule(t *testing.T) {
	tests := []struct {
		Name           string
		TestInput      RuleParams
		ExpectedResult bool
		ExpectedError  error
	}{
		{
			Name: "Phone Number is in area code 2",
			TestInput: RuleParams{
				PhoneNumber: "286-356-0375",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Phone Number is in area code 0",
			TestInput: RuleParams{
				PhoneNumber: "086-356-0375",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Phone Number is in area code 5",
			TestInput: RuleParams{
				PhoneNumber: "586-356-0375",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Phone Number is in area code 8",
			TestInput: RuleParams{
				PhoneNumber: "886-356-0375",
			},
			ExpectedResult: true,
			ExpectedError:  nil,
		},
		{
			Name: "Phone Number is in other area code",
			TestInput: RuleParams{
				PhoneNumber: "486-356-0375",
			},
			ExpectedResult: false,
			ExpectedError:  nil,
		},
		{
			Name: "Phone Number is in invalid area code",
			TestInput: RuleParams{
				PhoneNumber: "-486-356-0375",
			},
			ExpectedResult: false,
			ExpectedError:  errors.New("area code should be in between 0-9"),
		},
	}

	for _, test := range tests {
		phoneNumberInAllowedAreaRule := PhoneNumberInAllowedArea{}
		result, err := phoneNumberInAllowedAreaRule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)
	}
}
