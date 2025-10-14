package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

// rootCmd representa a base command quando chamado sem subcomandos
var rootCmd = &cobra.Command{
	Use:   "stresstest",
	Short: "Uma ferramenta de linha de comando para teste de carga em serviços web",
	Long: `StressTest é uma ferramenta de linha de comando construída em Go para realizar testes de carga em serviços web.
Permite especificar o número de requisições concorrentes e fornece relatórios detalhados
sobre tempos de resposta, códigos de status e métricas de desempenho.

Exemplo de uso:
  stresstest --url=https://example.com --requests=1000 --concurrency=10`,
	Run: runStressTest,
}

// Execute adiciona todos os subcomandos ao comando raiz e define as flags adequadamente.
// Isso é chamado por main.main(). Só precisa acontecer uma vez para o rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (obrigatório)")
	rootCmd.Flags().IntVar(&requests, "requests", 100, "Número total de requisições a serem feitas")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 10, "Número de requisições concorrentes")

	// Marcar flags obrigatórias
	rootCmd.MarkFlagRequired("url")
}

func runStressTest(cmd *cobra.Command, args []string) {
	// Validar entradas
	if requests <= 0 {
		fmt.Fprintf(os.Stderr, "Error: requests deve ser maior que 0\n")
		os.Exit(1)
	}

	if concurrency <= 0 {
		fmt.Fprintf(os.Stderr, "Error: concurrency deve ser maior que 0\n")
		os.Exit(1)
	}

	if concurrency > requests {
		concurrency = requests
		fmt.Printf("Warning: Concurrency reduzido para %d (não pode exceder o total de requisições)\n", concurrency)
	}

	fmt.Printf("Iniciando teste de carga...\n")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Total de Requisições: %d\n", requests)
	fmt.Printf("Concorrência: %d\n", concurrency)
	fmt.Println("-" + strings.Repeat("-", 50))

	// Executa o teste de carga
	result := performLoadTest(url, requests, concurrency)

	// Gera e exibe o relatório
	generateReport(result)
}
