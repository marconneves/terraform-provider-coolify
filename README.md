# Terraform Provider Coolify

## Visão Geral
O **Terraform Provider Coolify** permite que você gerencie recursos do Coolify (servidores, aplicações, bancos de dados, etc.) como código (IaC). Este provedor integra-se à API do Coolify para automatizar o provisionamento e a configuração da sua infraestrutura.

## Tecnologias
- **Linguagem:** Go (Golang)
- **Framework:** Terraform Plugin Framework (hashicorp/terraform-plugin-framework)
- **SDK:** coolify-sdk-go (SDK customizado para a API do Coolify)
- **Gestão de Dependências:** Go Modules

## Como rodar localmente

### Pré-requisitos
- [Go](https://golang.org/doc/install) instalando (versão recomendada no `.go-version`)
- [Terraform](https://developer.hashicorp.com/terraform/downloads) (versão 1.0+)
- [Docker](https://docs.docker.com/get-docker/) e Docker Compose (para testes locais)

### Passo a Passo
1. **Clone o repositório:**
   ```bash
   git clone https://github.com/marconneves/terraform-provider-coolify.git
   cd terraform-provider-coolify
   ```

2. **Compile o provedor:**
   ```bash
   go build -o terraform-provider-coolify
   ```

3. **Configuração para desenvolvimento local:**
   Para testar o provedor localmente sem publicá-lo, você pode usar um arquivo `.terraformrc` no seu home directory:
   ```hcl
   provider_installation {
     dev_overrides {
       "marconneves/coolify" = "/caminho/para/o/binario/do/provedor"
     }
     direct {}
   }
   ```

4. **Rodando testes com Docker:**
   O projeto utiliza Docker Compose para subir instâncias locais do Coolify ou bancos de dados para testes de aceitação:
   ```bash
   docker compose up -d
   ```
   Execute os testes:
   ```bash
   TF_ACC=1 go test ./... -v
   ```

## Variáveis de Ambiente
As seguintes variáveis de ambiente são utilizadas para configurar o provedor:

| Variável | Descrição | Obrigatório |
|----------|-----------|-------------|
| `COOLIFY_ENDPOINT` | URL da API do Coolify (ex: `https://app.coolify.io/api/v1`) | Sim |
| `COOLIFY_API_TOKEN` | Token de autenticação da API | Sim |

## Guia de Contribuição
1. Faça um Fork do repositório.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Siga os padrões de código (veja [AGENTS.md](./AGENTS.md)).
4. Certifique-se de que os testes estão passando.
5. Abra um Pull Request detalhando as mudanças.

---
Desenvolvido com ❤️ pela comunidade Coolify.