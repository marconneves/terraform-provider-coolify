package database

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
		"ValidExistingEngine": {
			Value: "mongodb",
			Error: false,
		},
		"ValidExistingEngineWithTrace": {
			Value: "mongo-db",
			Error: false,
		},
		"NotIsValidEngine": {
			Value: "oracledb",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			_, errors := ValidateEngine(testComponent.Value, testName)

			if len(errors) > 0 && !testComponent.Error {
				t.Errorf("ValidateEngine(%s) produced an unexpected error", testComponent.Value)
			} else if len(errors) == 0 && testComponent.Error {
				t.Errorf("ValidateEngine(%s) did not error", testComponent.Value)
			}
		})
	}
}
