package rules

import "errors"

type PhoneNumberInAllowedArea struct{}

var allowedAreaCodes = []byte{'0', '2', '5', '8'}

func (r PhoneNumberInAllowedArea) Evaluate(params RuleParams) (bool, error) {
	areaCode := params.PhoneNumber[0]
	if !isNumber(areaCode) {
		return false, errors.New("area code should be in between 0-9")
	}
	return r.isAllowedAreaCode(areaCode), nil
}

func isNumber(areaCode byte) bool {
	return areaCode >= '0' && areaCode <= '9'
}

func (r PhoneNumberInAllowedArea) isAllowedAreaCode(areaCode byte) bool {
	for _, code := range allowedAreaCodes {
		if code == areaCode {
			return true
		}
	}
	return false
}
