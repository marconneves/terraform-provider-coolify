package database

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateEngine(value interface{}, key string) (_ []string, _ []error) {
	var errs []error
	var warns []string
	engineBase, ok := value.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected engine to be string"))
		return warns, errs
	}

	engine := strings.ToLower(regexp.MustCompile(`[^a-z]+`).ReplaceAllString(engineBase, ""))

	enginesOptions := []string{"mongodb", "mysql", "mariadb", "postgresql", "redis", "couchdb", "edgedb"}

	if !contains(enginesOptions, engine) {
		errs = append(errs, fmt.Errorf("Engine not is a valid option. Got %s", value))
		return warns, errs
	}

	return warns, errs
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}