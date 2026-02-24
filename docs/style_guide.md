# Guia de Estilos de Código (Style Guide)

Este guia define os padrões estéticos e estruturais de código para garantir consistência e manutenibilidade.

## 1. Lógica e Estrutura
- **DTOs para Tudo:** Sempre utilizar structs de `DTO` (Data Transfer Objects) para validação de payloads da API. Nunca exponha structs internas diretamente.
- **Fail Fast:** Valide inputs no início das funções e retorne o erro imediatamente.
- **Acoplamento Mínimo:** O provedor deve interagir apenas com as interfaces públicas do SDK.

## 2. Nomenclatura e Tipagem
- **Clareza sobre Conclusão:** Use nomes de variáveis descritivos. Evite abreviações obscuras.
- **Tipagem Forte:** Utilize os tipos corretos para cada dado. Portas devem ser `int`, IDs devem seguir o formato original (`string` para UUIDs, `int` para IDs numéricos).
- **Terraform Types:** Utilize sempre `types.String`, `types.Int64`, `types.Bool` do plugin framework para garantir compatibilidade com o estado do Terraform.

## 3. Git e Entrega
- **Conventional Commits:** Siga rigorosamente:
  - `feat:` para novas funcionalidades.
  - `fix:` para correções de bugs.
  - `docs:` para alterações em documentação.
  - `refactor:` para mudanças no código que não corrigem bugs nem adicionam features.
- **Branches:** Utilize o padrão `feature/`, `fix/` ou `hotfix/`.

## 4. Testes
- **Testes de Aceitação:** Todo novo recurso deve vir acompanhado de testes de aceitação (`_test.go`).
- **Mocking:** Prefira mocks para testes unitários do SDK para não depender da API real em ambientes de CI.

---
*Manter o código limpo é responsabilidade de todos.*
