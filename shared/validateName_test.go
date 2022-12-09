package shared

import (
	"testing"
)


func TestValidateEngine(t *testing.T) {
	cases := map[string]struct {
		Value interface{}
		Error bool
	}{
		"NotString": {
			Value: 7,
			Error: true,
		},
		"ValidNameWithSpace": {
			Value: "Valid Name",
			Error: false,
		},
		"ValidNameWithTrace": {
			Value: "valid-name",
			Error: false,
		},
		"ValidNameWithUnderscore": {
			Value: "valid_name",
			Error: false,
		},
		"NotIsValidNameWithEmoji": {
			Value: "Invalid Name 🎉",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			_, errors := ValidateName(testComponent.Value, testName)

			if len(errors) > 0 && !testComponent.Error {
				t.Errorf("Valid Name(%s) produced an unexpected error", testComponent.Value)
			} else if len(errors) == 0 && testComponent.Error {
				t.Errorf("Valid Name(%s) did not error", testComponent.Value)
			}
		})
	}
}
