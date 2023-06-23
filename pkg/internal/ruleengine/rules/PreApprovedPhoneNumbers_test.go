package rules

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreApprovedPhoneNumbersRule(t *testing.T) {
	tests := []struct {
		Name                string
		TestInput           RuleParams
		PreApprovedFileName string
		ExpectedResult      bool
		ExpectedError       error
	}{
		{
			Name: "Phone number is pre approved",
			TestInput: RuleParams{
				PhoneNumber: "323-912-6988",
			},
			PreApprovedFileName: "pre-approved-phone-numbers_test.txt",
			ExpectedResult:      true,
			ExpectedError:       nil,
		},
		{
			Name: "Phone number is not pre approved",
			TestInput: RuleParams{
				PhoneNumber: "763-912-6988",
			},
			PreApprovedFileName: "pre-approved-phone-numbers_test.txt",
			ExpectedResult:      false,
			ExpectedError:       nil,
		},
		{
			Name: "Pre Approved File doesn't exist",
			TestInput: RuleParams{
				PhoneNumber: "763-912-6988",
			},
			PreApprovedFileName: "pre-approved-phone-numbers_test123.txt",
			ExpectedResult:      false,
			ExpectedError: fmt.Errorf("unable to fetch pre approved phone numbers. %w",
				errors.New("pre-approved-phone-numbers_test123.txt doesn't exist")),
		},
	}

	for _, test := range tests {
		preApprovedPhoneNumbersRule := PreApprovedPhoneNumbers{}
		APPROVED_PHONE_NUMBERS_FILE_NAME = test.PreApprovedFileName
		result, err := preApprovedPhoneNumbersRule.Evaluate(test.TestInput)
		assert.Equal(t, test.ExpectedError, err, test.Name)
		assert.Equal(t, test.ExpectedResult, result, test.Name)

	}
}
