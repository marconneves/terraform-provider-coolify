package database

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateDiagFunc(validateFunc func(interface{}, string) ([]string, []error)) schema.SchemaValidateDiagFunc {
	return func(i interface{}, path cty.Path) diag.Diagnostics {
		warnings, errs := validateFunc(i, fmt.Sprintf("%+v", path))
		var diags diag.Diagnostics
		for _, warning := range warnings {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  warning,
			})
		}
		for _, err := range errs {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}
		return diags
	}
}

func ValidateEngine(value interface{}, key string) (_ []string, _ []error) {
	var errs []error
	var warns []string
	// // 1. Validate if is a string
	// 2. Explode the string by ":"
	// 3. Validate if the first part is a valid engine
	// 4. Validate if the second part is a valid version from engine reference

	engineBase, ok := value.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected engine to be string"))
		return warns, errs
	}
	
    engineParts := strings.Split(engineBase, ":")
	if len(engineParts) != 2 {
		errs = append(errs, fmt.Errorf("Engine not is a valid option. Got %s", value))
		return warns, errs
	}
	
	engine := engineParts[0]
	version := engineParts[1]

	engine = strings.ToLower(regexp.MustCompile(`[^a-z]+`).ReplaceAllString(engine, ""))
	enginesOptions := []string{"mongodb", "mysql", "mariadb", "postgresql", "redis", "couchdb", "edgedb"}
	if !contains(enginesOptions, engine) {
		errs = append(errs, fmt.Errorf("Engine not is a valid option. Got %s", value))
		return warns, errs
	}

	versionsOptions := map[string][]string{
		"mongodb": {"4.2", "4.4", "5.0"},
		"mysql": {"5.7", "8.0"},
		"mariadb": {"10.2", "10.3", "10.4", "10.5","10.6", "10.7", "10.8"},
		"postgresql": {"10.22.0", "11.17.0", "12.12.0", "13.8.0", "14.5.0"},
		"redis": {"5.0", "6.0", "6.2", "7.0"},
		"couchdb": {"2.3.1", "3.1.2", "3.2.2"},
		"edgedb": {"1.4", "2.0", "2.1", "latest"},
	}

	if !contains(versionsOptions[engine], version) {
		errs = append(errs, fmt.Errorf("Version not is a valid option. Got %s", value))
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