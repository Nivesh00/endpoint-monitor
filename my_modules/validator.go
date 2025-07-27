package my_modules

import (
	"errors"
	"regexp"
)

// Checks to see if in expressions are in the response
// and if not_in expressions are not in the response
func ValidateResponse(response *string, in *[]string, not_in *[]string) (bool, error) {

	for _, item := range *in {

		re := regexp.MustCompile(`.*?` + item + `.*?`)
		if len(re.FindAllString(*response, 1)) == 0 {
			return false, errors.New("Expression " + item + " not found in response string")
		}
	}

	for _, item := range *not_in {

		re := regexp.MustCompile(`.*?` + item + `.*?`)
		if len(re.FindAllString(*response, 1)) == 1 {
			return false, errors.New("Expression " + item + " found in response string")
		}
	}

	return true, nil
}