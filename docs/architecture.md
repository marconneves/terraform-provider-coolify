# Arquitetura do Projeto (docs/architecture.md)

Este documento descreve a estrutura de pastas e o fluxo de dados do Terraform Provider Coolify.

## Estrutura de Pastas
```text
.
├── main.go                     # Ponto de entrada do provedor
├── coolify/                    # Lógica principal dos recursos
│   ├── [recurso]/              # Pasta específica por recurso (ex: server, database)
│   │   ├── resource.go         # Definição do Resource
│   │   ├── data_source.go      # Definição do Data Source
│   │   ├── model.go            # Mapeamento do modelo Terraform framework
│   │   ├── create.go           # Implementação do Create
│   │   ├── read.go             # Implementação do Read
│   │   ├── update.go           # Implementação do Update
│   │   └── delete.go           # Implementação do Delete
│   └── provider.go             # Configuração global do provedor
├── shared/                     # Código compartilhado entre recursos
├── internal/                    # Lógica interna e utilitários (não exportados)
├── examples/                   # Exemplos de uso (.tf)
├── templates/                  # Templates para geração de documentação (tfplugindocs)
└── docs/                       # Documentação autogerada e específica
```

## Fluxo de Dados
1. **Configuração (.tf):** O usuário define os recursos Coolify em arquivos HCL.
2. **Provider Initialization:** O Terraform carrega o provedor através do `main.go`.
3. **Resource Lifecycle:** Quando o Terraform executa um `apply`, ele chama as funções de ciclo de vida (`Create`, `Read`, `Update`, `Delete`) na pasta `coolify/[recurso]`.
4. **SDK Interaction:** O provedor utiliza o `coolify-sdk-go` para fazer chamadas HTTP à API do Coolify.
5. **State Management:** O provedor converte os dados da API para o `Model` do Terraform e persiste no estado (`.tfstate`).

## Considerações de Design
- **Desacoplamento:** O provedor não conhece os detalhes de implementação da API, ele depende exclusivamente do SDK.
- **Imutabilidade:** Onde possível, recursos são tratados como imutáveis para garantir a idempotência do Terraform.
