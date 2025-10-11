# StressTest CLI - Ferramenta de Teste de Carga

Uma ferramenta CLI construída em Go para realizar testes de carga em serviços web.

## Funcionalidades

- 🚀 **Teste de Carga Simultâneo**: Configure requisições simultâneas para testar seus serviços
- 📊 **Relatórios Detalhados**: Estatísticas abrangentes incluindo tempos de resposta, códigos de status e percentis
- ⚡ **Rápido & Eficiente**: Construído com o excelente suporte de simultaneidade do Go
- 📈 **Métricas de Desempenho**: Análise de tempos de resposta mín/máx/médio e percentis
- 🎯 **Análise de Erros**: Relatório categorizado de erros para melhor depuração

## Uso

### Parâmetros da Linha de Comando

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requisições a fazer (padrão: 100)
- `--concurrency`: Número de requisições simultâneas (padrão: 10)

### Exemplos

#### Uso Básico

```bash
# Teste com parâmetros padrão
./stresstest --url=https://httpbin.org/get

# Parâmetros personalizados
./stresstest --url=https://httpbin.org/get --requests=1000 --concurrency=50

# Testar a página inicial do Google
./stresstest --url=https://google.com --requests=100 --concurrency=10

# Teste de carga em um endpoint de API
./stresstest --url=https://api.example.com/users --requests=500 --concurrency=25

# Teste com alta simultaneidade
./stresstest --url=https://httpbin.org/delay/1 --requests=200 --concurrency=50
```

## Exemplo de Saída

```
Iniciando teste de carga...
URL: https://httpbin.org/get
Total de Requisições: 100
Concorrência: 10
==================================================

Relatório de Teste de Carga
==================================================
Tempo Total de Execução: 2.345s
Total de Requisições: 100
Requisições Bem-sucedidas (200): 95
Requisições Falhadas: 5
Taxa de Sucesso: 95.00%
Requisições por Segundo: 42.64

Distribuição de Códigos de Status HTTP:
-----------------------------------
Status 200: 95 requisições (95.00%)
Status 500: 3 requisições (3.00%)
Status 502: 2 requisições (2.00%)

Teste concluído com sucesso!
==================================================
```

## Explicação das Métricas do Relatório

### Estatísticas Básicas
- **Tempo Total de Execução**: Tempo total para completar todo o teste
- **Requisições Bem-sucedidas (200)**: Número de respostas HTTP 200
- **Requisições Falhadas**: Requisições que resultaram em erros ou códigos de status diferentes de 200
- **Taxa de Sucesso**: Porcentagem de requisições bem-sucedidas
- **Requisições por Segundo**: Medida de throughput

### Distribuição de Códigos de Status
Mostra a distribuição dos códigos de status HTTP retornados pelo servidor.

## Licença

Licença MIT - veja o arquivo LICENSE para detalhes.