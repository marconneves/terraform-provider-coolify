package client

import (
	"testing"
)

func TestListServer(t *testing.T) {
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

			_, errors := client.server.List()

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s) did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestGetServer(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		UUID   string
		Error  bool
	}{
		"ValidRequest": {
			Host:   host,
			ApiKey: apiKey,
			UUID:   "lo4sksgsks8kw8w0skog8c0s",
			Error:  false,
		},
		"WithInvalidId": {
			Host:   host,
			ApiKey: apiKey,
			UUID:   "",
			Error:  true,
		},
		"WithoutHost": {
			Host:   host,
			ApiKey: "",
			UUID:   "",
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(testComponent.Host, testComponent.ApiKey)

			_, errors := client.server.Get(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s) did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestCreateServer(t *testing.T) {
	cases := map[string]struct {
		Server *CreateServerDTO
		Error  bool
	}{
		"ValidServer": {
			Server: &CreateServerDTO{
				Name:            "My Server",
				Description:     "My Server Description",
				IP:              "127.0.0.1",
				Port:            22,
				User:            "root",
				PrivateKeyUUID:  "fggkoowk084k8okc8wk4g4o4",
				IsBuildServer:   true,
				InstantValidate: true,
			},
			Error: false,
		},
		"MissingName": {
			Server: &CreateServerDTO{
				Description:     "No Name Server",
				IP:              "127.0.0.1",
				Port:            22,
				User:            "root",
				PrivateKeyUUID:  "fggkoowk084k8okc8wk4g4o4",
				IsBuildServer:   true,
				InstantValidate: true,
			},
			Error: true,
		},
		"InvalidPrivateKey": {
			Server: &CreateServerDTO{
				Name:            "No PrivateKey UUID Valid",
				Description:     "Invalid IP",
				IP:              "127.0.0.1",
				Port:            22,
				User:            "root",
				PrivateKeyUUID:  "asjdaksdhjaskljdha",
				IsBuildServer:   true,
				InstantValidate: true,
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(host, apiKey)

			uuid, errors := client.server.Create(testComponent.Server)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server creation failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server creation succeeded unexpectedly: %s", *uuid)
			}
		})
	}
}

func TestUpdateServer(t *testing.T) {
	cases := map[string]struct {
		UUID   string
		Server *UpdateServerDTO
		Error  bool
	}{
		"ValidUpdate": {
			UUID: "lcs8ggw8cos48kw0sc0sk0gc",
			Server: &UpdateServerDTO{
				Name:        "Updated Server",
				Description: "Updated Description",
				IP:          "192.168.1.1",
				Port:        22,
				User:        "admin",
			},
			Error: false,
		},
		"InvalidUUID": {
			UUID: "",
			Server: &UpdateServerDTO{
				Name:        "Updated Server",
				Description: "Updated Description",
				IP:          "192.168.1.1",
				Port:        22,
				User:        "admin",
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(host, apiKey)

			errors := client.server.Update(testComponent.UUID, testComponent.Server)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server update failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server update succeeded unexpectedly")
			}
		})
	}
}

func TestDeleteServer(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidDelete": {
			UUID:  "lcs8ggw8cos48kw0sc0sk0gc",
			Error: false,
		},
		"InvalidUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(host, apiKey)

			errors := client.server.Delete(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server deletion failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server deletion succeeded unexpectedly")
			}
		})
	}
}

func TestListDomains(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidRequest": {
			UUID:  "lo4sksgsks8kw8w0skog8c0s",
			Error: false,
		},
		"InvalidUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(host, apiKey)

			_, errors := client.server.Domains(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Domain listing failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Domain listing succeeded unexpectedly")
			}
		})
	}
}

func TestListResources(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		UUID   string
		Error  bool
	}{
		"ValidRequest": {
			UUID:  "lo4sksgsks8kw8w0skog8c0s",
			Error: false,
		},
		"InvalidUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = NewClient(host, apiKey)

			_, errors := client.server.Resources(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Resource listing failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Resource listing succeeded unexpectedly")
			}
		})
	}
}
