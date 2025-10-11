package cmd

import (
	"fmt"
	"sort"
	"strings"
)

// generateReport cria e exibe um relatório abrangente dos resultados do teste de carga
func generateReport(result *LoadTestResult) {
	fmt.Println()
	fmt.Println("Relatório do Teste de Carga")
	fmt.Println("=" + strings.Repeat("=", 50))

	// Estatísticas Básicas
	fmt.Printf("Tempo Total de Execução: %v\n", result.TotalDuration)
	fmt.Printf("Total de Requisições: %d\n", result.TotalRequests)
	fmt.Printf("Requisições Bem-Sucedidas (200): %d\n", result.SuccessCount)
	fmt.Printf("Requisições com Falhas: %d\n", result.ErrorCount)

	successRate := float64(result.SuccessCount) / float64(result.TotalRequests) * 100
	fmt.Printf("Taxa de Sucesso: %.2f%%\n", successRate)

	if result.TotalDuration > 0 {
		requestsPerSecond := float64(result.TotalRequests) / result.TotalDuration.Seconds()
		fmt.Printf("Requisições por Segundo: %.2f\n", requestsPerSecond)
	}

	fmt.Println()

	// Estatísticas de Tempo de Resposta
	// if len(result.Results) > 0 && result.SuccessCount > 0 {
	// 	fmt.Println("Estatísticas de Tempo de Resposta:")
	// 	fmt.Println("-" + strings.Repeat("-", 30))
	// 	fmt.Printf("Min Tempo de Resposta: %v\n", result.MinDuration)
	// 	fmt.Printf("Max Tempo de Resposta: %v\n", result.MaxDuration)
	// 	fmt.Printf("Avg Tempo de Resposta: %v\n", result.AvgDuration)

	// 	// Calculando percentis
	// 	durations := make([]time.Duration, 0, result.SuccessCount)
	// 	for _, res := range result.Results {
	// 		if res.Error == nil {
	// 			durations = append(durations, res.Duration)
	// 		}
	// 	}

	// 	if len(durations) > 0 {
	// 		sort.Slice(durations, func(i, j int) bool {
	// 			return durations[i] < durations[j]
	// 		})

	// 		p50 := durations[len(durations)*50/100]
	// 		p90 := durations[len(durations)*90/100]
	// 		p95 := durations[len(durations)*95/100]
	// 		p99 := durations[len(durations)*99/100]

	// 		fmt.Printf("Percentil 50: %v\n", p50)
	// 		fmt.Printf("Percentil 90: %v\n", p90)
	// 		fmt.Printf("Percentil 95: %v\n", p95)
	// 		fmt.Printf("Percentil 99: %v\n", p99)
	// 	}

	// 	fmt.Println()
	// }

	// Distribuição de Códigos de Status
	fmt.Println("Distribuição de Códigos de Status HTTP:")
	fmt.Println("-" + strings.Repeat("-", 35))

	// Ordena os códigos de status para uma saída consistente
	var statusCodes []int
	for code := range result.StatusCodes {
		statusCodes = append(statusCodes, code)
	}
	sort.Ints(statusCodes)

	for _, code := range statusCodes {
		count := result.StatusCodes[code]
		percentage := float64(count) / float64(result.TotalRequests) * 100
		fmt.Printf("Status %d: %d requisições (%.2f%%)\n", code, count, percentage)
	}

	// Detalhes de Erro
	// if result.ErrorCount > 0 {
	// 	fmt.Println()
	// 	fmt.Println("Detalhes de Erro:")
	// 	fmt.Println("-" + strings.Repeat("-", 20))

	// 	errorCounts := make(map[string]int)
	// 	for _, res := range result.Results {
	// 		if res.Error != nil {
	// 			errorType := getErrorType(res.Error)
	// 			errorCounts[errorType]++
	// 		}
	// 	}

	// 	for errorType, count := range errorCounts {
	// 		percentage := float64(count) / float64(result.ErrorCount) * 100
	// 		fmt.Printf("%s: %d errors (%.2f%% of errors)\n", errorType, count, percentage)
	// 	}
	// }

	fmt.Println()
	fmt.Println("Teste concluído com sucesso!")
	fmt.Println("=" + strings.Repeat("=", 50))
}

// getErrorType categoriza os tipos de erro para relatório
// func getErrorType(err error) string {
// 	errStr := err.Error()

// 	if strings.Contains(errStr, "timeout") {
// 		return "Timeout"
// 	}
// 	if strings.Contains(errStr, "connection refused") {
// 		return "Conexão Recusada"
// 	}
// 	if strings.Contains(errStr, "no such host") {
// 		return "Resolução DNS"
// 	}
// 	if strings.Contains(errStr, "network is unreachable") {
// 		return "Rede Inacessível"
// 	}
// 	if strings.Contains(errStr, "certificate") {
// 		return "Certificado SSL"
// 	}

// 	return "Outro"
// }
