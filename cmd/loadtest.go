package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// TestResult armazena os resultados de uma única requisição HTTP
type TestResult struct {
	StatusCode  int
	Duration    time.Duration
	Error       error
	RequestTime time.Time
}

// LoadTestResult armazena os resultados agregados do teste de carga
type LoadTestResult struct {
	TotalRequests int
	TotalDuration time.Duration
	Results       []TestResult
	StatusCodes   map[int]int
	SuccessCount  int
	ErrorCount    int
	MinDuration   time.Duration
	MaxDuration   time.Duration
	AvgDuration   time.Duration
}

// performLoadTest executa o teste de carga com os parâmetros fornecidos
func performLoadTest(url string, totalRequests, concurrency int) *LoadTestResult {
	startTime := time.Now()

	// Cria canais para distribuição de trabalho e coleta de resultados
	requestChan := make(chan int, totalRequests)
	resultChan := make(chan TestResult, totalRequests)

	// Preenche o canal de requisições com números de requisição
	for i := 0; i < totalRequests; i++ {
		requestChan <- i + 1
	}
	close(requestChan)

	// Inicia goroutines de worker
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(url, requestChan, resultChan, &wg)
	}

	// Aguarda todos os workers concluírem
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Coleta resultados
	results := make([]TestResult, 0, totalRequests)
	statusCodes := make(map[int]int)
	successCount := 0
	errorCount := 0

	for result := range resultChan {
		results = append(results, result)

		if result.Error != nil {
			errorCount++
		} else {
			statusCodes[result.StatusCode]++
			if result.StatusCode == 200 {
				successCount++
			}
		}
	}

	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	// Calcula estatísticas de duração
	minDuration := time.Duration(0)
	maxDuration := time.Duration(0)
	totalResponseTime := time.Duration(0)
	validResponses := 0

	for _, result := range results {
		if result.Error == nil {
			if minDuration == 0 || result.Duration < minDuration {
				minDuration = result.Duration
			}
			if result.Duration > maxDuration {
				maxDuration = result.Duration
			}
			totalResponseTime += result.Duration
			validResponses++
		}
	}

	var avgDuration time.Duration
	if validResponses > 0 {
		avgDuration = totalResponseTime / time.Duration(validResponses)
	}

	return &LoadTestResult{
		TotalRequests: totalRequests,
		TotalDuration: totalDuration,
		Results:       results,
		StatusCodes:   statusCodes,
		SuccessCount:  successCount,
		ErrorCount:    errorCount,
		MinDuration:   minDuration,
		MaxDuration:   maxDuration,
		AvgDuration:   avgDuration,
	}
}

// worker performa requisições HTTP para o teste de carga
func worker(url string, requestChan <-chan int, resultChan chan<- TestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	for requestNum := range requestChan {
		result := makeRequest(client, url, requestNum)
		resultChan <- result
	}
}

// makeRequest realiza uma única requisição HTTP e mede o tempo de resposta
func makeRequest(client *http.Client, url string, requestNum int) TestResult {
	startTime := time.Now()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TestResult{
			Error:       fmt.Errorf("falha ao criar requisição: %v", err),
			RequestTime: startTime,
			Duration:    time.Since(startTime),
		}
	}

	// Adiciona User-Agent header para identificar nosso testador de carga
	req.Header.Set("User-Agent", "StressTest-CLI/1.0")

	resp, err := client.Do(req)
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	if err != nil {
		return TestResult{
			Error:       err,
			RequestTime: startTime,
			Duration:    duration,
		}
	}
	defer resp.Body.Close()

	return TestResult{
		StatusCode:  resp.StatusCode,
		Duration:    duration,
		RequestTime: startTime,
	}
}
