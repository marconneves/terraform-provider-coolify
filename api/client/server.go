package client

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

type ServerInstance struct {
	client *Client
}

func (c *Client) Server() *ServerInstance {
	return &ServerInstance{client: c}
}

type Server struct {
	UUID                          string         `json:"uuid"`
	Name                          string         `json:"name"`
	Description                   *string        `json:"description"`
	HighDiskUsageNotificationSent bool           `json:"high_disk_usage_notification_sent"`
	IP                            string         `json:"ip"`
	LogDrainNotificationSent      bool           `json:"log_drain_notification_sent"`
	Port                          string         `json:"port"`
	PrivateKeyID                  int            `json:"private_key_id"`
	Proxy                         *Proxy         `json:"proxy"`
	Settings                      ServerSettings `json:"settings"`
	SwarmCluster                  *string        `json:"swarm_cluster"`
	TeamID                        int            `json:"team_id"`
	UnreachableCount              int            `json:"unreachable_count"`
	UnreachableNotificationSent   bool           `json:"unreachable_notification_sent"`
	User                          string         `json:"user"`
	ValidationLogs                *string        `json:"validation_logs"`
	CreatedAt                     time.Time      `json:"created_at"`
	UpdatedAt                     time.Time      `json:"updated_at"`
}

type Proxy struct {
	Status    string `json:"status"`
	Type      string `json:"type"`
	ForceStop bool   `json:"force_stop"`
}

type ServerSettings struct {
	Id                         int       `json:"id"`
	ConcurrentBuilds           int       `json:"concurrent_builds"`
	DeleteUnusedNetworks       bool      `json:"delete_unused_networks"`
	DeleteUnusedVolumes        bool      `json:"delete_unused_volumes"`
	DockerCleanupFrequency     string    `json:"docker_cleanup_frequency"`
	DockerCleanupThreshold     int       `json:"docker_cleanup_threshold"`
	DynamicTimeout             int       `json:"dynamic_timeout"`
	ForceDisabled              bool      `json:"force_disabled"`
	ForceDockerCleanup         bool      `json:"force_docker_cleanup"`
	GenerateExactLabels        bool      `json:"generate_exact_labels"`
	IsBuildServer              bool      `json:"is_build_server"`
	IsCloudflareTunnel         bool      `json:"is_cloudflare_tunnel"`
	IsJumpServer               bool      `json:"is_jump_server"`
	IsLogdrainAxiomEnabled     bool      `json:"is_logdrain_axiom_enabled"`
	IsLogdrainCustomEnabled    bool      `json:"is_logdrain_custom_enabled"`
	IsLogdrainHighlightEnabled bool      `json:"is_logdrain_highlight_enabled"`
	IsLogdrainNewRelicEnabled  bool      `json:"is_logdrain_newrelic_enabled"`
	IsMetricsEnabled           bool      `json:"is_metrics_enabled"`
	IsReachable                bool      `json:"is_reachable"`
	IsServerAPIEnabled         bool      `json:"is_server_api_enabled"`
	IsSwarmManager             bool      `json:"is_swarm_manager"`
	IsSwarmWorker              bool      `json:"is_swarm_worker"`
	IsUsable                   bool      `json:"is_usable"`
	LogdrainAxiomApiKey        *string   `json:"logdrain_axiom_api_key"`
	LogdrainAxiomDatasetName   *string   `json:"logdrain_axiom_dataset_name"`
	LogdrainCustomConfig       *string   `json:"logdrain_custom_config"`
	LogdrainCustomConfigParser *string   `json:"logdrain_custom_config_parser"`
	LogdrainHighlightProjectId *string   `json:"logdrain_highlight_project_id"`
	LogdrainNewRelicBaseUri    *string   `json:"logdrain_newrelic_base_uri"`
	LogdrainNewRelicLicenseKey *string   `json:"logdrain_newrelic_license_key"`
	MetricsHistoryDays         int       `json:"metrics_history_days"`
	MetricsRefreshRateSeconds  int       `json:"metrics_refresh_rate_seconds"`
	MetricsToken               string    `json:"metrics_token"`
	ServerId                   int       `json:"server_id"`
	ServerTimezone             string    `json:"server_timezone"`
	WildcardDomain             *string   `json:"wildcard_domain"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

func (t *ServerInstance) List() (*[]Server, error) {
	body, err := t.client.httpRequest("servers", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]Server{})
}

func (t *ServerInstance) Get(uuid string) (*Server, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.httpRequest(fmt.Sprintf("servers/%v", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &Server{})
}

type CreateServerDTO struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	IP              string `json:"ip"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	PrivateKeyUUID  string `json:"private_key_uuid"`
	IsBuildServer   bool   `json:"is_build_server"`
	InstantValidate bool   `json:"instant_validate"`
}

type CreateServerResponse struct {
	UUID string `json:"uuid"`
}

func (t *ServerInstance) Create(server *CreateServerDTO) (*string, error) {
	buf, err := encodeRequest(server)
	if err != nil {
		return nil, err
	}

	body, err := t.client.httpRequest("servers", "POST", *buf)
	if err != nil {
		return nil, err
	}

	response, err := decodeResponse(body, &CreateServerResponse{})
	if err != nil {
		return nil, err
	}

	return &response.UUID, nil
}

func (t *ServerInstance) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	_, err := t.client.httpRequest(fmt.Sprintf("servers/%v", uuid), "DELETE")
	if err != nil {
		return err
	}

	return nil
}

type UpdateServerDTO struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	IP              string `json:"ip"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	PrivateKeyUUID  string `json:"private_key_uuid"`
	IsBuildServer   bool   `json:"is_build_server"`
	InstantValidate bool   `json:"instant_validate"`
}

func (t *ServerInstance) Update(uuid string, server *UpdateServerDTO) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	buf, err := encodeRequest(server)
	if err != nil {
		return err
	}

	_, err = t.client.httpRequest(fmt.Sprintf("servers/%v", uuid), "PATCH", *buf)
	return err
}

type Resource struct {
	Id        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

func (t *ServerInstance) Resources(uuid string) (*[]Resource, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.httpRequest(fmt.Sprintf("servers/%v/resources", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]Resource{})
}

type Domain struct {
	Id      int      `json:"id"`
	Domains []string `json:"domains"`
}

func (t *ServerInstance) Domains(uuid string) (*[]Domain, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.httpRequest(fmt.Sprintf("servers/%v/domains", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]Domain{})
}

func (t *ServerInstance) Validate(uuid string) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	_, err := t.client.httpRequest(fmt.Sprintf("servers/%v/validate", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return err
	}

	return nil
}
