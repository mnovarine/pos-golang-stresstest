#!/bin/bash

echo "=== Teste do StressTest CLI ==="
echo ""

echo "1. Teste com parâmetros padrão..."
go run main.go --url=https://httpbin.org/get

echo ""
echo "2. Parâmetros personalizados..."
go run main.go --url=https://httpbin.org/get --requests=1000 --concurrency=50

echo ""
echo "3. Testar a página inicial do Google..."
go run main.go --url=https://google.com --requests=100 --concurrency=10

echo ""
echo "4. Teste com alta simultaneidade..."
go run main.go --url=https://httpbin.org/delay/1 --requests=200 --concurrency=50

echo ""
echo "5. Teste de error handling com endpoint inexistente..."
go run main.go --url=https://httpbin.org/status/404 --requests=5 --concurrency=2

echo ""
echo "6. Teste de timeout..."
go run main.go --url=https://httpbin.org/delay/2 --requests=3 --concurrency=2

echo ""
echo "=== Teste Completo ==="