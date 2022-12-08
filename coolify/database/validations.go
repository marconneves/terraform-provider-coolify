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

func ValidateEngineImage(value interface{}, key string) (_ []string, _ []error) {
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