# AI Agent Rules & Project Standards (AGENTS.md)

Este arquivo define as regras e diretrizes para IAs (como Cursor, GitHub Copilot) e desenvolvedores que atuam neste repositório.

## Regras de Ouro
1. **Não commitar segredos:** Nunca suba arquivos `.env`, chaves privadas ou tokens. Certifique-se de que o `.gitignore` está atualizado.
2. **Qualidade de Código:** Siga rigorosamente os princípios de **Clean Code** e **SOLID**.
3. **Padrão de Provider:** Utilize o `Terraform Plugin Framework` (evite o SDKv2 legado, a menos que seja estritamente necessário por compatibilidade).
4. **Acoplamento:** Mantenha a lógica de negócio no SDK (`coolify-sdk-go`) e apenas o mapeamento de estado no Provider.

## Padrões de Código
- **DTOs:** Sempre use structs de DTO para validação de entrada/saída na comunicação com a API.
- **Mapeamento de Modelos:** Siga o padrão de nomenclatura `XModel` para recursos do Terraform e funções `mapXToModel` / `mapModelToX` para conversão.
- **Tratamento de Erros:** Não ignore erros. Use `diag.Diagnostics` para reportar erros ao Terraform de forma estruturada.
- **Conventional Commits:** Use o padrão `feat:`, `fix:`, `docs:`, `chore:`, etc.

## Processo de Deploy & GitOps
- **Automação:** O deploy do provider é automatizado via GitHub Actions (ver `.github/workflows`).
- **Webhooks:** Mudanças no SDK disparam builds no Provider para garantir compatibilidade.
- **Tags de Versão:** Siga o Semantic Versioning (SemVer). Versões fixas devem ser tageadas (ex: `v1.5.5`).

## Restrições Técnicas
- **Tipagem:** Seja rigoroso com tipos de dados (ex: `int32` vs `int64` vs `string` em portas). Valide schemas antes de qualquer push.
- **Dependências:** Evite adicionar dependências externas desnecessárias no Provider. Prefira o uso do SDK oficial.

---
*Este documento deve ser consultado por qualquer agente antes de iniciar uma nova task.*
