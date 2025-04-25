package report

import (
	"fmt"
	"time"

	"load-tester/request"
)

func GenerateReport(results []request.Result, duration time.Duration) {
	var totalRequests = len(results)
	var statusCount = make(map[int]int)

	// Contabiliza os códigos de status
	for _, result := range results {
		statusCount[result.StatusCode]++
	}

	// Exibe o relatório
	fmt.Println("Relatório do Teste de Carga")
	fmt.Println("===========================")
	fmt.Printf("Tempo total de execução: %s\n", duration)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("Status 200 (OK): %d\n", statusCount[200])
	fmt.Println("Distribuição de outros status:")
	for status, count := range statusCount {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}
