package shared

import (
	"fmt"
	"regexp"
)

func ValidateName(valueObject interface{}, key string) (_ []string, _ []error) {
	var errs []error
	var warns []string
	value, ok := valueObject.(string)

	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	
	return warns, errs
}