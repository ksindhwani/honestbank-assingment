package rules

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var APPROVED_PHONE_NUMBERS_FILE_NAME = "pre-approved-phone-numbers.txt"

type PreApprovedPhoneNumbers struct{}

func (r PreApprovedPhoneNumbers) Evaluate(params RuleParams) (bool, error) {
	approvedPhoneNumbers, err := fetchPreApprovedPhoneNumbers()
	if err != nil {
		return false, fmt.Errorf("unable to fetch pre approved phone numbers. %w", err)
	}
	for _, preApprovedPhoneNumber := range approvedPhoneNumbers {
		if params.PhoneNumber == preApprovedPhoneNumber {
			return true, nil
		}
	}
	return false, nil

}

func fetchPreApprovedPhoneNumbers() ([]string, error) {
	rootDir, err := findRootDirectory()
	if err != nil {
		return nil, fmt.Errorf("root directory not found. %w", err)
	}
	filePath := filepath.Join(rootDir, APPROVED_PHONE_NUMBERS_FILE_NAME)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New(APPROVED_PHONE_NUMBERS_FILE_NAME + " doesn't exist")
	}
	approvedPhoneNumbers := strings.Split(string(content), "\n")
	return approvedPhoneNumbers, nil
}

func findRootDirectory() (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Start from the current directory and traverse upwards until we find a go.mod file
	for {
		// Check if go.mod file exists in the current directory
		goModPath := filepath.Join(cwd, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return cwd, nil
		}

		// Move up to the parent directory
		parentDir := filepath.Dir(cwd)
		// If we reached the root directory ("/"), return an error indicating the project root was not found
		if parentDir == cwd {
			return "", fmt.Errorf("go.mod not found")
		}
		cwd = parentDir
	}
}
