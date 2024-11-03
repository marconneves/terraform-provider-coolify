package client

import (
	"bytes"
	"fmt"
	"time"
)

type TeamInstance struct {
	client *Client
}

type Team struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *TeamInstance) List() (*[]Team, error) {
	body, err := t.client.httpRequest("teams", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]Team{})
}

func (t *TeamInstance) Get(id int) (*Team, error) {
	body, err := t.client.httpRequest(fmt.Sprintf("teams/%v", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &Team{})
}

type Member struct {
	Id                   int     `json:"id"`                      // Identificador do usuário no banco de dados
	Name                 string  `json:"name"`                    // Nome do usuário
	Email                string  `json:"email"`                   // Email do usuário
	EmailVerifiedAt      *string `json:"email_verified_at"`       // Data de verificação do email do usuário
	CreatedAt            string  `json:"created_at"`              // Data de criação do usuário
	UpdatedAt            string  `json:"updated_at"`              // Data de atualização do usuário
	TwoFactorConfirmedAt *string `json:"two_factor_confirmed_at"` // Data de confirmação do segundo fator de autenticação
	ForcePasswordReset   bool    `json:"force_password_reset"`    // Flag para forçar redefinição de senha
	MarketingEmails      bool    `json:"marketing_emails"`        // Flag para receber emails de marketing
}

func (t *TeamInstance) GetMembers(id int) (*[]Member, error) {
	body, err := t.client.httpRequest(fmt.Sprintf("teams/%v/members", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]Member{})
}
