#!/bin/bash

# Verifica se o binário já existe, caso contrário compila o código
if [ ! -f "./stresstest" ]; then
    echo "Compilando o binário stresstest..."
    go build -o stresstest .
    if [ $? -ne 0 ]; then
        echo "Falha ao compilar. Por favor, verifique se há erros."
        exit 1
    fi
fi

echo "=== Teste do StressTest CLI ==="
echo ""

echo "1. Teste com parâmetros padrão..."
./stresstest --url=https://httpbin.org/get

echo ""
echo "2. Parâmetros personalizados..."
./stresstest --url=https://httpbin.org/get --requests=1000 --concurrency=50

echo ""
echo "3. Testar a página inicial do Google..."
./stresstest --url=https://google.com --requests=100 --concurrency=10

echo ""
echo "4. Teste com alta simultaneidade..."
./stresstest --url=https://httpbin.org/delay/1 --requests=200 --concurrency=50

echo ""
echo "5. Teste de error handling com endpoint inexistente..."
./stresstest --url=https://httpbin.org/status/404 --requests=5 --concurrency=2

echo ""
echo "6. Teste de timeout..."
./stresstest --url=https://httpbin.org/delay/2 --requests=3 --concurrency=2

echo ""
echo "=== Teste Completo ==="