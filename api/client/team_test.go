package client

import (
	"testing"
)

func TestListTeam(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		Error  bool
	}{
		"ValidRequest": {
			Host:   host,
			ApiKey: apiKey,
			Error:  false,
		},
		"WithoutHost": {
			Host:   host,
			ApiKey: "",
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(testComponent.Host, testComponent.ApiKey)

			_, errors := client.team.List()

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s)  did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestGetTeam(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		Id     int
		Error  bool
	}{
		"ValidRequest": {
			Host:   host,
			ApiKey: apiKey,
			Id:     1,
			Error:  false,
		},
		"WithInvalidId": {
			Host:   host,
			ApiKey: apiKey,
			Id:     1000,
			Error:  true,
		},
		"WithoutHost": {
			Host:   host,
			ApiKey: "",
			Id:     1,
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(testComponent.Host, testComponent.ApiKey)

			_, errors := client.team.Get(testComponent.Id)

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s)  did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestGetTeamMembers(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		Id     int
		Error  bool
	}{
		"ValidRequest": {
			Host:   host,
			ApiKey: apiKey,
			Id:     1,
			Error:  false,
		},
		"WithInvalidId": {
			Host:   host,
			ApiKey: apiKey,
			Id:     1000,
			Error:  true,
		},
		"WithoutHost": {
			Host:   host,
			ApiKey: "",
			Id:     1,
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(testComponent.Host, testComponent.ApiKey)

			_, errors := client.team.Members(testComponent.Id)

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s)  did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}
