# StressTest CLI

**Objetivo**: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Uso

### Parâmetros da Linha de Comando

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requisições a fazer (padrão: 100)
- `--concurrency`: Número de requisições simultâneas (padrão: 10)

### Exemplos

```bash
# Executar o arquivo teste.sh para testar todos os cenários de teste
./teste.sh

# Para testar os cenários específicos, utilizar os exemplos abaixo

# 1. Teste com parâmetros padrão...
go run main.go --url=https://httpbin.org/get

# 2. Parâmetros personalizados...
go run main.go --url=https://httpbin.org/get --requests=1000 --concurrency=50

# 3. Testar a página inicial do Google...
go run main.go --url=https://google.com --requests=100 --concurrency=10

# 4. Teste com alta simultaneidade...
go run main.go --url=https://httpbin.org/delay/1 --requests=200 --concurrency=50

# 5. Teste de error handling com endpoint inexistente...
go run main.go --url=https://httpbin.org/status/404 --requests=5 --concurrency=2

# 6. Teste de timeout...
go run main.go --url=https://httpbin.org/delay/2 --requests=3 --concurrency=2
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
